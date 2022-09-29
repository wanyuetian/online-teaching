package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "online-teaching/api/teaching/v1"
	"online-teaching/internal/ent"
)

// TeacherRepo is a Course repo.
type TeacherRepo interface {
	Update(context.Context, *ent.Teacher) (*ent.Teacher, error)
	Create(context.Context, *ent.Teacher) (*ent.Teacher, error)
	FindByID(context.Context, int) (*ent.Teacher, error)
	ListByKeyword(context.Context, string) ([]*ent.Teacher, error)
	Delete(context.Context, int) (*ent.Teacher, error)
}

// TeacherUsecase is a Course usecase.
type TeacherUsecase struct {
	repo TeacherRepo
	log  *log.Helper
}

func (uc *TeacherUsecase) SearchTeacher(ctx context.Context, keyword string) ([]*ent.Teacher, error) {
	return uc.repo.ListByKeyword(ctx, keyword)
}

func (uc *TeacherUsecase) GetTeacher(ctx context.Context, id int) (*ent.Teacher, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *TeacherUsecase) CreateTeacher(ctx context.Context, req *v1.CreateTeacherRequest) (*ent.Teacher, error) {
	teacher := &ent.Teacher{
		Name:   req.Name,
		Desc:   req.Desc,
		Avatar: req.Avatar,
		Title:  req.Title,
	}
	return uc.repo.Create(ctx, teacher)
}

func (uc *TeacherUsecase) UpdateTeacher(ctx context.Context, req *v1.UpdateTeacherRequest) (*ent.Teacher, error) {
	teacher := &ent.Teacher{
		ID:     int(req.TeacherId),
		Name:   req.Name,
		Desc:   req.Desc,
		Avatar: req.Avatar,
		Title:  req.Title,
	}
	return uc.repo.Update(ctx, teacher)
}

func (uc *TeacherUsecase) DeleteTeacher(ctx context.Context, id int32) (*ent.Teacher, error) {
	return uc.repo.Delete(ctx, int(id))
}

// NewTeacherUsecase new a Course usecase.
func NewTeacherUsecase(repo TeacherRepo, logger log.Logger) *TeacherUsecase {
	return &TeacherUsecase{repo: repo, log: log.NewHelper(logger)}
}
