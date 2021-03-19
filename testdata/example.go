package testdata

import "time"

// VENDOR implements json structure for vendor
type VENDOR struct {
	Name     string    `json:"name,omitempty"`
	Url      string    `json:"url,omitempty"`
	Added    time.Time `json:"added,omitempty"`
	Priority int       `json:"priority,omitempty"`
}

// EXAMPLE implements json structure for example
type EXAMPLE struct {
	Number           int                    `json:"number,omitempty"`
	Weight           float64                `json:"weight,omitempty"`
	AlternativeNames []interface{}          `json:"alternative_names,omitempty"`
	Order            []interface{}          `json:"order,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Timestamp        time.Time              `json:"timestamp,omitempty"`
	Children         map[string]interface{} `json:"children,omitempty"`
	Vendor           *VENDOR                `json:"vendor,omitempty"`
	Nullvalue        interface{}            `json:"nullvalue,omitempty"`
}
