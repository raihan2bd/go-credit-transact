package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel is the type for database connection values
type DBModel struct {
	DB *sql.DB
}

// models is the wrapper for all the models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Widget is the type for all widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"_"`
	UpdatedAt      time.Time `json:"_"`
}

// Order is the type for all orders
type Order struct {
	ID            int       `json:"id"`
	WidgetID      int       `json:"widget_id"`
	TransactionID int       `json:"transaction_id"`
	StatusID      int       `json:"status_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"_"`
	UpdatedAt     time.Time `json:"_"`
}

// Status is the type for all statuses
type Status struct {
	ID        int       `json:"id"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

// TransactionStatus is the type for all Transaction Statuses
type TransactionStatus struct {
	ID        int       `json:"id"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

// Transaction is the type for all Transactions
type Transaction struct {
	ID                  int       `json:"id"`
	Amount              int       `json:"amount"`
	Currency            string    `json:"currency"`
	LastFour            string    `json:"last_four"`
	BankReturnCode      string    `json:"bank_return_code"`
	TransactionStatusID int       `json:"transaction_status_id"`
	CreatedAt           time.Time `json:"_"`
	UpdatedAt           time.Time `json:"_"`
}

// User is the type for all users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

// GetWidget will get widget by id.
func (m *DBModel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select id, name from widgets where id = ?", id)

	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, nil
	}

	return widget, nil

}
