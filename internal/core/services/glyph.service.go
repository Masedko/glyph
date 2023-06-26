package services

import (
	"go-glyph-v2/internal/core/dtos"
	"go-glyph-v2/internal/core/models"
	"go-glyph-v2/internal/core/validator"
)

type GlyphServiceGlyphRepository interface {
	GetGlyphs(matchID int) ([]models.Glyph, error)
	GlyphsExist(matchID int) (bool, error)
	CreateGlyphs(newGlyphs []models.Glyph) error
}

type GlyphService struct {
	GlyphServiceGlyphRepository GlyphServiceGlyphRepository
}

func NewGlyphService(glyphServiceGlyphRepository GlyphServiceGlyphRepository) *GlyphService {
	return &GlyphService{
		GlyphServiceGlyphRepository: glyphServiceGlyphRepository,
	}
}

func (s *GlyphService) GetGlyphs(getGlyphs *dtos.GetGlyphs) (dtos.GlyphParse, error) {
	err := validator.ValidateStruct(getGlyphs)
	if err != nil {
		return dtos.GlyphParse{}, ValidateError{err}
	}

	matchParsed, err := s.GlyphServiceGlyphRepository.GlyphsExist(getGlyphs.MatchID)
	if err != nil {
		return dtos.GlyphParse{}, RepositoryError{err}
	}
	if !matchParsed {
		return dtos.GlyphParse{GlyphParsed: false}, nil
	}

	glyphs, err := s.GlyphServiceGlyphRepository.GetGlyphs(getGlyphs.MatchID)
	if err != nil {
		return dtos.GlyphParse{}, RepositoryError{err}
	}

	return dtos.GlyphParse{GlyphParsed: true, Glyphs: glyphs}, nil
}

func (s *GlyphService) CreateGlyphs(createGlyphs *dtos.CreateGlyphs) error {
	err := validator.ValidateStruct(createGlyphs)
	if err != nil {
		return ValidateError{err}
	}

	matchParsed, err := s.GlyphServiceGlyphRepository.GlyphsExist(createGlyphs.Glyphs[0].MatchID)
	if err != nil {
		return RepositoryError{err}
	}
	if matchParsed {
		return MatchAlreadyParsedError{}
	}

	err = s.GlyphServiceGlyphRepository.CreateGlyphs(createGlyphs.Glyphs)
	if err != nil {
		return RepositoryError{err}
	}
	return nil
}
