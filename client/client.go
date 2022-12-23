package client

import (
	"context"
	"time"

	"github.com/alramdein/auth-service/pb"

	grpcpool "github.com/processout/grpc-go-pool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type client struct {
	Conn *grpcpool.Pool
}

func NewClient(target string, timeout time.Duration, idleConnPool, maxConnPool int) (pb.AuthServiceClient, error) {
	factory := newFactory(target, timeout)

	pool, err := grpcpool.New(factory, idleConnPool, maxConnPool, time.Second)
	if err != nil {
		log.Errorf("Error : %v", err)
		return nil, err
	}

	return &client{
		Conn: pool,
	}, nil
}

func newFactory(target string, timeout time.Duration) grpcpool.Factory {
	return func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(target, grpc.WithInsecure(), withClientUnaryInterceptor(timeout))
		if err != nil {
			log.Errorf("Error : %v", err)
			return nil, err
		}

		return conn, err
	}
}

func withClientUnaryInterceptor(timeout time.Duration) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	})
}

func (u *client) HasAccess(ctx context.Context, req *pb.HassAccessRequest, opts ...grpc.CallOption) (*pb.HasAccessResponse, error) {
	conn, err := u.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewAuthServiceClient(conn.ClientConn)
	return cli.HasAccess(ctx, req, opts...)
}
