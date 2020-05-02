package server

import (
	"context"
	"net"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	pb "github.com/schmurfy/concourse-test/service/generated_pb/service"
)

type service struct {
	logger *zap.Logger
	redis  *redis.Client
}

func (s *service) GetAddresses(context.Context, *empty.Empty) (*pb.GetAddressesResponse, error) {
	ret := &pb.GetAddressesResponse{}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ret.Ipv4S = append(ret.Ipv4S, v.String())
			case *net.IPAddr:
				ret.Ipv4S = append(ret.Ipv4S, v.String())
			}
		}
	}

	return ret, nil
}

func (s *service) IncrementRedis(ctx context.Context, req *pb.IncrementRedisRequest) (*empty.Empty, error) {
	time.Sleep(5 * time.Second)

	val, err := s.redis.Get(req.Key).Result()
	if err != nil {
		s.logger.Error("failed to get key",
			zap.Error(err),
			zap.String("host", "redis"),
		)

		// initialize the key
		err = s.redis.Set(req.Key, 0, 0).Err()
		if err != nil {
			s.logger.Fatal("failed to iniialize key",
				zap.Error(err),
			)
		}
	}

	s.logger.Info("increment",
		zap.String("value", val),
	)

	err = s.redis.Incr(req.Key).Err()
	if err != nil {
		s.logger.Error("failed to increment key",
			zap.Error(err),
		)
	}

	return &empty.Empty{}, nil
}

func Start(logger *zap.Logger, redis *redis.Client) {
	grpcServer := grpc.NewServer()

	svc := &service{
		logger: logger,
		redis:  redis,
	}

	pb.RegisterServiceServer(grpcServer, svc)
	reflection.Register(grpcServer)

	srv := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, srv)

	conn, err := net.Listen("tcp", ":8000")
	if err != nil {
		logger.Fatal("error while intializing grpc server",
			zap.Error(err),
			zap.String("listen_address", ":8080"),
		)
	}

	grpcServer.Serve(conn)
}
