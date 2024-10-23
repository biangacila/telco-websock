package grpc

import (
	"context"
	"encoding/json"
	"github.com/biangacila/telco-websock/application/services"
	"github.com/biangacila/telco-websock/domain/valueobjects"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Define the server that implements the gRPC service
type server struct {
	UnimplementedDashboardServiceServer
	dashboardService *services.DashboardService
}

// Implement the UpdateDashboard method
func (s *server) UpdateDashboard(ctx context.Context, req *DashboardRequest) (*DashboardResponse, error) {
	log.Printf("Received request to update dashboard for user: %s with data: %s", req.GetUserCode(), req.GetData())

	// Here you can update the dashboard using your existing WebSocket manager
	userCode := valueobjects.UserCode{Code: req.GetUserCode()}
	var payload = make(map[string]interface{})
	_ = json.Unmarshal([]byte(req.GetData()), &payload)
	s.dashboardService.UpdateDashboardInfo(userCode, payload)
	//wsManager.Broadcast(req.GetUserCode(), req.GetData())

	return &DashboardResponse{
		Message: "Dashboard updated successfully",
	}, nil
}

func StartGRPCServer(dashboardService services.DashboardService) {
	lis, err := net.Listen("tcp", ":50051") // gRPC server listens on port 50051
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	RegisterDashboardServiceServer(grpcServer, &server{
		dashboardService: &dashboardService,
	})

	// Enable reflection
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051 ...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
