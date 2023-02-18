package main

import (
	"database/sql"
	"net"

	"github.com/fernandohtr/grpc-case-study/internal/database"
	"github.com/fernandohtr/grpc-case-study/internal/pb"
	"github.com/fernandohtr/grpc-case-study/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, error := sql.Open("sqlite3", "./db.sqlite")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, error := net.Listen("tcp", ":50051")
	if error != nil {
		panic(error)
	}

	if error := grpcServer.Serve(listener); error != nil {
		panic(error)
	}
}
