package repository

import (
	"context"
	"time"

	"github.com/szczynk/MyGram/models"
	"gorm.io/gorm"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) *photoRepo {
	return &photoRepo{db}
}

func (pr photoRepo) Fetch(c context.Context, m *[]models.Photo) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = pr.db.Debug().WithContext(ctx).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username")
		}).
		Find(&m).Error
	if err != nil {
		return err
	}
	return
}

func (pr *photoRepo) Store(c context.Context, m *models.Photo) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = pr.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (pr photoRepo) GetByID(c context.Context, m *models.Photo, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = pr.db.Debug().WithContext(ctx).First(&m, id).Error
	if err != nil {
		return err
	}
	return
}

func (pr photoRepo) GetByUserID(c context.Context, m *models.Photo, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = pr.db.Debug().WithContext(ctx).Select("user_id").First(&m, id).Error
	if err != nil {
		return err
	}
	return
}

func (pr *photoRepo) Update(c context.Context, mu models.Photo, id uint) (photo models.Photo, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	photo = models.Photo{}
	err = pr.db.Debug().WithContext(ctx).First(&photo, id).Error
	if err != nil {
		return photo, err
	}

	err = pr.db.Debug().WithContext(ctx).Model(&photo).Where("id = ?", id).
		Updates(mu).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (pr photoRepo) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = pr.db.Debug().WithContext(ctx).First(&models.Photo{}, id).Error
	if err != nil {
		return err
	}

	err = pr.db.Debug().WithContext(ctx).Delete(&models.Photo{}, id).Error
	if err != nil {
		return err
	}
	return
}
