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
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	// Date  int
}

// Repo for Product Type
type Repo struct {
	Date     int `json:"date"`
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
