package repository

import (
	"gorm.io/gorm"
	"tiktok-arena/internal/core/models"
)

type GlyphRepository struct {
	db *gorm.DB
}

func NewGlyphRepository(db *gorm.DB) *GlyphRepository {
	return &GlyphRepository{db: db}
}

func (r *GlyphRepository) GetGlyphs(matchID int) ([]models.Glyph, error) {
	var glyphs []models.Glyph
	record := r.db.Where("match_id = ?", matchID).Find(&glyphs)
	return glyphs, record.Error
}

func (r *GlyphRepository) GlyphsExist(matchID int) (bool, error) {
	var glyph models.Glyph
	record := r.db.
		First(&glyph, "match_id = ?", matchID) // TODO
	if record.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	return glyph.MatchID != 0, record.Error
}

func (r *GlyphRepository) CreateGlyphs(newGlyphs []models.Glyph) error {
	record := r.db.
		Create(newGlyphs)
	return record.Error
}
