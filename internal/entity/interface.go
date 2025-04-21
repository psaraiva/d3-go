package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	Update(order *Order) error
	// GetTotal() (int, error)
}
