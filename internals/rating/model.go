package rating

import (
	"github.com/google/uuid"
)

type Rating struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}
