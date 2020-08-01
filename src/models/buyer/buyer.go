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
	// Date int    `json:"date"`
}

// Repo for Buyer Type
type Repo struct {
	Date   int     `json:"date"`
	Buyers []Buyer `json:"all"`
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

//SetDateAll for every buyer
// func (r *Repo) SetDateAll() {
// 	for i := range r.Buyers {
// 		r.Buyers[i].Date = r.Date
// 	}
// }
