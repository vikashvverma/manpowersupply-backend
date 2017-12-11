package party

type Party  struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}

type Query struct {
	ID string `json:"id"`
	QueryerID string `json:"queryer_id"`
	Query string `json:"query"`
	QueryDate string `json:"query_date"`
}