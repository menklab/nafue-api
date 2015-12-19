package config

import (
	"os"
)

var (
	//path to media assets (typically in a s3 bucket)
	//Database and Redis Connections
	DBUrl              = ""
	DBName             = os.Getenv("DBNAME")
	RedisDb            = os.Getenv("REDIS_DB")
	RedisAuth          = os.Getenv("REDIS_AUTH")
	HubspotKey         = os.Getenv("HUBSPOT_KEY")
	HubSpotContentApi  = os.Getenv("HUBSPOT_CONTENT_API")
	HubSpotFormApi     = os.Getenv("HUBSPOT_FORM_API")
	HubSpotPostDataApi = os.Getenv("HUBSPOT_POST_DATA_API")
	RFQGuid            = os.Getenv("RFQ_GUID")
	RQFPortal          = os.Getenv("RFQ_PORTAL")
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
