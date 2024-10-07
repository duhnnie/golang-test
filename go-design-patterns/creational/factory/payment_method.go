package factory

import (
	"fmt"
)

const Cash = 1
const DebitCard = 2

type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using cash", amount)
}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card", amount)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}
