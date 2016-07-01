package models

type BasicAnalytics struct {
	FileCount int `json:"fileCount,omitempty" db:"fileCount"`
}
