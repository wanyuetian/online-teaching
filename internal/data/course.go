package data

import (
	"context"
	"online-teaching/internal/ent"
	"online-teaching/internal/ent/course"

	"online-teaching/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type courseRepo struct {
	data *Data
	log  *log.Helper
}

func (r *courseRepo) ListByKeyword(ctx context.Context, keyword string) ([]*ent.Course, error) {
	return r.data.Client.Course.Query().WithTeachers().Where(course.Or(
		course.NameContains(keyword),
		course.DescContains(keyword)),
	).All(ctx)
}

func (r *courseRepo) ListAll(ctx context.Context) ([]*ent.Course, error) {
	//TODO implement me
	panic("implement me")
}

// NewCourseRepo .
func NewCourseRepo(data *Data, logger log.Logger) biz.CourseRepo {
	return &courseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *courseRepo) Create(ctx context.Context, g *ent.Course) (*ent.Course, error) {
	//r.data.Client.CourseTeacher.Create()
	return r.data.Client.Course.Create().SetName(g.Name).
		SetDesc(g.Desc).
		SetImage(g.Image).
		Save(ctx)
}

func (r *courseRepo) Update(ctx context.Context, g *ent.Course) (*ent.Course, error) {
	return g, nil
}

func (r *courseRepo) FindByID(context.Context, int64) (*ent.Course, error) {
	return nil, nil
}
