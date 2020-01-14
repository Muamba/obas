package misc

import (
	"net/http"
	"obas/config"
	"strings"
)

const (
	LayoutOBAS        = "2006-01-02"
	DangerAlertStyle  = "alert-danger"
	SuccessAlertStyle = "alert-success"
	NonMatricProgressBarIncrementValue = 4.1666666667
	MatricProgressBarIncrementVale = 5.882353
)

type PageToast struct {
	AlertType string
	AlertInfo string
}

/**
Get date in YYYYMMDD format
*/
func GetDate_YYYYMMDD(dateString string) string {
	return strings.Split(dateString, " ")[0]
}

/**
Check if session has alert message
*/
func CheckForSessionAlert(app *config.Env, r *http.Request) PageToast {
	message := app.Session.GetString(r.Context(), "message")
	messageType := app.Session.GetString(r.Context(), "message-type")
	var alert PageToast
	if message != "" && messageType != "" {
		alert = PageToast{messageType, message}
		app.Session.Remove(r.Context(), "message")
		app.Session.Remove(r.Context(), "message-type")
	}
	return alert
}

/**
Set session alert message
 */
func SetSessionMessage(app *config.Env, r *http.Request, messageType string, message string) {
	app.Session.Put(r.Context(), "message-type", messageType)
	app.Session.Put(r.Context(), "message", message)
}
