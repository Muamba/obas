package academics

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	academicsDomain "obas/domain/academics"
	academicsIO "obas/io/academics"
)

/**
Get subjects for course given course id
*/
func GetSubjectsForCourse(app *config.Env, courseId string) ([]academicsDomain.CourseSubject, genericHelper.PageToast) {
	var courseSubjects []academicsDomain.CourseSubject
	var alert genericHelper.PageToast
	courseSubjects, err := academicsIO.GetCourseSubjects(courseId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve subjects for course!"}
	}
	return courseSubjects, alert
}

/**
Get Subject details given subject id
*/
func GetSubject(app *config.Env, subjectId string) (academicsDomain.Subject, genericHelper.PageToast) {
	var subject academicsDomain.Subject
	var alert genericHelper.PageToast
	subject, err := academicsIO.GetSubject(subjectId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve subject!"}
	}
	return subject, alert
}

/**
Get course name given course id
 */
func GetCourseName(app *config.Env, courseId string) (string, genericHelper.PageToast) {
	var courseName string
	var alert genericHelper.PageToast
	course, alert := GetCourse(app, courseId)
	if alert.AlertInfo == "" {
		courseName = course.CourseName
	}
	return courseName, alert

}

/**
Get course details given course id
 */
func GetCourse(app *config.Env, courseId string) (academicsDomain.Course, genericHelper.PageToast) {
	var course academicsDomain.Course
	var alert genericHelper.PageToast
	course, err := academicsIO.GetCourse(courseId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve course!"}
	}
	return course, alert
}
