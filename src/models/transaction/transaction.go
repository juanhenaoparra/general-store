package transaction

// Getter interface
type Getter interface {
	GetAll() []Transaction
}

// Adder interface
type Adder interface {
	Add(transaction Transaction)
}

// Transaction Type
type Transaction struct {
	UID      string   `json:"uid,omitempty"`
	ID       string   `json:"id,omitempty"`
	Date     *int     `json:"date,omitempty"`
	BuyerID  string   `json:"by_buyer,omitempty"`
	IP       string   `json:"since_ip,omitempty"`
	Device   string   `json:"since_device,omitempty"`
	Products []string `json:"have_products"`
	DType    []string `json:"dgraph.type,omitempty"`
}

// Repo for Transaction Type
type Repo struct {
	Transactions []Transaction `json:"all"`
}

// New Transaction Struct
func New() *Repo {
	return &Repo{
		Transactions: []Transaction{},
	}
}

// Add a new Transaction Struct
func (r *Repo) Add(transaction Transaction) {
	r.Transactions = append(r.Transactions, transaction)
}

// GetAll of the  Transactions
func (r *Repo) GetAll() []Transaction {
	return r.Transactions
}
