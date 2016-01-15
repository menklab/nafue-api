package config

import (
	"os"
)

var (
	S3Bucket		   = "menklab.com"
	S3Key			   = "sfds"
	S3PutTTL		   = 15
//	S3Bucket		   = os.Getenv("S3_BUCKET")
	S3Location		   = "us-east-1"
//	S3Location		   = os.Getenv("S3_LOCATION")
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
