package main

import "fmt"

// PaymentStrategy - интерфейс стратегии оплаты.
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CardPayment - оплата банковской картой.
type CardPayment struct {
	CardNumber string
}

func (c *CardPayment) Pay(amount float64) string {
	last4 := c.CardNumber
	if len(last4) > 4 {
		last4 = last4[len(last4)-4:]
	}
	return fmt.Sprintf("Оплата %.2f руб. банковской картой ****%s выполнена успешно", amount, last4)
}

// CashPayment - оплата наличными курьеру.
type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) string {
	return fmt.Sprintf("Заказ оформлен на сумму %.2f руб. с оплатой наличными при получении", amount)
}

// SBPPayment - оплата через СБП.
type SBPPayment struct {
	Phone string
}

func (s *SBPPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата %.2f руб. через СБП подтверждена для номера %s", amount, s.Phone)
}

// OrderProcessor - контекст, использующий выбранную стратегию.
type OrderProcessor struct {
	strategy PaymentStrategy
	orderID  int
	total    float64
}

func NewOrderProcessor(orderID int, total float64) *OrderProcessor {
	return &OrderProcessor{orderID: orderID, total: total}
}

func (o *OrderProcessor) SetStrategy(strategy PaymentStrategy) {
	o.strategy = strategy
}

func (o *OrderProcessor) Checkout() string {
	if o.strategy == nil {
		return "Стратегия оплаты не выбрана"
	}
	return fmt.Sprintf("Заказ #%d: %s", o.orderID, o.strategy.Pay(o.total))
}

func main() {
	order := NewOrderProcessor(101, 1590.50)

	order.SetStrategy(&CardPayment{CardNumber: "2200123412345678"})
	fmt.Println(order.Checkout())

	order.SetStrategy(&CashPayment{})
	fmt.Println(order.Checkout())

	order.SetStrategy(&SBPPayment{Phone: "+7-900-123-45-67"})
	fmt.Println(order.Checkout())
}
