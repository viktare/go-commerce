package models

type Product struct {
	ID    string `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Code  string `json:"code" db:"code"`
}
