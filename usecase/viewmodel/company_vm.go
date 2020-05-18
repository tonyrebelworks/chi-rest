package viewmodel

// CompanyVM ...
type CompanyVM struct {
	ID            uint          ` json:"id"`
	Code          string        ` json:"code"`
	CompanyName   string        ` json:"companyName"`
	Logo          string        ` json:"logo"`
	Description   string        ` json:"description"`
	Website       string        ` json:"website"`
	Established   string        ` json:"established"`
	NoOfEmployees int           ` json:"noOfEmployees"`
	Strength      []string      ` json:"strength"`
	Weakness      []string      ` json:"weakness"`
	Products      []ProductVM   ` json:"products"`
	Promotions    []PromotionVM ` json:"promotions"`
	CreatedAt     string        ` json:"createdAt"`
	UpdatedAt     string        ` json:"updatedAt"`
}
