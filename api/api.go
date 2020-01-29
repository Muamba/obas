package api

import (
	"gopkg.in/resty.v1"
	"obas/config"
)

//const BASE_URL string = "https://ict.cput.ac.za/obasapi"
const BASE_URL string = "http://155.238.32.219:9000"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
