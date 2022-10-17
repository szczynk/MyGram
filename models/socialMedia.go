package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`

	Name           string `json:"name" valid:"required~Social media name is required" example:"johndee13" gorm:"not null"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social media url is required"  example:"johndee13url" gorm:"not null"`

	User *User `json:"user,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (s *SocialMedia) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}

	return
}

func (s *SocialMedia) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}
	return
}

type SocialMediaUsecase interface {
	Fetch(context.Context, *[]SocialMedia, uint) error
	Store(context.Context, *SocialMedia) error
	GetByUserID(context.Context, *SocialMedia, uint) error
	Update(context.Context, SocialMedia, uint) (SocialMedia, error)
	Delete(context.Context, uint) error
}

type SocialMediaRepo interface {
	Fetch(context.Context, *[]SocialMedia, uint) error
	Store(context.Context, *SocialMedia) error
	GetByUserID(context.Context, *SocialMedia, uint) error
	Update(context.Context, SocialMedia, uint) (SocialMedia, error)
	Delete(context.Context, uint) error
}
