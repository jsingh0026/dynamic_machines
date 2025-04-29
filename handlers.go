package main

import (
	"context"
	"fmt"
	"time"

	pb "machines/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var store = NewMachineStore()

type MachineServiceServer struct {
	pb.UnimplementedMachineServiceServer
}

func (s *MachineServiceServer) StartMachine(ctx context.Context, req *pb.StartMachineRequest) (*pb.MachineResponse, error) {
	if req.Id == "" || req.Host == "" {
		return nil, status.Errorf(codes.InvalidArgument, "both 'id' and 'host' are required")
	}

	m := store.StartMachine(req.Id, req.Host)
	return &pb.MachineResponse{
		Id:      m.ID,
		Host:    m.Host,
		Running: m.Running,
		Started: m.Started.Format(time.RFC3339),
	}, nil
}

func (s *MachineServiceServer) StopMachine(ctx context.Context, req *pb.StopMachineRequest) (*pb.StopMachineResponse, error) {
	success := store.StopMachine(req.Id)
	return &pb.StopMachineResponse{Success: success}, nil
}

func (s *MachineServiceServer) CloneMachine(ctx context.Context, req *pb.CloneMachineRequest) (*pb.MachineResponse, error) {
	m := store.CloneMachine(req.OldId, req.NewId, req.NewHost)
	fmt.Printf("OldId: %q, NewId: %q, NewHost: %q\n", req.OldId, req.NewId, req.NewHost)
	if m == nil {
		return nil, status.Error(codes.NotFound, "original machine not found or not running")
	}
	return &pb.MachineResponse{
		Id:      m.ID,
		Host:    m.Host,
		Running: m.Running,
		Started: m.Started.Format(time.RFC3339),
	}, nil
}

func (s *MachineServiceServer) ListMachines(ctx context.Context, _ *pb.ListMachinesRequest) (*pb.ListMachinesResponse, error) {
	machines := store.ListMachines()
	var res []*pb.MachineResponse
	for _, m := range machines {
		res = append(res, &pb.MachineResponse{
			Id:      m.ID,
			Host:    m.Host,
			Running: m.Running,
			Started: m.Started.Format(time.RFC3339),
		})
	}
	return &pb.ListMachinesResponse{Machines: res}, nil
}
