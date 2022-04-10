package status

import (
	"github.com/gofrs/uuid"
)

type Status struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt string    `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt string    `json:"updated_at" gorm:"type:timestamp;not null"`
	DeletedAt string    `json:"deleted_at" gorm:"type:timestamp;default:null"`
}

func (Status) TableName() string {
	return "status"
}
