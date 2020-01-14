package model

// Customer is the model of a person/business entity who receives credits and makes payments
type Customer struct {
	ID            int           `db:"id" json:"id"`
	Name          string        `db:"name" json:"name"`
	ShortName     string        `db:"short_name" json:"shortname"`
	DeliveryRoute string        `db:"delivery_route" json:"deliveryroute"`
	Contact       []int         `db:"contact" json:"contact"`
	Credits       []Transaction `db:"credits" json:"-"`
	Payments      []Transaction `db:"payments" json:"-"`
	DueAmount     float64       `db:"due_amount" json:"dueamount"`
}

// CalculateDueAmount is a helper method to calculate the total due amount for a customer
func (customer *Customer) CalculateDueAmount() {
	sumCredits := float64(0)
	for _, c := range customer.Credits {
		sumCredits += c.Amount
	}

	sumPayments := float64(0)
	for _, p := range customer.Payments {
		sumPayments += p.Amount
	}

	customer.DueAmount = sumCredits - sumPayments
}
