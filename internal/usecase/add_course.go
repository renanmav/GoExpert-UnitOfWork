package usecase

import (
	"context"

	"github.com/renanmav/GoExpert-UnitOfWork/internal/entity"
	"github.com/renanmav/GoExpert-UnitOfWork/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int32
}

type AddCourseUseCase struct {
	CategoryRepository repository.CategoryRepositoryInterface
	CourseRepository   repository.CourseRepositoryInterface
}

func NewAddCourseUseCase(categoryRepository repository.CategoryRepositoryInterface, courseRepository repository.CourseRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CategoryRepository: categoryRepository,
		CourseRepository:   courseRepository,
	}
}

func (uc *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}
	err := uc.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}
	err = uc.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
