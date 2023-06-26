package services

import (
	"fmt"
	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
	"go-glyph-v2/internal/core/dtos"
	"go-glyph-v2/internal/core/models"
	"golang.org/x/exp/slices"
	"math"
	"os"
	"strconv"
)

type MantaService struct {
}

func NewMantaService() *MantaService {
	return &MantaService{}
}

func (s MantaService) GetGlyphsFromDem(match dtos.Match) ([]models.Glyph, error) {
	filename := fmt.Sprintf("internal/data/demos/%d.dem", match.ID)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filename)
	defer f.Close()
	p, err := manta.NewStreamParser(f)
	defer p.Stop()
	if err != nil {
		return nil, err
	}
	var gameCurrentTime, gameStartTime float64
	var gamePaused bool
	var pauseStartTick int32
	var totalPausedTicks int32
	var glyphs []models.Glyph
	var glyph models.Glyph
	var heroplayers []dtos.HeroPlayer
	for i := 0; i < 10; i++ {
		heroplayers = append(heroplayers, dtos.HeroPlayer{})
	}

	p.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
		if m.GetOrderType() == int32(dota.DotaunitorderT_DOTA_UNIT_ORDER_GLYPH) {
			entity := p.FindEntity(m.GetEntindex())
			glyph = models.Glyph{
				MatchID:     match.ID,
				Username:    entity.Get("m_iszPlayerName").(string),
				UserSteamID: fmt.Sprint(entity.Get("m_steamID").(uint64)),
				Minute:      uint32(gameCurrentTime-gameStartTime) / 60,
				Second:      uint32(math.Round(gameCurrentTime-gameStartTime)) % 60,
				Team:        entity.Get("m_iTeamNum").(uint64),
			}
			if !slices.Contains(glyphs, glyph) {
				glyphs = append(glyphs, glyph)
			}
		}
		return nil
	})
	p.OnEntity(func(e *manta.Entity, op manta.EntityOp) error {
		if e.GetClassName() != "CDOTAGamerulesProxy" && e.GetClassName() != "CDOTA_PlayerResource" {
			return nil
		}
		if e.GetClassName() == "CDOTAGamerulesProxy" {
			gameStartTime = float64(e.Get("m_pGameRules.m_flGameStartTime").(float32))
			gamePaused = e.Get("m_pGameRules.m_bGamePaused").(bool)
			pauseStartTick = e.Get("m_pGameRules.m_nPauseStartTick").(int32)
			totalPausedTicks = e.Get("m_pGameRules.m_nTotalPausedTicks").(int32)
			if gamePaused {
				gameCurrentTime = float64((pauseStartTick - totalPausedTicks) / 30)
			} else {
				gameCurrentTime = float64((int32(p.NetTick) - totalPausedTicks) / 30)
			}
			return nil
		}
		if gameCurrentTime < 1100 && e.GetClassName() == "CDOTA_PlayerResource" {
			for i := 0; i < 10; i++ {
				heroplayers[i].HeroID, _ = e.GetInt32("m_vecPlayerTeamData.000" + strconv.Itoa(i) + ".m_nSelectedHeroID")
				heroplayers[i].PlayerID, _ = e.GetUint64("m_vecPlayerData.000" + strconv.Itoa(i) + ".m_iPlayerSteamID")
			}
			return nil
		}
		return nil
	})

	err = p.Start()
	if err != nil {
		return nil, err
	}
	for k := range glyphs {
		for l := range heroplayers {
			if fmt.Sprint(glyphs[k].UserSteamID) == fmt.Sprint(heroplayers[l].PlayerID) {
				glyphs[k].HeroID = heroplayers[l].HeroID
			}
		}
	}
	if len(glyphs) == 0 {
		return glyphs, NoGlyphsError{}
	}
	return glyphs, nil
}
