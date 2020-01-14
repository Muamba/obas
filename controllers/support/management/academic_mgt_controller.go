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
	r.Get("/delete/subject/{resetKey}", DeleteSubjectManagementHandler(app))
	r.Get("/delete/course/{courseId}", DeletecourseManagementHandler(app))
	r.Get("/delete/coursesubject/{courseId}/{subjectId}", DeletecourseSubjectManagementHandler(app))

	r.Post("/create/course", CreateCourseManagementHandler(app))
	r.Post("/create/subject", CreateSubjectManagementHandler(app))
	r.Post("/create/courseSubject", CreateCourseSubjectManagementHandler(app))

	return r
}

func DeletecourseSubjectManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseId := chi.URLParam(r, "courseId")
		SubjectId := chi.URLParam(r, "subjectId")
		fmt.Println("courseId", courseId)
		fmt.Println("SubjectId", SubjectId)
		_ = app.Session.Destroy(r.Context())
		courseSubjectObject := academicsDomain.CourseSubject{courseId, SubjectId}
		if courseSubjectObject.SubjectId != "" {
			_, err := academics.DeleteCourseSubject(courseSubjectObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		fmt.Println("course subject to detele>>>", courseSubjectObject)
		app.Session.Put(r.Context(), "tab", "tab3")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func DeletecourseManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseId := chi.URLParam(r, "courseId")
		_ = app.Session.Destroy(r.Context())
		courseObject, err := academics.GetCourse(courseId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if courseObject.Id != "" {
			_, err := academics.DeleteCourse(courseObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func DeleteSubjectManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subjectId := chi.URLParam(r, "resetKey")
		_ = app.Session.Destroy(r.Context())
		fmt.Println("subject Id to delete>>>>", subjectId)
		subjectObject, err := academics.GetSubject(subjectId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if subjectObject.Id != "" {
			_, err := academics.DeleteSubject(subjectObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func CreateCourseSubjectManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		_ = app.Session.Destroy(r.Context())
		subjectId := r.PostFormValue("subjectId")
		courseId := r.PostFormValue("courseId")

		fmt.Println(subjectId, "<<<< subjectId||courseId>>>>", courseId)
		if subjectId != "" || courseId != "" {
			newcCourseSubject := academicsDomain.CourseSubject{courseId, subjectId}
			_, err := academics.CreateCourseSubject(newcCourseSubject)
			if err != nil {
				fmt.Println("An error in CreateCourseManagementHandler create myCourse")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func CreateSubjectManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		_ = app.Session.Destroy(r.Context())
		subjectName := r.PostFormValue("subjectName")
		subjectDesc := r.PostFormValue("Description")

		if subjectName != "" || subjectDesc != "" {
			newcCourse := academicsDomain.Subject{"", subjectName, subjectDesc}
			_, err := academics.CreateSubject(newcCourse)
			if err != nil {
				fmt.Println("An error in CreateCourseManagementHandler create myCourse")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
}

func CreateCourseManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		_ = app.Session.Destroy(r.Context())
		courseName := r.PostFormValue("courseName")
		courseDescription := r.PostFormValue("courseDescription")

		if courseName != "" || courseDescription != "" {
			newcCourse := academicsDomain.Course{"", courseName, courseDescription}
			_, err := academics.SaveCourse(newcCourse)
			if err != nil {
				fmt.Println("An error in CreateCourseManagementHandler create myCourse")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/academics", 301)
	}
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
		fmt.Println("All the courseSubjects", courseSubjects)
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
					fmt.Println(myCourseSubject.CourseId, "<<<CourseId,  SubjectId>>>>>", myCourseSubject.SubjectId)
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
	Tab3 string
}

func getTabs(tab string) tabs {

	switch tab {
	case "tab1":
		return tabs{"active show", "", ""}
	case "tab2":
		return tabs{"", "active show", ""}
	case "tab3":
		return tabs{"", "", "active show"}
	default:
		return tabs{"active show", "", ""}
	}
}
