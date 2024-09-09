package entity

type Category struct {
	ID        int32
	Name      string
	CourseIDs []int32
}

func (c *Category) AddCourse(courseID int32) {
	c.CourseIDs = append(c.CourseIDs, courseID)
}

type Course struct {
	ID         int32
	Name       string
	CategoryID int32
}
