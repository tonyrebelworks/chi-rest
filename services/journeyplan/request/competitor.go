package request

// AddCompanyRequest ...
type AddCompanyRequest struct {
	CompanyName   string `json:"companyName" validate:"required"`
	Logo          string `json:"logo"`
	Description   string `json:"description" `
	Website       string `json:"website"`
	Established   string `json:"established"`
	NoOfEmployees int64  `json:"noOfEmployees"`
	Strength      string `json:"strength"`
	Weakness      string `json:"weakness"`
}

// UpdateCompanyRequest ...
type UpdateCompanyRequest struct {
	CompanyName   string `json:"companyName" validate:"required"`
	Logo          string `json:"logo"`
	Description   string `json:"description" `
	Website       string `json:"website"`
	Established   string `json:"established"`
	NoOfEmployees int64  `json:"noOfEmployees"`
	Strength      string `json:"strength"`
	Weakness      string `json:"weakness"`
}

// AddProductRequest ...
type AddProductRequest struct {
	CompanyCode        string ` json:"companyID"  validate:"required"`
	ProductName        string ` json:"name"  validate:"required"`
	ProductImage       string ` json:"image"`
	ProductDescription string ` json:"description"`
	TargetMarket       string ` json:"targetMarket"`
	ProductCategory    string ` json:"category" validate:"required"`
	Price              int    ` json:"price" validate:"required"`
	Variant            string ` json:"variant" validate:"required"`
	Notes              string ` json:"notes"`
}

// UpdateProductRequest ...
type UpdateProductRequest struct {
	CompanyCode        string ` json:"companyID"`
	ProductName        string ` json:"name"`
	ProductImage       string ` json:"image"`
	ProductDescription string ` json:"description"`
	TargetMarket       string ` json:"targetMarket"`
	ProductCategory    string ` json:"category"`
	Price              int    ` json:"price"`
	Variant            string ` json:"variant"`
	Notes              string ` json:"notes"`
}

// AddPromotionRequest ...
type AddPromotionRequest struct {
	CompanyCode       string ` json:"companyID"  validate:"required"`
	PromoTitle        string ` json:"title" validate:"required"`
	PromoImage        string ` json:"image"`
	PromoType         string ` json:"type"  validate:"required"`
	DisplayLocation   string ` json:"displayLocation"`
	PromoStart        string ` json:"start"  validate:"required"`
	PromoEnd          string ` json:"end"`
	IndefiniteEndDate int64  ` json:"indefiniteEndDate"`
	PromoDetail       string ` json:"detail"`
}

// // ObjectPromoTypeVM ...
// type ObjectPromoTypeVM struct {
// 	Label string ` json:"label"`
// 	Value string ` json:"value"`
// }

// UpdatePromotionRequest ...
type UpdatePromotionRequest struct {
	CompanyCode       string ` json:"companyID"`
	PromoTitle        string ` json:"title"`
	PromoImage        string ` json:"image"`
	PromoType         string ` json:"type"`
	DisplayLocation   string ` json:"displayLocation"`
	PromoStart        string ` json:"start"`
	PromoEnd          string ` json:"end"`
	IndefiniteEndDate int64  ` json:"indefiniteEndDate"`
	PromoDetail       string ` json:"detail"`
}

// AddCategoryRequest ...
type AddCategoryRequest struct {
	CategoryName string ` json:"categoryName"  validate:"required"`
}

// UpdateCategoryRequest ...
type UpdateCategoryRequest struct {
	CompanyCode string ` json:"companyCode"`
}

// AddDownloadRequest ...
type AddDownloadRequest struct {
	DownloadOn                 string ` json:"downloadOn"`
	Type                       string ` json:"type"`
	NumberOfProductOrPromotion int    ` json:"numberOfProductOrPromotion"`
	Start                      string ` json:"start"`
	End                        string ` json:"end"`
	URLRef                     string ` json:"urlRef"`
	Result                     string ` json:"result"`
}

// UpdateDownloadRequest ...
type UpdateDownloadRequest struct {
	DownloadOn                 string ` json:"downloadOn"`
	Type                       string ` json:"type"`
	NumberOfProductOrPromotion string ` json:"numberOfProductOrPromotion"`
	Start                      string ` json:"start"`
	End                        string ` json:"end"`
	URLRef                     string ` json:"urlRef"`
	Result                     string ` json:"result"`
}
