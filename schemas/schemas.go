package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primaryKey"`
	Name            string
	Email           string
	Address         string
	UserType        string
	PasswordHash    string
	ProfileHeadline string
	ProfileID       uuid.UUID
	Profile         Profile
}

type Profile struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primaryKey"`
	ResumeFileAddress string
	Skills            string
	Education         string
	Experience        string
	Name              string
	Email             string
	Phone             string
}

type Job struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primaryKey"`
	Title             string
	Description       string
	PostedOn          time.Time
	TotalApplications int
	CompanyName       string
	PostedByID        uuid.UUID `gorm:"unique"`
	PostedBy          User      `gorm:"foreignKey:PostedByID"`
}
