package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 8575
// Hash 6380
// Hash 2168
// Hash 8931
// Hash 3256
// Hash 6929
// Hash 9870
// Hash 3006
// Hash 3767
// Hash 8672
// Hash 9680
// Hash 7314
// Hash 6746
// Hash 2282
// Hash 5306
// Hash 3619
// Hash 2170
// Hash 6518
// Hash 2079
// Hash 8620
// Hash 3811
// Hash 2249
// Hash 8574
// Hash 7142
// Hash 8721
// Hash 5678
// Hash 8542