package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	// Mock data
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
		Phone: "099-999-9999", // Task 2: เพิ่มข้อมูลเบอร์โทรศัพท์
	}, nil
}

// Task 1: เพิ่มฟังก์ชัน ListStudents เพื่อส่งรายชื่อนักเรียนกลับไป
func (s *server) ListStudents(ctx context.Context, in *pb.Empty) (*pb.StudentListResponse, error) {
	log.Println("Received request for ListStudents")

	// สร้างข้อมูลจำลอง (Mock Data) เป็น Slice ของนักเรียน
	students := []*pb.StudentResponse{
		{
			Id:    101,
			Name:  "Alice Johnson",
			Major: "Computer Science",
			Email: "alice@university.com",
			Phone: "099-999-9999",
		},
		{
			Id:    102,
			Name:  "Bob Smith",
			Major: "Data Science",
			Email: "bob@university.com",
			Phone: "088-777-6666",
		},
		{
			Id:    103,
			Name:  "Charlie Brown",
			Major: "Software Engineering",
			Email: "charlie@university.com",
			Phone: "077-333-3333",
		},
	}

	return &pb.StudentListResponse{
		Student: students,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
