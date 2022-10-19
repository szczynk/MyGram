package repository

import (
	"context"
	"errors"
	"time"

	"github.com/szczynk/MyGram/helpers"
	"github.com/szczynk/MyGram/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (ur *userRepo) Register(c context.Context, m *models.User) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = ur.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (ur userRepo) Login(c context.Context, m *models.User) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	originalPassword := m.Password

	err = ur.db.Debug().WithContext(ctx).Where("email = ?", m.Email).
		Take(&m).Error
	if err != nil {
		return err
	}

	isValid := helpers.CompareHash(m.Password, originalPassword)
	if !isValid {
		return errors.New("invalid password")
	}

	return
}

func (ur *userRepo) Update(c context.Context, mu models.User, id uint) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	user = models.User{}
	err = ur.db.Debug().WithContext(ctx).First(&user, id).Error
	if err != nil {
		return user, err
	}

	err = ur.db.Debug().WithContext(ctx).Model(&user).Where("id = ?", id).
		Updates(mu).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur userRepo) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = ur.db.Debug().WithContext(ctx).First(&models.User{}, id).Error
	if err != nil {
		return err
	}

	err = ur.db.Debug().WithContext(ctx).Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return
}
