package business

import "github.com/google/uuid"

type Business struct {
		ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
