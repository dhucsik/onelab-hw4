package models

import "time"

type Transaction struct {
	ID            string    `json:"id"`
	SenderID      string    `json:"sender_id"`
	PaymentTime   time.Time `json:"payment_time"`
	BookID        string    `json:"book_id"`
	PaymentAmount float64   `json:"payment_amount"`
}
