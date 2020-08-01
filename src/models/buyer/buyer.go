package buyer

// Getter interface
type Getter interface {
	GetAll() []Buyer
}

// Adder interface
type Adder interface {
	Add(buyer Buyer)
}

// Buyer Type
type Buyer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Date int
}

// Repo for Buyer Type
type Repo struct {
	Buyers []Buyer
}

// New Buyer Struct
func New() *Repo {
	return &Repo{
		Buyers: []Buyer{},
	}
}

// Add a new Buyer Struct
func (r *Repo) Add(buyer Buyer) {
	r.Buyers = append(r.Buyers, buyer)
}

// GetAll of the  Buyers
func (r *Repo) GetAll() []Buyer {
	return r.Buyers
}
