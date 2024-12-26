package product

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}