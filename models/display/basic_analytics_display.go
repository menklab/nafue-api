package display

import "fmt"

type BasicAnalyticsDisplay struct {
	FileCount int `json:"fileCount,omitempty"`
}

func (self *BasicAnalyticsDisplay) ToString() string {
	return fmt.Sprintf(
		"{FileCount: %v}",
		self.FileCount,
	)
}
