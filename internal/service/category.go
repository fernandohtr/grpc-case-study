package service

import (
	"github.com/fernandohtr/grpc-case-study/internal/database"
	"github.com/fernandohtr/grpc-case-study/internal/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(context context.Context, input *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, error := c.CategoryDB.Create(input.Name, input.Description)
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Error to create category: %v", error)
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategory(context context.Context, input *pb.Blank) (*pb.CategoryList, error) {
	categories, error := c.CategoryDB.FindAll()
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Error to list categories: %v", error)
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(context context.Context, input *pb.CategoryGetRequest) (*pb.Category, error) {
	category, error := c.CategoryDB.Find(input.Id)
	if error != nil {
		return nil, status.Errorf(codes.Internal, "Error to get category: %v", error)
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}
