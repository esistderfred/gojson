package testdata

import "time"

// CHILDREN implements json structure for children
type CHILDREN struct {
}

// VENDOR implements json structure for vendor
type VENDOR struct {
	Priority int       `json:"priority,omitempty"`
	Name     string    `json:"name,omitempty"`
	Url      string    `json:"url,omitempty"`
	Added    time.Time `json:"added,omitempty"`
}

// EXAMPLE implements json structure for example
type EXAMPLE struct {
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Children  *CHILDREN   `json:"children,omitempty"`
	Vendor    *VENDOR     `json:"vendor,omitempty"`
	Nullvalue interface{} `json:"nullvalue,omitempty"`
	Name      string      `json:"name,omitempty"`
	Number    int         `json:"number,omitempty"`
	Weight    float64     `json:"weight,omitempty"`
}
