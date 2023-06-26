package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-glyph-v2/internal/core/dtos"
	"go-glyph-v2/internal/core/models"
	"strconv"
)

type GlyphService interface {
	GetGlyphs(getGlyphs *dtos.GetGlyphs) (dtos.GlyphParse, error)
	CreateGlyphs(createGlyphs *dtos.CreateGlyphs) error
}

type StratzService interface {
	GetMatchFromStratzAPI(matchID int) (dtos.Match, error)
}

type ValveService interface {
	RetrieveFile(match dtos.Match) error
}

type MantaService interface {
	GetGlyphsFromDem(match dtos.Match) ([]models.Glyph, error)
}

type GlyphController struct {
	GlyphService  GlyphService
	StratzService StratzService
	ValveService  ValveService
	MantaService  MantaService
}

func NewGlyphController(glyphService GlyphService, stratzService StratzService, valveService ValveService, mantaService MantaService) *GlyphController {
	return &GlyphController{
		GlyphService:  glyphService,
		StratzService: stratzService,
		ValveService:  valveService,
		MantaService:  mantaService,
	}
}

// GetGlyphs
//
//	@Summary		Get glyphs
//	@Description	Get glyphs using match id
//	@Tags			glyph
//	@Accept			json
//	@Produce		json
//	@Param			matchID					path		string						true	"Match ID"
//	@Success		200						{object}	[]models.Glyph				"Glyphs"
//	@Failure		400						{object}	dtos.MessageResponseType	"Glyphs parse error"
//	@Router			/api/glyph/{matchID}	[post]
func (cr *GlyphController) GetGlyphs(c *fiber.Ctx) error {
	matchIDString := c.Params("matchID")
	matchID, err := strconv.Atoi(matchIDString)
	if err != nil {
		return errors.New("wrong matchID(cannot convert to integer)")
	}
	// Check if parsed match is stored in db and retrieve if stored
	getGlyphes := &dtos.GetGlyphs{MatchID: matchID}
	glyphParse, err := cr.GlyphService.GetGlyphs(getGlyphes)
	if err != nil {
		return err
	}
	// If match is parsed -> return parsed match
	if glyphParse.GlyphParsed == true {
		return c.Status(fiber.StatusCreated).JSON(glyphParse.Glyphs)
	}
	// If not in db
	// Make request to STRATZ API
	match, err := cr.StratzService.GetMatchFromStratzAPI(matchID)
	if err != nil {
		return err
	}
	// Download from valve cluster
	err = cr.ValveService.RetrieveFile(match)
	if err != nil {
		return err
	}
	// Parse using Manta(Dotabuff golang parser)
	glyphs, err := cr.MantaService.GetGlyphsFromDem(match)
	if err != nil {
		return err
	}
	// Save parsed match to database
	createGlyphs := dtos.CreateGlyphs{Glyphs: glyphs}
	err = cr.GlyphService.CreateGlyphs(&createGlyphs)
	if err != nil {
		return err
	}
	// Return parsed match
	return c.Status(fiber.StatusCreated).JSON(glyphs)
}
