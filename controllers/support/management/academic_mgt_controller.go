package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	academicsDomain "obas/domain/academics"
	"obas/io/academics"
)

func AcademicManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", AcademiManagementHandler(app))

	return r
}

type CourseSubjectHolder struct {
	CourseId    string
	SubjectId   string
	CourseName  string
	SubjectName string
}

func AcademiManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var courseSubjectHolder []CourseSubjectHolder
		courses, err := academics.GetAllCourses()
		if err != nil {
			fmt.Println("An error in AcademiManagementHandler reading courses")
			app.ErrorLog.Println(err.Error())
		}
		subjects, errr := academics.GetSubjects()
		if errr != nil {
			fmt.Println("An error in AcademiManagementHandler reading subjects")
			app.ErrorLog.Println(errr.Error())
		}
		courseSubjects, errrr := academics.GetAllCourseSubject()
		if errrr != nil {
			fmt.Println("An error in AcademiManagementHandler reading courseSubjects")
			app.ErrorLog.Println(errrr.Error())
		} else {
			for _, myCourseSubject := range courseSubjects {
				course, err := academics.GetCourse(myCourseSubject.CourseId)
				if err != nil {
					fmt.Println("An error in AcademiManagementHandler reading cours")
					app.ErrorLog.Println(err.Error())
				}
				subject, err := academics.GetSubject(myCourseSubject.SubjectId)
				if err != nil {
					fmt.Println("An error in AcademiManagementHandler reading subject")
					app.ErrorLog.Println(err.Error())
				}
				if subject.Name != "" || course.CourseName != "" {
					myCourseSubjectHolder := CourseSubjectHolder{myCourseSubject.CourseId, myCourseSubject.SubjectId, course.CourseName, subject.Name}
					courseSubjectHolder = append(courseSubjectHolder, myCourseSubjectHolder)
				}
			}
		}

		type PageData struct {
			Subjects       []academicsDomain.Subject
			Courses        []academicsDomain.Course
			CourseSubjects []CourseSubjectHolder
			MyActiveTab    tabs
		}
		tab := app.Session.GetString(r.Context(), "tab")
		activeTab := getTabs(tab)

		Data := PageData{subjects, courses, courseSubjectHolder, activeTab}
		files := []string{
			app.Path + "content/tech/tech_admin_academics.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

type tabs struct {
	Tab1 string
	Tab2 string
}

func getTabs(tab string) tabs {

	switch tab {
	case "tab1":
		return tabs{"active show", ""}
	case "tab2":
		return tabs{"", "active show"}
	default:
		return tabs{"active show", ""}
	}
}
