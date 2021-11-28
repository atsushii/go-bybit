package bybit

type NewOrderService struct {
	Client         *client
	Symbol         string
	Side           bool
	OrderType      string
	Quantity       string
	StopPrice      int
	TimeInForce    string
	BasePrice      int
	Price          int
	TriggerBy      bool
	CloseOnTrigger bool
	OrderLinkId    string
}

// set Symbol
func (s *NewOrderService) symbol(symbol string) *NewOrderService {
	s.Symbol = symbol
	return s
}

// set Side
func (s *NewOrderService) side(side bool) *NewOrderService {
	s.Side = side
	return s
}

// set OrderType
func (s *NewOrderService) orderType(orderType string) *NewOrderService {
	s.OrderType = orderType
	return s
}

// set Quantity
func (s *NewOrderService) quantity(quantity string) *NewOrderService {
	s.Quantity = quantity
	return s
}

// set StopPrice
func (s *NewOrderService) stopPrice(stopPrice int) *NewOrderService {
	s.StopPrice = stopPrice
	return s
}

// set TimeInForce
func (s *NewOrderService) timeInForce(timeInForce string) *NewOrderService {
	s.TimeInForce = timeInForce
	return s
}

// set BasePrice
func (s *NewOrderService) basePrice(basePrice int) *NewOrderService {
	s.BasePrice = basePrice
	return s
}

// set Price
func (s *NewOrderService) price(price int) *NewOrderService {
	s.Price = price
	return s
}

// set TriggerBy
func (s *NewOrderService) triggerBy(triggerBy bool) *NewOrderService {
	s.TriggerBy = triggerBy
	return s
}

// set CloseOnTrigger
func (s *NewOrderService) closeOnTrigger(closeOnTrigger bool) *NewOrderService {
	s.CloseOnTrigger = closeOnTrigger
	return s
}

// set OrderLinkId
func (s *NewOrderService) orderLinkId(orderLinkId string) *NewOrderService {
	s.OrderLinkId = orderLinkId
	return s
}