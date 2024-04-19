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
	Password        string
	ProfileHeadline string
	ProfileID       uuid.UUID
	Profile         Profile
}

func (u *User) Valid() error {
	panic("unimplemented")
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

type File struct {
	gorm.Model
	UserID   uuid.UUID
	User     User
	FileName string
	FileType string
	FileData []byte // Store file content as bytes
}