package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	UserID  uint `json:"user_id"`
	PhotoID uint `json:"photo_id"`

	Message string `json:"message" valid:"required~Comment message is required" example:"Sheesh!" gorm:"not null"`

	User  *User  `json:"user,omitempty"`
	Photo *Photo `json:"photo,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return
}

func (c *Comment) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}
	return
}

type CommentUsecase interface {
	Fetch(context.Context, *[]Comment, uint) error
	Store(context.Context, *Comment) error
	GetByUserID(context.Context, *Comment, uint) error
	Update(context.Context, Comment, uint) (Photo, error)
	Delete(context.Context, uint) error
}

type CommentRepo interface {
	Fetch(context.Context, *[]Comment, uint) error
	Store(context.Context, *Comment) error
	GetByUserID(context.Context, *Comment, uint) error
	Update(context.Context, Comment, uint) (Photo, error)
	Delete(context.Context, uint) error
}
