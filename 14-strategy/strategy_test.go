package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	payment := NewPayment("Ada", "0001", 123, &Cash{})
	payment.Pay()
	payment = NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
