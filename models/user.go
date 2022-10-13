package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `json:"id" gorm: "primary_key:auto_increment;"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Key       string    `json:"key" gorm:"type:varchar";not null;uniqueIndex`
	Secret    string    `json:"secret" gorm:"type:varchar;not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "signup":
		if u.Name == "" {
			return errors.New("required name")
		} else if u.Key == "" {
			return errors.New("required key")
		} else if u.Secret == "" {
			return errors.New("required secret")
		}
	}
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = strings.TrimSpace(u.Name)
	u.Key = strings.TrimSpace(u.Key)
	u.Secret = strings.TrimSpace(u.Secret)
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	if err := db.Debug().Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) GetUserInfo(db *gorm.DB, key string) (*User, error) {
	err := db.Debug().Model(User{}).Where("key = ?", key).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}
