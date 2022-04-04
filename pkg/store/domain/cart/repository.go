package cart

import "gorm.io/gorm"

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}
func (r *CartRepository) Migration() {
	r.db.AutoMigrate(&Cart{})
}