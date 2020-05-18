package viewmodel

// CategoryVM ...
type CategoryVM struct {
	ID           uint   ` json:"id"`
	Code         string ` json:"code"`
	CategoryName string ` json:"categoryName"`
	CreatedAt    string ` json:"created_at"`
}
