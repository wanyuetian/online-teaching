package service

import (
	"context"
	pb "online-teaching/api/teaching/v1"
	"online-teaching/internal/biz"
)

// CourseService is a course service.
type CourseService struct {
	pb.UnimplementedCourseServer

	uc *biz.CourseUsecase
}

// NewCourseService new a course service.
func NewCourseService(uc *biz.CourseUsecase) *CourseService {
	return &CourseService{uc: uc}
}

// SearchCourse implements teaching.CourseService.
func (s *CourseService) SearchCourse(ctx context.Context, req *pb.SearchCoursesRequest) (*pb.SearchCoursesReply, error) {
	reply := &pb.SearchCoursesReply{}
	courses, err := s.uc.SearchCourse(ctx, req.Keyword)

	if err == nil {
		return reply, err
	}

	reply.Data = &pb.SearchCoursesReply_Data{
		Courses: make([]*pb.SearchCoursesReply_Course, 0),
	}
	for _, course := range courses {
		c := &pb.SearchCoursesReply_Course{
			Name:           course.Name,
			Desc:           course.Desc,
			Image:          course.Image,
			Tags:           course.Tags,
			Classification: course.Classification,
			Teachers:       make([]*pb.SearchCoursesReply_Teacher, 0),
		}
		for _, t := range course.Edges.Teachers {
			c.Teachers = append(c.Teachers, &pb.SearchCoursesReply_Teacher{
				TeacherId: int32(t.ID),
				Name:      t.Name,
				Desc:      t.Desc,
				Avatar:    t.Avatar,
				Title:     t.Title,
			})
		}
		reply.Data.Courses = append(reply.Data.Courses, c)
	}
	return reply, nil
}

func (s *CourseService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseReply, error) {
	reply := &pb.CreateCourseReply{
		Data:    &pb.CreateCourseReply_Data{},
		Code:    0,
		Message: "",
	}
	course, err := s.uc.CreateCourse(ctx, req)
	if err != nil {
		return nil, err
	}
	reply.Data = &pb.CreateCourseReply_Data{
		Name:  course.Name,
		Desc:  course.Desc,
		Image: course.Image,
	}
	return reply, nil
}
