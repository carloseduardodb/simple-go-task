package models

type Task struct {
	Id        int    `db:"id"`
	Text      string `db:"text"`
	Completed bool   `db:"completed"`
	CreatedAt string `db:"created_at"`
}
