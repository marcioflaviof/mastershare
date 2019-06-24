package models

type Table struct {
	Master   string    `json: "master"`
	Users    []User    `json: "users"`
	Products []Product `json: "products"`
	//how much they will pay
	Account float32 `json:"account"`
}

type SecureTable struct {
	Master string `json:"master"`
}

func (t *Table) TotalValue() (p float64) {

	for _, product := range t.Products {
		p += product.Price
	}

	return
}

func (t *Table) ShareAllBills(b float64) {
	
	for i, _ := range t.Users {
		t.Users[i].Bill = b
	}

}
