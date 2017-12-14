package party

import "time"

type Party struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Query   Query  `json:"query"`
}

type Query struct {
	ID        int64     `json:"id"`
	QueryerID int64     `json:"queryer_id"`
	Query     string    `json:"query"`
	QueryDate time.Time `json:"query_date"`
}
