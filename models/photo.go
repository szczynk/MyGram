package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`

	Title    string `json:"title" valid:"required~Photo title is required" example:"My Sweet Photo" gorm:"not null"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo url is required" example:"https://szczynk.github.io/blog/_nuxt/img/128affc.png" gorm:"not null"`
	Caption  string `json:"caption" example:"Beautiful as it is"`

	User     *User      `json:"user,omitempty"`
	Comments *[]Comment `json:"-" gorm:"constraint:OnDelete:SET NULL;"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}

	return
}

func (p *Photo) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

type PhotoUsecase interface {
	Fetch(context.Context, *[]Photo) error
	Store(context.Context, *Photo) error
	GetByID(context.Context, *Photo, uint) error
	GetByUserID(context.Context, *Photo, uint) error
	Update(context.Context, Photo, uint) (Photo, error)
	Delete(context.Context, uint) error
}

type PhotoRepo interface {
	Fetch(context.Context, *[]Photo) error
	Store(context.Context, *Photo) error
	GetByID(context.Context, *Photo, uint) error
	GetByUserID(context.Context, *Photo, uint) error
	Update(context.Context, Photo, uint) (Photo, error)
	Delete(context.Context, uint) error
}
