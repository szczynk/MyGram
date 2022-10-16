package usecase

import (
	"context"

	"github.com/szczynk/MyGram/models"
)

type photoUsecase struct {
	pr models.PhotoRepo
}

func NewPhotoUsecase(pr models.PhotoRepo) *photoUsecase {
	return &photoUsecase{pr}
}

func (puc *photoUsecase) Fetch(c context.Context, pa *models.Pagination) (err error) {
	if err = puc.pr.Fetch(c, pa); err != nil {
		return err
	}
	return
}

func (puc *photoUsecase) Store(c context.Context, m *models.Photo) (err error) {
	if err = puc.pr.Store(c, m); err != nil {
		return err
	}
	return
}

func (puc *photoUsecase) GetByID(c context.Context, m *models.Photo, id uint) (err error) {
	if err = puc.pr.GetByID(c, m, id); err != nil {
		return err
	}
	return
}

func (puc *photoUsecase) GetByUserID(c context.Context, m *models.Photo, id uint) (err error) {
	if err = puc.pr.GetByUserID(c, m, id); err != nil {
		return err
	}
	return
}

func (puc *photoUsecase) Update(c context.Context, mu models.Photo, id uint) (p models.Photo, err error) {
	p, err = puc.pr.Update(c, mu, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (puc *photoUsecase) Delete(c context.Context, id uint) (err error) {
	if err = puc.pr.Delete(c, id); err != nil {
		return err
	}
	return
}
