package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/renanmav/GoExpert-UnitOfWork/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses?parseTime=true")
	assert.NoError(t, err)

	dtb.Exec("DROP TABLE if exists `courses`;")
	dtb.Exec("DROP TABLE if exists `categories`;")

	dtb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dtb.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Golang",
		CourseName:       "Golang Course",
		CourseCategoryID: 1, // HERE: try changing to 2
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCategoryRepository(dtb), repository.NewCourseRepository(dtb))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
