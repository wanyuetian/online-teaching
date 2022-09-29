package data

import (
	"context"
	"online-teaching/internal/biz"
	"online-teaching/internal/ent"
	"online-teaching/internal/ent/teacher"

	"github.com/go-kratos/kratos/v2/log"
)

type teacherRepo struct {
	data *Data
	log  *log.Helper
}

func (r *teacherRepo) Delete(ctx context.Context, id int) (*ent.Teacher, error) {
	return r.data.Client.Teacher.UpdateOneID(id).SetIsDeleted(true).Save(ctx)
}

func (r *teacherRepo) Update(ctx context.Context, teacher *ent.Teacher) (*ent.Teacher, error) {
	_, err := r.FindByID(ctx, teacher.ID)
	if err != nil {
		return nil, err
	}

	return r.data.Client.Teacher.UpdateOne(teacher).
		SetName(teacher.Name).
		SetDesc(teacher.Desc).
		SetAvatar(teacher.Avatar).
		SetTitle(teacher.Title).
		Save(ctx)
}

func (r *teacherRepo) Create(ctx context.Context, teacher *ent.Teacher) (*ent.Teacher, error) {
	return r.data.Client.Teacher.Create().
		SetName(teacher.Name).
		SetDesc(teacher.Desc).
		SetAvatar(teacher.Avatar).
		SetTitle(teacher.Title).
		Save(ctx)
}

func (r *teacherRepo) FindByID(ctx context.Context, id int) (*ent.Teacher, error) {
	return r.data.Client.Teacher.Get(ctx, id)
}

func (r *teacherRepo) ListByKeyword(ctx context.Context, keyword string) ([]*ent.Teacher, error) {
	return r.data.Client.Teacher.Query().Where(teacher.Or(
		teacher.NameContains(keyword),
		teacher.DescContains(keyword)),
		teacher.IsDeleted(false),
	).All(ctx)
}

// NewTeacherRepo .
func NewTeacherRepo(data *Data, logger log.Logger) biz.TeacherRepo {
	return &teacherRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
