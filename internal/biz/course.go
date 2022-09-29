package biz

import (
	"context"
	v1 "online-teaching/api/teaching/v1"
	"online-teaching/internal/ent"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// CourseRepo is a Course repo.
type CourseRepo interface {
	Create(context.Context, *ent.Course) (*ent.Course, error)
	Update(context.Context, *ent.Course) (*ent.Course, error)
	FindByID(context.Context, int64) (*ent.Course, error)
	ListByKeyword(context.Context, string) ([]*ent.Course, error)
	ListAll(context.Context) ([]*ent.Course, error)
}

// CourseUsecase is a Course usecase.
type CourseUsecase struct {
	repo CourseRepo
	log  *log.Helper
}

// NewCourseUsecase new a Course usecase.
func NewCourseUsecase(repo CourseRepo, logger log.Logger) *CourseUsecase {
	return &CourseUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SearchCourse search Courses, and returns the Courses.
func (uc *CourseUsecase) SearchCourse(ctx context.Context, keyword string) ([]*ent.Course, error) {
	return uc.repo.ListByKeyword(ctx, keyword)
}

// CreateCourse create Course, and returns the Course.
func (uc *CourseUsecase) CreateCourse(ctx context.Context, req *v1.CreateCourseRequest) (*ent.Course, error) {
	course := &ent.Course{
		Name:           req.Name,
		Desc:           req.Desc,
		Image:          req.Image,
		Tags:           req.Tags,
		Classification: req.Classification,
	}

	return uc.repo.Create(ctx, course)
}
