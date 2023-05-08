package main

import (
	server "github.com/cristiancll/qrpay-be/internal/api"
)

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/qrpay?sslmode=disable"
	s := server.New(dbURL)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}
