package config

import (
	"os"
)

var (
	S3Key = os.Getenv("S3_KEY")
	S3PutTTL = os.Getenv("S3_PUT_TTL")
	S3Bucket = os.Getenv("S3_BUCKET")
	S3Location = os.Getenv("S3_LOCATION")
	PORT = os.Getenv("PORT")
)

const (
	PathSeperator string = string(os.PathSeparator)

//path to customer facing views to be compiled at app start
	ViewPath string = "" + PathSeperator + "" + PathSeperator

//directory containing css, js, and img files for public site
	PublicDir string = "" + PathSeperator

	AdminPath string = "" + PathSeperator + "" + PathSeperator

//name of the site as it shows up right of the pipe in the page title
	SiteName string = ""
)
