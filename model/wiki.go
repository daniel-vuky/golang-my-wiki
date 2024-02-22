package model

type Wiki struct {
	WikiId     uint64 `json:"wiki_id"`
	Title      string `json:"title"`
	CategoryId uint64 `json:"category_id"`
	UserId     uint64 `json:"user_id"`
	Body       string `json:"body"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
