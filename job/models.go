package job

type Job struct {
	ID          int64  `json:"id"`
	JobID       int64  `json:"job_id"`
	TypeID      int64  `json:"type_id"`
	Title       string `json:"title"`
	Industry    string `json:"industry"`
	Location    string `json:"location"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Available   bool   `json:"available"`
}

type Type struct {
	ID       int64  `json:"id"`
	TypeID   int64  `json:"type_id"`
	Industry string `json:"industry"`
}
