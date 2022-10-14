package usecase

import (
	"context"

	"github.com/szczynk/MyGram/models"
)

type commentUsecase struct {
	cr models.CommentRepo
}

func NewCommentUsecase(cr models.CommentRepo) *commentUsecase {
	return &commentUsecase{cr}
}

func (cuc *commentUsecase) Fetch(c context.Context, m *[]models.Comment) (err error) {
	if err = cuc.cr.Fetch(c, m); err != nil {
		return err
	}
	return
}

func (cuc *commentUsecase) Store(c context.Context, m *models.Comment) (err error) {
	if err = cuc.cr.Store(c, m); err != nil {
		return err
	}
	return
}

func (cuc *commentUsecase) GetByUserID(c context.Context, m *models.Comment, id uint) (err error) {
	if err = cuc.cr.GetByUserID(c, m, id); err != nil {
		return err
	}
	return
}

func (cuc *commentUsecase) Update(c context.Context, mu models.Comment, id uint) (p models.Comment, err error) {
	p, err = cuc.cr.Update(c, mu, id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (cuc *commentUsecase) Delete(c context.Context, id uint) (err error) {
	if err = cuc.cr.Delete(c, id); err != nil {
		return err
	}
	return
}
