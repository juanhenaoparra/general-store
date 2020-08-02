package product

// Getter interface
type Getter interface {
	GetAll() []Product
}

// Adder interface
type Adder interface {
	Add(product Product)
}

// Product Type
type Product struct {
	UID   string   `json:"uid,omitempty"`
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price int      `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

// Repo for Product Type
type Repo struct {
	Products []Product
}

// New Product Struct
func New() *Repo {
	return &Repo{
		Products: []Product{},
	}
}

// Add a new Product Struct
func (r *Repo) Add(product Product) {
	r.Products = append(r.Products, product)
}

// GetAll of the  Products
func (r *Repo) GetAll() []Product {
	return r.Products
}
