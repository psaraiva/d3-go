package database

import (
	"database/sql"

	"psaraiva/d3/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Update(order *entity.Order) error {
	stmt, err := r.Db.Prepare("UPDATE orders SET price = ?, tax = ?, final_price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.Price, order.Tax, order.FinalPrice, order.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() (*[]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &orders, nil
}
