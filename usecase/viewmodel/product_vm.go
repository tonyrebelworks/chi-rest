package viewmodel

// ProductVM ...
type ProductVM struct {
	ID                 uint              ` json:"id"`
	Code               string            ` json:"code"`
	CompanyCode        string            ` json:"companyCode"`
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
	UpdatedAt          string            ` json:"updatedAt"`
}

// ProductCategoryVM ...
type ProductCategoryVM struct {
	ID           uint   ` json:"id"`
	CategoryName string ` json:"categoryName"`
}

// ProductCompanyVM ...
type ProductCompanyVM struct {
	ID          string ` json:"id"`
	CompanyName string ` json:"companyName"`
}
