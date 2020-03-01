package core

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/hunterlong/statping/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// grpcServerDef is function type.
// Consumed by Test data.
type grpcServerDef func(int) *grpc.Server

// Test Data: Simulates testing scenarios
var testdata = []struct {
	grpcService   grpcServerDef
	clientChecker *Service
}{
	{
		grpcService: func(port int) *grpc.Server {
			return grpcServer(port, true)
		},
		clientChecker: ReturnService(&types.Service{
			Name:           "GRPC Client -> Server with Health check",
			Domain:         "localhost",
			Port:           50053,
			Expected:       types.NewNullString("status:SERVING"),
			ExpectedStatus: 1,
			Type:           "grpc",
			Timeout:        3,
			VerifySSL:      types.NewNullBool(false),
		}),
	},
	{
		grpcService: func(port int) *grpc.Server {
			return grpcServer(port, true)
		},
		clientChecker: ReturnService(&types.Service{
			Name:           "GRPC Client check TLS -> Server with Health check but without TLS",
			Domain:         "localhost",
			Port:           50054,
			Expected:       types.NewNullString(""),
			ExpectedStatus: 0,
			Type:           "grpc",
			Timeout:        1,
			VerifySSL:      types.NewNullBool(true),
		}),
	},
	{
		grpcService: func(port int) *grpc.Server {
			return grpcServer(port, false)
		},
		clientChecker: ReturnService(&types.Service{
			Name:           "GRPC Client -> Server without Health check",
			Domain:         "localhost",
			Port:           50055,
			Expected:       types.NewNullString(""),
			ExpectedStatus: 0,
			Type:           "grpc",
			Timeout:        1,
			VerifySSL:      types.NewNullBool(false),
		}),
	},
}

// grpcServer creates grpc Service with optional parameters.
func grpcServer(port int, enableHealthCheck bool) *grpc.Server {
	portString := strconv.Itoa(port)
	server := grpc.NewServer()
	lis, err := net.Listen("tcp", "localhost:"+portString)
	if err != nil {
		fmt.Println(err)
	}

	if enableHealthCheck {
		healthServer := health.NewServer()
		healthServer.SetServingStatus("Test GRPC Service", healthpb.HealthCheckResponse_SERVING)
		healthpb.RegisterHealthServer(server, healthServer)
		go server.Serve(lis)
	}
	return server
}

// TestGrpcCheck ranges over the testdata struct.
// Examines GrpcCheck() function
func TestGrpcCheck(t *testing.T) {
	for _, testscenario := range testdata {
		v := testscenario
		t.Run(v.clientChecker.Name, func(t *testing.T) {
			t.Parallel()
			server := v.grpcService(v.clientChecker.Port)
			defer server.Stop()
			v.clientChecker.checkGrpc(false)
			if v.clientChecker.LastStatusCode != v.clientChecker.ExpectedStatus || strings.TrimSpace(v.clientChecker.LastResponse) != v.clientChecker.Expected.String {
				t.Errorf("Expected message: '%v', Got message: '%v' , Expected Status: '%v', Got Status: '%v'", v.clientChecker.Expected, v.clientChecker.LastResponse, v.clientChecker.ExpectedStatus, v.clientChecker.LastStatusCode)
			}
		})
	}
}
