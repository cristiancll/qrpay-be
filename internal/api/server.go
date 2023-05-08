package server

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/handler"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	dbURL string

	db      *pgxpool.Pool
	context context.Context
}

func (s *Server) startDatabase() error {
	// Create a new context
	s.context = context.Background()

	// Create a new connection pool
	db, err := pgxpool.New(s.context, s.dbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	// Ping the database to check if it's still alive.
	err = db.Ping(s.context)
	if err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}

	// Set the database connection pool to the Server struct
	s.db = db
	return nil
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func (s *Server) initializeAPI() error {
	// Create a new instance of the UserRepository
	userRepo := repository.NewUserRepository(s.db)
	if err := userRepo.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate user repository: %v", err)
	}
	authRepo := repository.NewAuthRepository(s.db)
	if err := authRepo.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate auth repository: %v", err)
	}

	// Create a new instance of the UserService, passing in the UserRepository
	userService := service.NewUserService(s.db, userRepo, authRepo)

	// Create a new instance of the UserHandler, passing in the UserService
	userHandler := handler.NewUserHandler(userService)

	// Create a new instance of the grpc.Server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	// Register the UserHandler with the grpc.Server
	proto.RegisterUserServiceServer(grpcServer, userHandler)

	// Create a new TCP listener on port 50051
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Println("Starting server on port 9090...")
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}

func New(dbURL string) *Server {
	return &Server{dbURL: dbURL}
}

func (s *Server) Start() error {
	err := s.startDatabase()
	if err != nil {
		return fmt.Errorf("could not start database: %w", err)
	}
	err = s.initializeAPI()
	if err != nil {
		return fmt.Errorf("could not start api: %w", err)
	}
	return nil
}
