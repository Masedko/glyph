package scopes

import "gorm.io/gorm"

// Scopes for search and pagination

func Search(searchText string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if searchText == "" {
			return db
		}
		return db.Select("*, levenshtein(name, ?) as distance", searchText).Order("distance")
	}
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * pageSize
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
