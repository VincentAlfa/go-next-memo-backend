package model

type Memo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Memo string `json:"memo"`
	User_id int `json:"user_id"`
}