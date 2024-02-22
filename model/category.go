package model

type Category struct {
	CategoryId       uint64 `json:"category_id"`
	Name             string `json:"name"`
	UserId           uint64 `json:"user_id"`
	ShortDescription string `json:"short_description"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
