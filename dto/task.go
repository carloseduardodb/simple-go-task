package dto

type Task struct {
	Id        *int    `json:"id,omitempty"`
	Text      string  `json:"text"`
	Completed bool    `json:"completed"`
	CreatedAt *string `json:"created_at,omitempty"`
}
