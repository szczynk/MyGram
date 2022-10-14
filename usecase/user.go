package usecase

import (
	"context"

	"github.com/szczynk/MyGram/models"
)

type userUsecase struct {
	ur models.UserRepo
}

func NewUserUsecase(ur models.UserRepo) *userUsecase {
	return &userUsecase{ur}
}

func (uuc *userUsecase) Register(c context.Context, m *models.User) (err error) {
	if err = uuc.ur.Register(c, m); err != nil {
		return err
	}
	return
}

func (uuc *userUsecase) Login(c context.Context, m *models.User) (err error) {
	if err = uuc.ur.Login(c, m); err != nil {
		return err
	}
	return
}

func (uuc *userUsecase) Update(c context.Context, mu models.User, id uint) (u models.User, err error) {
	u, err = uuc.ur.Update(c, mu, id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (uuc *userUsecase) Delete(c context.Context, id uint) (err error) {
	if err = uuc.ur.Delete(c, id); err != nil {
		return err
	}
	return
}
