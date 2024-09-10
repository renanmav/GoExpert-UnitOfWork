package usecase

import (
	"context"

	"github.com/renanmav/GoExpert-UnitOfWork/internal/entity"
	"github.com/renanmav/GoExpert-UnitOfWork/internal/repository"
	"github.com/renanmav/GoExpert-UnitOfWork/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int32
}

type AddCourseUseCaseUow struct {
	uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		uow: uow,
	}
}

func (uc *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCaseUow) error {
	return uc.uow.Do(ctx, func(uow uow.UowInterface) error {
		// Everything inside this function is a unit of work
		// If an error occurs, the transaction will be rolled back
		// If everything is ok, the transaction will be committed

		categoryRepository := uc.getCategoryRepository(ctx)
		courseRepository := uc.getCourseRepository(ctx)

		// Create category
		category := entity.Category{
			Name: input.CategoryName,
		}
		err := categoryRepository.Insert(ctx, category)
		if err != nil {
			return err
		}

		// Create course
		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		err = courseRepository.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uc *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := uc.uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface)
}

func (uc *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := uc.uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface)
}
