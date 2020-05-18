package usecase

import (
	"chi-rest/model"
	"chi-rest/usecase/viewmodel"
	"strings"
	"time"
)

// GetAllCompany ...
func (uc UC) GetAllCompany() ([]viewmodel.CompanyVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.CompanyOp.GetAllCompany(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.CompanyVM, 0)
	for _, a := range data {
		dataProduct, err := model.ProductOp.GetByCompanyCode(uc.DB, a.Code)
		if err != nil {
			return nil, pagination, err
		}

		dataPromotion, err := model.PromotionOp.GetByCompanyCode(uc.DB, a.Code)
		if err != nil {
			return nil, pagination, err
		}

		company, err := model.CompanyOp.GetDetailCompany(uc.DB, a.Code)
		if err != nil {
			return nil, pagination, err
		}

		productRes := []viewmodel.ProductVM{}
		for _, a := range dataProduct {
			tempRes := viewmodel.ProductVM{
				ID:                 a.ID,
				Code:               a.Code,
				CompanyCode:        a.CompanyCode,
				ProductName:        a.ProductName,
				ProductImage:       a.ProductImage,
				ProductDescription: a.ProductDescription.String,
				TargetMarket:       a.TargetMarket.String,
				ProductCategory: viewmodel.ProductCategoryVM{
					ID:           a.ID,
					CategoryName: a.ProductCategory,
				},
				Price:   a.Price,
				Variant: a.Variant.String,
				Notes:   a.Notes.String,
				Company: viewmodel.ProductCompanyVM{
					ID:          a.CompanyCode,
					CompanyName: company.CompanyName,
				},
			}
			productRes = append(productRes, tempRes)
		}

		promotionRes := []viewmodel.PromotionVM{}
		for _, a := range dataPromotion {
			tempRes := viewmodel.PromotionVM{
				ID:          a.ID,
				Code:        a.Code,
				CompanyCode: a.CompanyCode,
				PromoTitle:  a.PromoTitle,
				PromoImage:  a.PromoImage,
				PromoType: viewmodel.PromoTypeVM{
					Label: "a",
					Value: "b",
				},
				DisplayLocation:   a.DisplayLocation.String,
				PromoStart:        a.PromoStart.String,
				PromoEnd:          a.PromoEnd.String,
				IndefiniteEndDate: a.IndefiniteEndDate.Int64,
				PromoDetail:       a.PromoDetail.String,
			}
			promotionRes = append(promotionRes, tempRes)
		}

		strength := a.Strength
		arrStrength := strings.Split(strength, "|")

		weak := a.Weakness
		arrWeak := strings.Split(weak, "|")

		resMap = append(resMap, viewmodel.CompanyVM{
			ID:            a.ID,
			Code:          a.Code,
			CompanyName:   a.CompanyName,
			Logo:          a.Logo,
			Description:   a.Description,
			Website:       a.Website,
			Established:   a.Established,
			NoOfEmployees: a.NoOfEmployees,
			Strength:      arrStrength,
			Weakness:      arrWeak,
			Products:      productRes,
			Promotions:    promotionRes,
			CreatedAt:     a.CreatedAt.String,
			UpdatedAt:     a.UpdatedAt.String,
		})
	}

	return resMap, pagination, err
}

// GetDetailCompany ...
func (uc UC) GetDetailCompany(code string) (viewmodel.CompanyVM, error) {

	data, err := model.CompanyOp.GetDetailCompany(uc.DB, code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	dataProduct, err := model.ProductOp.GetByCompanyCode(uc.DB, data.Code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	dataPromotion, err := model.PromotionOp.GetByCompanyCode(uc.DB, data.Code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	company, err := model.CompanyOp.GetDetailCompany(uc.DB, data.Code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	productRes := []viewmodel.ProductVM{}
	for _, a := range dataProduct {
		tempRes := viewmodel.ProductVM{
			ID:                 a.ID,
			Code:               a.Code,
			CompanyCode:        a.CompanyCode,
			ProductName:        a.ProductName,
			ProductImage:       a.ProductImage,
			ProductDescription: a.ProductDescription.String,
			TargetMarket:       a.TargetMarket.String,
			ProductCategory: viewmodel.ProductCategoryVM{
				ID:           a.ID,
				CategoryName: a.ProductCategory,
			},
			Price:   a.Price,
			Variant: a.Variant.String,
			Notes:   a.Notes.String,
			Company: viewmodel.ProductCompanyVM{
				ID:          a.CompanyCode,
				CompanyName: company.CompanyName,
			},
			CreatedAt: a.CreatedAt.String,
		}
		productRes = append(productRes, tempRes)
	}

	promotionRes := []viewmodel.PromotionVM{}
	for _, a := range dataPromotion {
		tempRes := viewmodel.PromotionVM{
			ID:          a.ID,
			Code:        a.Code,
			CompanyCode: a.CompanyCode,
			PromoTitle:  a.PromoTitle,
			PromoImage:  a.PromoImage,
			PromoType: viewmodel.PromoTypeVM{
				Label: "a",
				Value: "b",
			},
			DisplayLocation:   a.DisplayLocation.String,
			PromoStart:        a.PromoStart.String,
			PromoEnd:          a.PromoEnd.String,
			IndefiniteEndDate: a.IndefiniteEndDate.Int64,
			PromoDetail:       a.PromoDetail.String,
		}
		promotionRes = append(promotionRes, tempRes)
	}

	strength := data.Strength
	arrStrength := strings.Split(strength, "|")

	weak := data.Weakness
	arrWeak := strings.Split(weak, "|")

	res := viewmodel.CompanyVM{
		ID:            data.ID,
		Code:          data.Code,
		CompanyName:   data.CompanyName,
		Logo:          data.Logo,
		Description:   data.Description,
		Website:       data.Website,
		Established:   data.Established,
		NoOfEmployees: data.NoOfEmployees,
		Strength:      arrStrength,
		Weakness:      arrWeak,
		Products:      productRes,
		Promotions:    promotionRes,
		CreatedAt:     data.CreatedAt.String,
		UpdatedAt:     data.UpdatedAt.String,
	}

	return res, err
}

// AddCompany ...
func (uc UC) AddCompany(
	code string,
	companyName string,
	logo string,
	description string,
	website string,
	established string,
	noOfEmployees int64,
	strength string,
	weakness string,

) (int64, error) {

	dt, err := model.CompanyOp.StoreCompany(uc.DB, code, companyName, logo, description, website, established, noOfEmployees, strength, weakness, time.Now().UTC())
	return dt, err
}

// UpdateCompany ...
func (uc UC) UpdateCompany(
	code string,
	companyName string,
	logo string,
	description string,
	website string,
	established string,
	noOfEmployees int64,
	strength string,
	weakness string,

) (int64, error) {
	dt, err := model.CompanyOp.UpdateCompany(uc.DB, code, companyName, logo, description, website, established, noOfEmployees, strength, weakness, time.Now().UTC())
	return dt, err
}

// DeleteCompany ...
func (uc UC) DeleteCompany(code string) ([]*model.CompanyEntity, error) {

	dt, err := model.CompanyOp.DeleteCompany(uc.DB, code, time.Now().UTC())
	return dt, err
}

// GetAllProduct ...
func (uc UC) GetAllProduct() ([]viewmodel.ProductVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.ProductOp.GetAllProduct(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.ProductVM, 0)
	for _, a := range data {

		company, err := model.CompanyOp.GetDetailCompany(uc.DB, a.CompanyCode)
		if err != nil {
			return nil, pagination, err
		}

		category, err := model.CategoryOp.GetByCode(uc.DB, a.ProductCategory)
		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.ProductVM{
			ID:                 a.ID,
			Code:               a.Code,
			CompanyCode:        a.CompanyCode,
			ProductName:        a.ProductName,
			ProductImage:       a.ProductImage,
			ProductDescription: a.ProductDescription.String,
			TargetMarket:       a.TargetMarket.String,
			ProductCategory: viewmodel.ProductCategoryVM{
				ID:           category.ID,
				CategoryName: category.CategoryName,
			},
			Price:   a.Price,
			Variant: a.Variant.String,
			Notes:   a.Notes.String,
			Company: viewmodel.ProductCompanyVM{
				ID:          a.CompanyCode,
				CompanyName: company.CompanyName,
			},
			CreatedAt: a.CreatedAt.String,
			UpdatedAt: a.UpdatedAt.String,
		})
	}

	return resMap, pagination, err
}

// GetDetailProduct ...
func (uc UC) GetDetailProduct(code string) (viewmodel.ProductVM, error) {

	data, err := model.ProductOp.GetDetailProduct(uc.DB, code)
	if err != nil {
		return viewmodel.ProductVM{}, err
	}

	company, err := model.CompanyOp.GetDetailCompany(uc.DB, data.CompanyCode)
	if err != nil {
		return viewmodel.ProductVM{}, err
	}
	res := viewmodel.ProductVM{
		ID:                 data.ID,
		Code:               data.Code,
		CompanyCode:        data.CompanyCode,
		ProductName:        data.ProductName,
		ProductImage:       data.ProductImage,
		ProductDescription: data.ProductDescription.String,
		TargetMarket:       data.TargetMarket.String,
		ProductCategory: viewmodel.ProductCategoryVM{
			ID:           data.ID,
			CategoryName: data.ProductCategory,
		},
		Price:   data.Price,
		Variant: data.Variant.String,
		Notes:   data.Notes.String,
		Company: viewmodel.ProductCompanyVM{
			ID:          data.CompanyCode,
			CompanyName: company.CompanyName,
		},
		CreatedAt: data.CreatedAt.String,
		UpdatedAt: data.UpdatedAt.String,
	}

	return res, err
}

// AddProduct ...
func (uc UC) AddProduct(
	code string,
	companyCode string,
	productName string,
	productImage string,
	productDescription string,
	targetMarket string,
	productCategory string,
	price int,
	variant string,
	notes string,

) (int64, error) {

	dt, err := model.ProductOp.StoreProduct(uc.DB, code, companyCode, productName, productImage, productDescription, targetMarket, productCategory, price, variant, notes, time.Now().UTC())
	return dt, err
}

// UpdateProduct ...
func (uc UC) UpdateProduct(
	code string,
	companyCode string,
	productName string,
	productImage string,
	productDescription string,
	targetMarket string,
	productCategory string,
	price int,
	variant string,
	notes string,

) (int64, error) {
	dt, err := model.ProductOp.UpdateProduct(uc.DB, code, companyCode, productName, productImage, productDescription, targetMarket, productCategory, price, variant, notes, time.Now().UTC())
	return dt, err
}

// DeleteProduct ...
func (uc UC) DeleteProduct(code string) ([]*model.ProductEntity, error) {

	dt, err := model.ProductOp.DeleteProduct(uc.DB, code, time.Now().UTC())
	return dt, err
}

// GetAllPromotion ...
func (uc UC) GetAllPromotion() ([]viewmodel.PromotionVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.PromotionOp.GetAllPromotion(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.PromotionVM, 0)
	for _, a := range data {

		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.PromotionVM{
			ID:          a.ID,
			Code:        a.Code,
			CompanyCode: a.CompanyCode,
			PromoTitle:  a.PromoTitle,
			PromoImage:  a.PromoImage,
			PromoType: viewmodel.PromoTypeVM{
				Label: "a",
				Value: "b",
			},
			DisplayLocation:   a.DisplayLocation.String,
			PromoStart:        a.PromoStart.String,
			PromoEnd:          a.PromoEnd.String,
			IndefiniteEndDate: a.IndefiniteEndDate.Int64,
			PromoDetail:       a.PromoDetail.String,
			CreatedAt:         a.CreatedAt.String,
			UpdatedAt:         a.UpdatedAt.String,
		})
	}

	return resMap, pagination, err
}

// GetDetailPromotion ...
func (uc UC) GetDetailPromotion(code string) (viewmodel.PromotionVM, error) {

	data, err := model.PromotionOp.GetDetailPromotion(uc.DB, code)
	if err != nil {
		return viewmodel.PromotionVM{}, err
	}

	res := viewmodel.PromotionVM{
		ID:          data.ID,
		Code:        data.Code,
		CompanyCode: data.CompanyCode,
		PromoTitle:  data.PromoTitle,
		PromoImage:  data.PromoImage,
		PromoType: viewmodel.PromoTypeVM{
			Label: "a",
			Value: "b",
		},
		DisplayLocation:   data.DisplayLocation.String,
		PromoStart:        data.PromoStart.String,
		PromoEnd:          data.PromoEnd.String,
		IndefiniteEndDate: data.IndefiniteEndDate.Int64,
		PromoDetail:       data.PromoDetail.String,
		CreatedAt:         data.CreatedAt.String,
		UpdatedAt:         data.UpdatedAt.String,
	}

	return res, err
}

// AddPromotion ...
func (uc UC) AddPromotion(
	code string,
	companyCode string,
	promoTitle string,
	promoImage string,
	promoType string,
	displayLocation string,
	promoStart string,
	promoEnd string,
	indefiniteEndDate int64,
	promoDetail string,

) (int64, error) {

	dt, err := model.PromotionOp.StorePromotion(uc.DB, code, companyCode, promoTitle, promoImage, promoType, displayLocation, promoStart, promoEnd, indefiniteEndDate, promoDetail, time.Now().UTC())
	return dt, err
}

// UpdatePromotion ...
func (uc UC) UpdatePromotion(
	code string,
	companyCode string,
	promoTitle string,
	promoImage string,
	promoType string,
	displayLocation string,
	promoStart string,
	promoEnd string,
	indefiniteEndDate int64,
	promoDetail string,

) (int64, error) {
	dt, err := model.PromotionOp.UpdatePromotion(uc.DB, code, companyCode, promoTitle, promoImage, promoType, displayLocation, promoStart, promoEnd, indefiniteEndDate, promoDetail, time.Now().UTC())
	return dt, err
}

// DeletePromotion ...
func (uc UC) DeletePromotion(code string) ([]*model.PromotionEntity, error) {

	dt, err := model.PromotionOp.DeletePromotion(uc.DB, code, time.Now().UTC())
	return dt, err
}

// GetAllCategory ...
func (uc UC) GetAllCategory() ([]viewmodel.CategoryVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.CategoryOp.GetAllCategory(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.CategoryVM, 0)
	for _, a := range data {

		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.CategoryVM{
			ID:           a.ID,
			Code:         a.Code,
			CategoryName: a.CategoryName,
			CreatedAt:    a.CreatedAt.String,
		})
	}

	return resMap, pagination, err
}

// AddCategory ...
func (uc UC) AddCategory(
	code string,
	categoryName string,

) (int64, error) {

	dt, err := model.CategoryOp.StoreCategory(uc.DB, code, categoryName, time.Now().UTC())
	return dt, err
}

// DeleteCategory ...
func (uc UC) DeleteCategory(code string) ([]*model.CategoryEntity, error) {

	dt, err := model.CategoryOp.DeleteCategory(uc.DB, code, time.Now().UTC())
	return dt, err
}

// SearchExport ...
func (uc UC) SearchExport() ([]viewmodel.ProductVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.ProductOp.GetAllProduct(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.ProductVM, 0)
	for _, a := range data {

		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.ProductVM{
			ID:                 a.ID,
			Code:               a.Code,
			CompanyCode:        a.CompanyCode,
			ProductName:        a.ProductName,
			ProductImage:       a.ProductImage,
			ProductDescription: a.ProductDescription.String,
			TargetMarket:       a.TargetMarket.String,
			ProductCategory: viewmodel.ProductCategoryVM{
				ID:           a.ID,
				CategoryName: a.ProductCategory,
			},
			Price:     a.Price,
			Variant:   a.Variant.String,
			Notes:     a.Notes.String,
			CreatedAt: a.CreatedAt.String,
			UpdatedAt: a.UpdatedAt.String,
		})
	}

	return resMap, pagination, err
}
