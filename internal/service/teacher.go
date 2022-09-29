package service

import (
	"context"
	pb "online-teaching/api/teaching/v1"
	"online-teaching/internal/biz"
)

type TeacherService struct {
	pb.UnimplementedTeacherServer
	uc *biz.TeacherUsecase
}

func NewTeacherService(uc *biz.TeacherUsecase) *TeacherService {
	return &TeacherService{uc: uc}
}

func (s *TeacherService) SearchTeacher(ctx context.Context, req *pb.SearchTeachersRequest) (*pb.SearchTeachersReply, error) {
	reply := &pb.SearchTeachersReply{}
	teachers, err := s.uc.SearchTeacher(ctx, req.Keyword)

	if err != nil {
		return reply, err
	}

	reply.Data = &pb.SearchTeachersReply_Data{
		Teachers: make([]*pb.SearchTeachersReply_Teacher, 0),
	}
	for _, teacher := range teachers {
		reply.Data.Teachers = append(reply.Data.Teachers, &pb.SearchTeachersReply_Teacher{
			Name:   teacher.Name,
			Desc:   teacher.Desc,
			Avatar: teacher.Avatar,
			Title:  teacher.Title,
			Id:     int32(teacher.ID),
		})
	}
	return reply, nil
}
func (s *TeacherService) GetTeacher(ctx context.Context, req *pb.GetTeacherRequest) (*pb.GetTeacherReply, error) {
	reply := &pb.GetTeacherReply{}
	teacher, err := s.uc.GetTeacher(ctx, int(req.TeacherId))

	if err != nil {
		return reply, err
	}

	reply.Data = &pb.GetTeacherReply_Data{
		Name:   teacher.Name,
		Desc:   teacher.Desc,
		Avatar: teacher.Avatar,
		Title:  teacher.Title,
		Id:     int32(teacher.ID),
	}

	return reply, nil
}
func (s *TeacherService) DeleteTeacher(ctx context.Context, req *pb.DeleteTeacherRequest) (*pb.DeleteTeacherReply, error) {
	reply := &pb.DeleteTeacherReply{}
	teacher, err := s.uc.DeleteTeacher(ctx, req.TeacherId)

	if err != nil {
		return reply, err
	}

	reply.Data = &pb.DeleteTeacherReply_Data{
		Teacher: &pb.DeleteTeacherReply_Teacher{
			Name:   teacher.Name,
			Desc:   teacher.Desc,
			Avatar: teacher.Avatar,
			Title:  teacher.Title,
		},
	}

	return reply, nil
}
func (s *TeacherService) UpdateTeacher(ctx context.Context, req *pb.UpdateTeacherRequest) (*pb.UpdateTeacherReply, error) {
	reply := &pb.UpdateTeacherReply{}
	teacher, err := s.uc.UpdateTeacher(ctx, req)
	if err != nil {
		return reply, err
	}
	reply.Data = &pb.UpdateTeacherReply_Data{Teacher: &pb.UpdateTeacherReply_Teacher{
		TeacherId: int32(teacher.ID),
		Name:      teacher.Name,
		Desc:      teacher.Desc,
		Avatar:    teacher.Avatar,
		Title:     teacher.Title,
	}}
	return reply, nil
}
func (s *TeacherService) CreateTeacher(ctx context.Context, req *pb.CreateTeacherRequest) (*pb.CreateTeacherReply, error) {
	reply := &pb.CreateTeacherReply{}
	teacher, err := s.uc.CreateTeacher(ctx, req)
	if err != nil {
		return reply, err
	}
	reply.Data = &pb.CreateTeacherReply_Data{Teacher: &pb.CreateTeacherReply_Teacher{
		TeacherId: int32(teacher.ID),
		Name:      teacher.Name,
		Desc:      teacher.Desc,
		Avatar:    teacher.Avatar,
		Title:     teacher.Title,
	}}
	return reply, nil
}
