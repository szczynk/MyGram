package repository

import (
	"context"
	"time"

	"github.com/szczynk/MyGram/models"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *commentRepo {
	return &commentRepo{db}
}

func (cr commentRepo) Fetch(c context.Context, m *[]models.Comment, userID uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = cr.db.Debug().WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username")
		}).
		Preload("Photo", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "UserID", "Title", "PhotoUrl", "Caption")
		}).
		Find(&m).Error
	if err != nil {
		return err
	}
	return
}

func (cr *commentRepo) Store(c context.Context, m *models.Comment) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = cr.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (cr commentRepo) GetByUserID(c context.Context, m *models.Comment, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = cr.db.Debug().WithContext(ctx).Select("user_id").First(&m, id).Error
	if err != nil {
		return err
	}
	return
}

func (cr *commentRepo) Update(c context.Context, mu models.Comment, id uint) (photo models.Photo, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	comment := models.Comment{}
	photo = models.Photo{}

	err = cr.db.Debug().WithContext(ctx).First(&comment, id).Error
	if err != nil {
		return photo, err
	}

	err = cr.db.Debug().WithContext(ctx).Model(&comment).Where("id = ?", id).
		Updates(mu).Error
	if err != nil {
		return photo, err
	}

	err = cr.db.Debug().WithContext(ctx).First(&photo, comment.PhotoID).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (cr commentRepo) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = cr.db.Debug().WithContext(ctx).First(&models.Comment{}, id).Error
	if err != nil {
		return err
	}

	err = cr.db.Debug().WithContext(ctx).Delete(&models.Comment{}, id).Error
	if err != nil {
		return err
	}
	return
}
