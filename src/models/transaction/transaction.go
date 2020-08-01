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
	ID      string `json:"id"`
	BuyerID string `json:"buyer_id"`
	IP      string `json:"ip"`
	Device  string `json:"device"`
	// Date     int
	Products []string `json:"product_ids"`
}

// Repo for Transaction Type
type Repo struct {
	Date         int           `json:"date"`
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
