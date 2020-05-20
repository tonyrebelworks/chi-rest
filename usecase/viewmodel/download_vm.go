package viewmodel

// DownloadVM ...
type DownloadVM struct {
	ID                         uint                     ` json:"id"`
	Code                       string                   ` json:"code"`
	DownloadOn                 string                   ` json:"downloadOn"`
	Type                       string                   ` json:"type"`
	NumberOfProductOrPromotion string                   ` json:"numberOfProductOrPromotion"`
	Date                       DateVM                   ` json:"date"`
	URLRef                     string                   ` json:"urlRef"`
	Result                     []map[string]interface{} ` json:"result"`
	// Result    string ` json:"result"`
	CreatedAt string ` json:"createdAt"`
}

// DateVM ...
type DateVM struct {
	Start string ` json:"start"`
	End   string ` json:"end"`
}
