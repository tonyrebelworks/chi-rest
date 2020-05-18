package viewmodel

// PromotionVM ...
type PromotionVM struct {
	ID                uint        ` json:"id"`
	Code              string      ` json:"code"`
	CompanyCode       string      ` json:"companyCode"`
	PromoTitle        string      ` json:"title"`
	PromoImage        string      ` json:"image"`
	PromoType         PromoTypeVM ` json:"type"`
	DisplayLocation   string      ` json:"displayLocation"`
	PromoStart        string      ` json:"start"`
	PromoEnd          string      ` json:"end"`
	IndefiniteEndDate int64       ` json:"indefiniteEndDate"`
	PromoDetail       string      ` json:"detail"`
	CreatedAt         string      ` json:"createdAt"`
	UpdatedAt         string      ` json:"updatedAt"`
}

// PromoTypeVM ...
type PromoTypeVM struct {
	Label string ` json:"label"`
	Value string ` json:"value"`
}
