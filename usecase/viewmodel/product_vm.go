package viewmodel

// ProductVM ...
type ProductVM struct {
	ID                 uint              ` json:"id"`
	Code               string            ` json:"code"`
	ProductName        string            ` json:"name"`
	ProductImage       string            ` json:"image"`
	ProductDescription string            ` json:"description"`
	TargetMarket       string            ` json:"targetMarket"`
	ProductCategory    ProductCategoryVM ` json:"category"`
	Price              int               ` json:"price"`
	Variant            string            ` json:"variant"`
	Notes              string            ` json:"notes"`
	Company            ProductCompanyVM  ` json:"company"`
	CreatedAt          string            ` json:"createdAt"`
	// CompanyCode        string            ` json:"companyCode"`
	// UpdatedAt          string            ` json:"updatedAt"`
}

// ProductCategoryVM ...
type ProductCategoryVM struct {
	Code         string ` json:"code"`
	CategoryName string ` json:"categoryName"`
}

// ProductCompanyVM ...
type ProductCompanyVM struct {
	Code        string ` json:"code"`
	CompanyName string ` json:"companyName"`
}
