package ip

// Getter interface
type Getter interface {
	GetAll() []IP
}

// Adder interface
type Adder interface {
	Add(ip IP)
}

// IP Type
type IP struct {
	UID     string   `json:"uid,omitempty"`
	Address string   `json:"address,omitempty"`
	DType   []string `json:"dgraph.type,omitempty"`
}

// Repo for IP Type
type Repo struct {
	IPS []IP `json:"all"`
}

// New IP Struct
func New() *Repo {
	return &Repo{
		IPS: []IP{},
	}
}

// Add a new IP Struct
func (r *Repo) Add(ip IP) {
	r.IPS = append(r.IPS, ip)
}

// GetAll of the  IPS
func (r *Repo) GetAll() []IP {
	return r.IPS
}
