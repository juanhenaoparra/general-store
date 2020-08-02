package device

// Getter interface
type Getter interface {
	GetAll() []Device
}

// Adder interface
type Adder interface {
	Add(device Device)
}

// Device Type
type Device struct {
	UID   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

// Repo for Device Type
type Repo struct {
	Devices []Device
}

// New Device Struct
func New() *Repo {
	return &Repo{
		Devices: []Device{},
	}
}

// Add a new Device Struct
func (r *Repo) Add(device Device) {
	r.Devices = append(r.Devices, device)
}

// GetAll of the  Devices
func (r *Repo) GetAll() []Device {
	return r.Devices
}
