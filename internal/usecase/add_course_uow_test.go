package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/renanmav/GoExpert-UnitOfWork/internal/db"
	"github.com/renanmav/GoExpert-UnitOfWork/internal/repository"
	"github.com/renanmav/GoExpert-UnitOfWork/pkg/uow"
	"github.com/stretchr/testify/assert"
)

func TestAddCourseUow(t *testing.T) {
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses?parseTime=true")
	assert.NoError(t, err)

	dtb.Exec("DROP TABLE if exists `courses`;")
	dtb.Exec("DROP TABLE if exists `categories`;")

	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dtb.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()

	uow := uow.NewUow(dtb)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dtb)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dtb)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCaseUow{
		CategoryName:     "Golang",
		CourseName:       "Golang Course",
		CourseCategoryID: 1, // HERE: try changing to 2
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
