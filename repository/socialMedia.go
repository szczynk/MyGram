package repository

import (
	"context"
	"time"

	"github.com/szczynk/MyGram/models"
	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) *socialMediaRepo {
	return &socialMediaRepo{db}
}

func (sr socialMediaRepo) Fetch(c context.Context, m *[]models.SocialMedia, userID uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Email", "Username", "ProfileImageUrl")
		}).
		Find(&m).Error
	if err != nil {
		return err
	}
	return
}

func (sr *socialMediaRepo) Store(c context.Context, m *models.SocialMedia) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (sr socialMediaRepo) GetByUserID(c context.Context, m *models.SocialMedia, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).Select("user_id").First(&m, id).Error
	if err != nil {
		return err
	}
	return
}

func (sr *socialMediaRepo) Update(c context.Context, mu models.SocialMedia, id uint) (socialMedia models.SocialMedia, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	socialMedia = models.SocialMedia{}
	err = sr.db.Debug().WithContext(ctx).First(&socialMedia, id).Error
	if err != nil {
		return socialMedia, err
	}

	err = sr.db.Debug().WithContext(ctx).Model(&socialMedia).Where("id = ?", id).
		Updates(mu).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (sr socialMediaRepo) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = sr.db.Debug().WithContext(ctx).First(&models.SocialMedia{}, id).Error
	if err != nil {
		return err
	}

	err = sr.db.Debug().WithContext(ctx).Delete(&models.SocialMedia{}, id).Error
	if err != nil {
		return err
	}
	return
}
