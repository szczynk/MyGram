package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/szczynk/MyGram/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Username        string `json:"username" valid:"required~Username is required" example:"Johndee" gorm:"not null;uniqueIndex;"`
	Email           string `json:"email" valid:"required~Email is required,email~Invalid email format" example:"johndee@gmail.com" gorm:"not null;uniqueIndex;"`
	Password        string `json:"password,omitempty" valid:"required~Password is required,minstringlength(6)~Your password must be at least 6 characters long" example:"12345678" gorm:"not null"`
	Age             int    `json:"age,omitempty" valid:"required~Age is required,range(8|100)~Your age must be at least greater than 8 years old" example:"8" gorm:"not null"`
	ProfileImageUrl string `json:"profile_image_url,omitempty" example:"https://avatars.dicebear.com/api/identicon/your-custom-seed.svg"`

	Photos      *[]Photo     `json:"-" gorm:"constraint:OnDelete:SET NULL;"`
	Comments    *[]Comment   `json:"-" gorm:"constraint:OnDelete:SET NULL;"`
	SocialMedia *SocialMedia `json:"-" gorm:"constraint:OnDelete:SET NULL;"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	u.Password = helpers.Hash(u.Password)
	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}
	return
}

type UserUsecase interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}

type UserRepo interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}
