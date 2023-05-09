package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/handler"
	"github.com/cristiancll/qrpay-be/internal/api/middleware"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type Server struct {
	settingsPath string

	db      *pgxpool.Pool
	context context.Context
}

func (s *Server) startDatabase() error {
	// Create a new context
	s.context = context.Background()

	// Create a new connection pool
	c := configs.Get().Database
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Name)
	db, err := pgxpool.New(s.context, url)
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

func loadTLSCredentials() credentials.TransportCredentials {
	cert := &tls.Config{
		Certificates: []tls.Certificate{*configs.Get().Keys.TLS},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(cert)
}

func (s *Server) initializeAPI() error {

	// Create Repositories
	userRepo := repository.NewUserRepository(s.db)
	if err := userRepo.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate user repository: %v", err)
	}
	authRepo := repository.NewAuthRepository(s.db)
	if err := authRepo.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate auth repository: %v", err)
	}

	// Create Services
	userService := service.NewUserService(s.db, userRepo, authRepo)
	authService := service.NewAuthService(s.db, authRepo, userRepo)

	// Create Handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	creds := loadTLSCredentials()

	// Create a new gRPC server
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.RateLimiterUnaryInterceptor,
			middleware.LoggingUnaryInterceptor,
			middleware.AuthInterceptor,
		),
		grpc.ChainStreamInterceptor(
			middleware.RateLimiterStreamInterceptor,
			middleware.LoggingStreamInterceptor,
		),
	)

	// Register the Services
	proto.RegisterUserServiceServer(grpcServer, userHandler)
	proto.RegisterAuthServiceServer(grpcServer, authHandler)

	// Create a new TCP listener
	address := fmt.Sprintf(":%d", configs.Get().Server.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Start the gRPC server
	fmt.Printf("gRPC server listening at %s\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}

func New(settingsPath string) *Server {
	return &Server{settingsPath: settingsPath}
}

func (s *Server) Start() error {
	err := configs.Load(s.settingsPath)
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}
	err = s.startDatabase()
	if err != nil {
		return fmt.Errorf("could not start database: %w", err)
	}
	err = s.initializeAPI()
	if err != nil {
		return fmt.Errorf("could not start api: %w", err)
	}
	return nil
}
