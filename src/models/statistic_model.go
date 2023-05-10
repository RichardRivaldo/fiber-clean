package models

type Statistics struct {
	UserCount       int64 `json:"user_count"`
	CourseCount     int64 `json:"course_count"`
	FreeCourseCount int64 `json:"free_course_count"`
}
