# gojson

Create Json-Interface structures from sample json


## Installation
```
go install
```

## Example

From given input ```testdata/example.json```:
```json
{
  "name": "Entitiy",
  "number": 3,
  "weight": 4.3,
  "timestamp": "2019-11-12T14:08:00Z",
  "children": {},
  "vendor": {
    "name": "Supervendor",
    "url": "http://www.vendor.com",
    "added": "2019-11-12T14:08:00Z",
    "priority": 12
  },
  "nullvalue": null
}
```

Run the command
```
cd testdata
gojson example.json
```

The program will create ```example.go```:
```golang
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
```