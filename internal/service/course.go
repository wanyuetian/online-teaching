package service

import (
	"context"
	v1 "online-teaching/api/teaching/v1"
	"online-teaching/internal/biz"
)

// CourseService is a course service.
type CourseService struct {
	v1.UnimplementedCourseServer

	uc *biz.CourseUsecase
}

// NewCourseService new a course service.
func NewCourseService(uc *biz.CourseUsecase) *CourseService {
	return &CourseService{uc: uc}
}

// SearchCourse implements teaching.CourseService.
func (s *CourseService) SearchCourse(ctx context.Context, req *v1.SearchCourseRequest) (*v1.SearchCourseReply, error) {
	reply := &v1.SearchCourseReply{}
	courses, err := s.uc.SearchCourse(ctx, req.Keyword)

	if err == nil {
		reply.Code = -1
		return reply, err
	}

	reply.Data = &v1.SearchCourseReply_Data{
		Courses: make([]*v1.SearchCourseReply_Course, 0),
	}
	for _, course := range courses {
		reply.Data.Courses = append(reply.Data.Courses, &v1.SearchCourseReply_Course{
			Name:            course.Name,
			Desc:            course.Desc,
			BackgroundImage: course.BackgroundImage,
		})
	}
	return reply, nil
}

func (s *CourseService) CreateCourse(ctx context.Context, req *v1.CreateCourseRequest) (*v1.CreateCourseReply, error) {
	reply := &v1.CreateCourseReply{
		Data:    &v1.CreateCourseReply_Data{Course: &v1.CreateCourseReply_Course{}},
		Code:    0,
		Message: "",
	}
	course, err := s.uc.CreateCourse(ctx, req)
	if err != nil {
		reply.Code = -1
		return nil, err
	}
	reply.Data.Course = &v1.CreateCourseReply_Course{
		Name:            course.Name,
		Desc:            course.Desc,
		BackgroundImage: course.BackgroundImage,
		Summary:         course.Summary,
	}
	return reply, nil
}
