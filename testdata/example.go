package testdata

import "time"

// VENDOR implements json structure for vendor
type VENDOR struct {
	Url      string    `json:"url,omitempty"`
	Added    time.Time `json:"added,omitempty"`
	Priority int       `json:"priority,omitempty"`
	Name     string    `json:"name,omitempty"`
}

// EXAMPLE implements json structure for example
type EXAMPLE struct {
	Timestamp time.Time              `json:"timestamp,omitempty"`
	Children  map[string]interface{} `json:"children,omitempty"`
	Vendor    *VENDOR                `json:"vendor,omitempty"`
	Nullvalue interface{}            `json:"nullvalue,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Number    int                    `json:"number,omitempty"`
	Weight    float64                `json:"weight,omitempty"`
}
