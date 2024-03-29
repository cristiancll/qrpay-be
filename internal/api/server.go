package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/middleware"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type Server struct {
	settingsPath string

	db      *pgxpool.Pool
	context context.Context

	repos    *repositories
	handlers *handlers
	services *services
}

func (s *Server) createDatabaseIfNotExists(err error) error {
	err = errors.Unwrap(err)
	if er, ok := err.(*pgconn.PgError); ok {
		if er.Code == "3D000" {
			c := configs.Get().Database
			url := fmt.Sprintf("postgres://%s:%s@%s:%d/?sslmode=disable", c.Username, c.Password, c.Host, c.Port)
			db, err := pgxpool.New(s.context, url)
			if err != nil {
				return errs.New(err, errCode.Internal)
			}
			defer db.Close()
			query := fmt.Sprintf("CREATE DATABASE %s", c.Name)
			_, err = db.Exec(s.context, query)
			if err != nil {
				return errs.New(err, errCode.Internal)
			}
		}
	}
	return nil
}

func (s *Server) startDatabase() error {
	// Create a new context
	s.context = context.Background()

	// Create a new connection pool
	c := configs.Get().Database
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Name)
	db, err := pgxpool.New(s.context, url)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}

	// Ping the database to check if it's still alive.
	err = db.Ping(s.context)
	if err != nil {
		err = s.createDatabaseIfNotExists(err)
		if err == nil {
			return s.startDatabase()
		}
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
	err := s.createRepositories()
	if err != nil {
		return err
	}

	// Create Services
	s.createServices()

	// Create Handlers
	s.createHandlers()

	// Load TLS keys for HTTPS
	creds := loadTLSCredentials()

	// Create a new gRPC server
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			middleware.RateLimiterUnaryInterceptor,
			middleware.LoggingUnaryInterceptor,
			middleware.AuthInterceptor,
			middleware.ErrorUnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			middleware.RateLimiterStreamInterceptor,
			middleware.LoggingStreamInterceptor,
			middleware.ErrorStreamInterceptor,
		),
	)

	// Register the Services
	s.registerServices(grpcServer)

	// Create a new TCP listener
	address := fmt.Sprintf(":%d", configs.Get().Server.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}

	// Start the gRPC server
	fmt.Printf("gRPC server listening at %s\n", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		return errs.New(err, errCode.Internal)
	}

	return nil
}

func New(settingsPath string) *Server {
	return &Server{
		settingsPath: settingsPath,
		repos:        &repositories{},
		handlers:     &handlers{},
		services:     &services{},
	}
}

func (s *Server) Start() error {
	err := configs.Load(s.settingsPath)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	err = s.startDatabase()
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	err = s.initializeAPI()
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}
