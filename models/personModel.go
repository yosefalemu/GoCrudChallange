package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Person struct {
	ID        string         `gorm:"type:uuid;primarykey;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Age       int            `gorm:"not null" json:"age"`
	Hobbies   pq.StringArray `gorm:"type:text[];default:'{}'" json:"hobbies"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate hook to generate UUID if it’s not already set
func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}
