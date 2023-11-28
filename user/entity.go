package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"not null;type:varchar(100)"`
	Email    string `json:"email" gorm:"not null;type:varchar(100)"`
	Password string `json:"-" gorm:"not null;type:varchar(255)"`
	Phone    string `json:"phone" gorm:"not null;type:varchar(20)"`
	Addres   string `json:"addres" gorm:"not null;type:varchar(255)"`
	Gender   string `json:"gender" gorm:"not null;type:varchar(10)"`
	Profile  string `json:"profile" gorm:"not null;type:varchar(255)"`
	// CreatedAt time.Time `json:"createdAt" gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
	// UpdatedAt time.Time `json:"updatedAt" gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
	// DeletedAt time.Time `json:"deletedAt" gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"`
}
