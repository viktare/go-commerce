package models

type Product struct {
	ID    string `json:"id" db:"id"`
	TITLE string `json:"title" db:"title"`
	CODE  string `json:"code" db:"code"`
}
