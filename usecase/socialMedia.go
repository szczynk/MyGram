package usecase

import (
	"context"

	"github.com/szczynk/MyGram/models"
)

type socialMediaUsecase struct {
	sr models.SocialMediaRepo
}

func NewSocialMediaUsecase(sr models.SocialMediaRepo) *socialMediaUsecase {
	return &socialMediaUsecase{sr}
}

func (suc *socialMediaUsecase) Fetch(c context.Context, m *[]models.SocialMedia, userID uint) (err error) {
	if err = suc.sr.Fetch(c, m, userID); err != nil {
		return err
	}
	return
}

func (suc *socialMediaUsecase) Store(c context.Context, m *models.SocialMedia) (err error) {
	if err = suc.sr.Store(c, m); err != nil {
		return err
	}
	return
}

func (suc *socialMediaUsecase) GetByUserID(c context.Context, m *models.SocialMedia, id uint) (err error) {
	if err = suc.sr.GetByUserID(c, m, id); err != nil {
		return err
	}
	return
}

func (suc *socialMediaUsecase) Update(c context.Context, mu models.SocialMedia, id uint) (p models.SocialMedia, err error) {
	p, err = suc.sr.Update(c, mu, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (suc *socialMediaUsecase) Delete(c context.Context, id uint) (err error) {
	if err = suc.sr.Delete(c, id); err != nil {
		return err
	}
	return
}
