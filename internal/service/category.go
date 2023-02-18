package service

import (
	"github.com/fernandohtr/grpc-case-study/internal/pb"
	"github.com/fernandohtr/grpc-case-study/internal/repository/database"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Categoty
}

func NewCategoryService(categoryDB database.Categoty) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(context context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, error := c.CategoryDB.Create(input.Name, input.Description)
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Error to create category: %v", error)
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
