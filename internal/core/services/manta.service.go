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
	// Open file to parse
	f, err := os.Open(filename)
	if err != nil {
		return nil, OpenFileError{filename: filename, error: err}
	}
	// Handle defer errors
	defer func(name string) {
		if tempErr := os.Remove(name); tempErr != nil {
			err = RemoveFileError{filename: filename, error: tempErr}
		}
	}(filename)
	defer func(f *os.File) {
		if tempErr := f.Close(); tempErr != nil {
			err = CloseFileError{filename: filename, error: tempErr}
		}
	}(f)
	// Create stream parser
	p, err := manta.NewStreamParser(f)
	if err != nil {
		return nil, ParserCreationError{err}
	}
	defer p.Stop()
	// Declare some variables for parsing
	var (
		gameCurrentTime, gameStartTime float64
		gamePaused                     bool
		pauseStartTick                 int32
		totalPausedTicks               int32

		heroPlayers = make([]dtos.HeroPlayer, 10)
		glyphs      []models.Glyph
		glyph       models.Glyph

		magicTime = 1100.0 // Time when heroes loaded TODO
	)

	p.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
		if m.GetOrderType() == int32(dota.DotaunitorderT_DOTA_UNIT_ORDER_GLYPH) {
			entity := p.FindEntity(m.GetEntindex())
			glyph = models.Glyph{
				MatchID:     match.ID,
				Username:    entity.Get("m_iszPlayerName").(string),
				UserSteamID: strconv.FormatInt(int64(entity.Get("m_steamID").(uint64)), 10),
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
		switch e.GetClassName() {
		case "CDOTAGamerulesProxy":
			gameStartTime = float64(e.Get("m_pGameRules.m_flGameStartTime").(float32))
			gamePaused = e.Get("m_pGameRules.m_bGamePaused").(bool)
			pauseStartTick = e.Get("m_pGameRules.m_nPauseStartTick").(int32)
			totalPausedTicks = e.Get("m_pGameRules.m_nTotalPausedTicks").(int32)
			if gamePaused {
				gameCurrentTime = float64((pauseStartTick - totalPausedTicks) / 30)
			} else {
				gameCurrentTime = float64((int32(p.NetTick) - totalPausedTicks) / 30)
			}
		case "CDOTA_PlayerResource":
			if gameCurrentTime < magicTime {
				for i := 0; i < 10; i++ {
					heroPlayers[i].HeroID, _ = e.GetInt32("m_vecPlayerTeamData.000" + strconv.Itoa(i) + ".m_nSelectedHeroID")
					heroPlayers[i].PlayerID, _ = e.GetUint64("m_vecPlayerData.000" + strconv.Itoa(i) + ".m_iPlayerSteamID")
				}
			}
		}
		return nil
	})

	if err = p.Start(); err != nil {
		return nil, ParserError{err}
	}

	for k := range glyphs {
		for l := range heroPlayers {
			if glyphs[k].UserSteamID == strconv.FormatInt(int64(heroPlayers[l].PlayerID), 10) {
				glyphs[k].HeroID = heroPlayers[l].HeroID
				break
			}
		}
	}
	if len(glyphs) == 0 {
		return glyphs, NoGlyphsError{}
	}
	return glyphs, err
}
