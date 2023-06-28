package models

type Glyph struct {
	MatchID     int    `gorm:"not null;default:null"`
	Username    string `gorm:"not null;default:null"`
	UserSteamID string `gorm:"not null;default:null"`
	Minute      uint32 `gorm:"not null;default:0"`
	Second      uint32 `gorm:"not null;default:0"`
	Team        uint64 `gorm:"not null;default:2"` // Radiant team is 2 and dire team is 3
	HeroID      int32  `gorm:"not null;default:1"` // ID of hero (https://liquipedia.net/dota2/MediaWiki:Dota2webapi-heroes.json)
}
