package domain

type Course struct {
	Id         string `json:"id"`
	CourseName string `json:"courseName"`
	CourseDesc string `json:"courseDesc"`
}

type CourseSubject struct {
	CourseId  string `json:"courseId"`
	SubjectId string `json:"subjectId"`
}
type Subject struct {
	Id          string `json:"id"`
	Name        string `json:"subjectName"`
	SubjectDesc string `json:"subjectDesc"`
}
