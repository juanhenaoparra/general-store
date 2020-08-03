package datemodel

// Date save date object
type Date struct {
	UID   string   `json:"uid,omitempty"`
	Date  *int     `json:"timestamp,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}
