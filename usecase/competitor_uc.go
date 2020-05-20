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
	for _, com := range data {
		dataProduct, err := model.ProductOp.GetByCompanyCode(uc.DB, com.Code)
		if err != nil {
			return nil, pagination, err
		}

		dataPromotion, err := model.PromotionOp.GetByCompanyCode(uc.DB, com.Code)
		if err != nil {
			return nil, pagination, err
		}

		productRes := []viewmodel.ProductVM{}
		for _, prod := range dataProduct {
			tempRes := viewmodel.ProductVM{
				ID:                 prod.ID,
				Code:               prod.Code,
				ProductName:        prod.ProductName,
				ProductImage:       prod.ProductImage,
				ProductDescription: prod.ProductDescription.String,
				TargetMarket:       prod.TargetMarket.String,
				ProductCategory: viewmodel.ProductCategoryVM{
					Code:         prod.Code,
					CategoryName: prod.ProductCategory,
				},
				Price:   prod.Price,
				Variant: prod.Variant.String,
				Notes:   prod.Notes.String,
				Company: viewmodel.ProductCompanyVM{
					Code:        com.Code,
					CompanyName: com.CompanyName,
				},
				CreatedAt: prod.CreatedAt.String,
			}
			productRes = append(productRes, tempRes)
		}

		promotionRes := []viewmodel.PromotionVM{}
		for _, prom := range dataPromotion {
			tempRes := viewmodel.PromotionVM{
				ID:         prom.ID,
				Code:       prom.Code,
				PromoTitle: prom.PromoTitle,
				PromoImage: prom.PromoImage,
				PromoType: viewmodel.PromoTypeVM{
					Label: prom.PromoType.String,
					Value: prom.PromoType.String,
				},
				DisplayLocation:   prom.DisplayLocation.String,
				PromoStart:        prom.PromoStart.String,
				PromoEnd:          prom.PromoEnd.String,
				IndefiniteEndDate: prom.IndefiniteEndDate.Int64,
				PromoDetail:       prom.PromoDetail.String,
				Company: viewmodel.PromoCompanyVM{
					Code:        com.Code,
					CompanyName: com.CompanyName,
				},
				CreatedAt: prom.CreatedAt.String,
			}
			promotionRes = append(promotionRes, tempRes)
		}

		strength := com.Strength
		arrStrength := strings.Split(strength, "|")

		weak := com.Weakness
		arrWeak := strings.Split(weak, "|")

		resMap = append(resMap, viewmodel.CompanyVM{
			ID:            com.ID,
			Code:          com.Code,
			CompanyName:   com.CompanyName,
			Logo:          com.Logo,
			Description:   com.Description,
			Website:       com.Website,
			Established:   com.Established,
			NoOfEmployees: com.NoOfEmployees,
			Strength:      arrStrength,
			Weakness:      arrWeak,
			Products:      productRes,
			Promotions:    promotionRes,
			CreatedAt:     com.CreatedAt.String,
			// UpdatedAt:     com.UpdatedAt.String,
		})
	}

	return resMap, pagination, err
}

// GetDetailCompany ...
func (uc UC) GetDetailCompany(code string) (viewmodel.CompanyVM, error) {

	com, err := model.CompanyOp.GetDetailCompany(uc.DB, code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	dataProduct, err := model.ProductOp.GetByCompanyCode(uc.DB, com.Code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	dataPromotion, err := model.PromotionOp.GetByCompanyCode(uc.DB, com.Code)
	if err != nil {
		return viewmodel.CompanyVM{}, err
	}

	productRes := []viewmodel.ProductVM{}
	for _, prod := range dataProduct {
		tempRes := viewmodel.ProductVM{
			ID:                 prod.ID,
			Code:               prod.Code,
			ProductName:        prod.ProductName,
			ProductImage:       prod.ProductImage,
			ProductDescription: prod.ProductDescription.String,
			TargetMarket:       prod.TargetMarket.String,
			ProductCategory: viewmodel.ProductCategoryVM{
				Code:         prod.Code,
				CategoryName: prod.ProductCategory,
			},
			Price:   prod.Price,
			Variant: prod.Variant.String,
			Notes:   prod.Notes.String,
			Company: viewmodel.ProductCompanyVM{
				Code:        com.Code,
				CompanyName: com.CompanyName,
			},
			CreatedAt: prod.CreatedAt.String,
		}
		productRes = append(productRes, tempRes)
	}

	promotionRes := []viewmodel.PromotionVM{}
	for _, prom := range dataPromotion {
		tempRes := viewmodel.PromotionVM{
			ID:         prom.ID,
			Code:       prom.Code,
			PromoTitle: prom.PromoTitle,
			PromoImage: prom.PromoImage,
			PromoType: viewmodel.PromoTypeVM{
				Label: prom.PromoType.String,
				Value: prom.PromoType.String,
			},
			DisplayLocation:   prom.DisplayLocation.String,
			PromoStart:        prom.PromoStart.String,
			PromoEnd:          prom.PromoEnd.String,
			IndefiniteEndDate: prom.IndefiniteEndDate.Int64,
			PromoDetail:       prom.PromoDetail.String,
			Company: viewmodel.PromoCompanyVM{
				Code:        com.Code,
				CompanyName: com.CompanyName,
			},
			CreatedAt: prom.CreatedAt.String,
		}
		promotionRes = append(promotionRes, tempRes)
	}

	strength := com.Strength
	arrStrength := strings.Split(strength, "|")

	weak := com.Weakness
	arrWeak := strings.Split(weak, "|")

	res := viewmodel.CompanyVM{
		ID:            com.ID,
		Code:          com.Code,
		CompanyName:   com.CompanyName,
		Logo:          com.Logo,
		Description:   com.Description,
		Website:       com.Website,
		Established:   com.Established,
		NoOfEmployees: com.NoOfEmployees,
		Strength:      arrStrength,
		Weakness:      arrWeak,
		Products:      productRes,
		Promotions:    promotionRes,
		CreatedAt:     com.CreatedAt.String,
		// UpdatedAt:     data.UpdatedAt.String,
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

	dataProduct, err := model.ProductOp.GetAllProduct(uc.DB)
	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.ProductVM, 0)
	for _, prod := range dataProduct {

		company, err := model.CompanyOp.GetDetailCompany(uc.DB, prod.CompanyCode)
		if err != nil {
			return nil, pagination, err
		}

		category, err := model.CategoryOp.GetByCode(uc.DB, prod.ProductCategory)
		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.ProductVM{
			ID:                 prod.ID,
			Code:               prod.Code,
			ProductName:        prod.ProductName,
			ProductImage:       prod.ProductImage,
			ProductDescription: prod.ProductDescription.String,
			TargetMarket:       prod.TargetMarket.String,
			ProductCategory: viewmodel.ProductCategoryVM{
				Code:         category.Code,
				CategoryName: category.CategoryName,
			},
			Price:   prod.Price,
			Variant: prod.Variant.String,
			Notes:   prod.Notes.String,
			Company: viewmodel.ProductCompanyVM{
				Code:        company.Code,
				CompanyName: company.CompanyName,
			},
			CreatedAt: prod.CreatedAt.String,
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

	category, err := model.CategoryOp.GetByCode(uc.DB, data.ProductCategory)
	if err != nil {
		return viewmodel.ProductVM{}, err
	}

	res := viewmodel.ProductVM{
		ID:                 data.ID,
		Code:               data.Code,
		ProductName:        data.ProductName,
		ProductImage:       data.ProductImage,
		ProductDescription: data.ProductDescription.String,
		TargetMarket:       data.TargetMarket.String,
		ProductCategory: viewmodel.ProductCategoryVM{
			Code:         category.Code,
			CategoryName: category.CategoryName,
		},
		Price:   data.Price,
		Variant: data.Variant.String,
		Notes:   data.Notes.String,
		Company: viewmodel.ProductCompanyVM{
			Code:        company.Code,
			CompanyName: company.CompanyName,
		},
		CreatedAt: data.CreatedAt.String,
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
	for _, prom := range data {
		if err != nil {
			return nil, pagination, err
		}

		company, err := model.CompanyOp.GetDetailCompany(uc.DB, prom.CompanyCode)
		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, viewmodel.PromotionVM{
			ID:         prom.ID,
			Code:       prom.Code,
			PromoTitle: prom.PromoTitle,
			PromoImage: prom.PromoImage,
			PromoType: viewmodel.PromoTypeVM{
				Label: prom.PromoType.String,
				Value: prom.PromoType.String,
			},
			DisplayLocation:   prom.DisplayLocation.String,
			PromoStart:        prom.PromoStart.String,
			PromoEnd:          prom.PromoEnd.String,
			IndefiniteEndDate: prom.IndefiniteEndDate.Int64,
			PromoDetail:       prom.PromoDetail.String,
			CreatedAt:         prom.CreatedAt.String,
			Company: viewmodel.PromoCompanyVM{
				Code:        company.Code,
				CompanyName: company.CompanyName,
			},
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

	company, err := model.CompanyOp.GetDetailCompany(uc.DB, data.CompanyCode)
	if err != nil {
		return viewmodel.PromotionVM{}, err
	}

	res := viewmodel.PromotionVM{
		ID:         data.ID,
		Code:       data.Code,
		PromoTitle: data.PromoTitle,
		PromoImage: data.PromoImage,
		PromoType: viewmodel.PromoTypeVM{
			Label: data.PromoType.String,
			Value: data.PromoType.String,
		},
		DisplayLocation:   data.DisplayLocation.String,
		PromoStart:        data.PromoStart.String,
		PromoEnd:          data.PromoEnd.String,
		IndefiniteEndDate: data.IndefiniteEndDate.Int64,
		PromoDetail:       data.PromoDetail.String,
		CreatedAt:         data.CreatedAt.String,
		Company: viewmodel.PromoCompanyVM{
			Code:        company.Code,
			CompanyName: company.CompanyName,
		},
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
func (uc UC) SearchExport(filters map[string]string) ([]map[string]interface{}, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	types := filters["types"]
	startFrom := filters["start_from"]
	endTo := filters["end_to"]

	if types == "product" {
		dataProduct, err := model.ProductOp.GetProductByDate(uc.DB, startFrom, endTo)
		if err != nil {
			return nil, pagination, err
		}

		resMap := make([]map[string]interface{}, 0)
		for _, prod := range dataProduct {

			if err != nil {
				return nil, pagination, err
			}

			companyProd, err := model.CompanyOp.GetDetailCompany(uc.DB, prod.CompanyCode)
			if err != nil {
				return nil, pagination, err
			}

			categoryProd, err := model.CategoryOp.GetByCode(uc.DB, prod.ProductCategory)
			if err != nil {
				return nil, pagination, err
			}

			resMap = append(resMap, map[string]interface{}{
				"id":   prod.ID,
				"type": "product",
				"code": prod.Code,
				"company": viewmodel.ProductCompanyVM{
					Code:        companyProd.Code,
					CompanyName: companyProd.CompanyName,
				},
				"name":         prod.ProductName,
				"image":        prod.ProductImage,
				"description":  prod.ProductDescription.String,
				"targetMarket": prod.TargetMarket.String,
				"category": viewmodel.ProductCategoryVM{
					Code:         categoryProd.Code,
					CategoryName: categoryProd.CategoryName,
				},
				"price":     prod.Price,
				"variant":   prod.Variant.String,
				"notes":     prod.Notes.String,
				"createdAt": prod.CreatedAt.String,
			})
		}

		return resMap, pagination, err
	} else {
		dataPromotion, err := model.PromotionOp.GetPromotionByDate(uc.DB, startFrom, endTo)
		if err != nil {
			return nil, pagination, err
		}

		resMap := make([]map[string]interface{}, 0)
		for _, prom := range dataPromotion {

			if err != nil {
				return nil, pagination, err
			}

			companyPromotion, err := model.CompanyOp.GetDetailCompany(uc.DB, prom.CompanyCode)
			if err != nil {
				return nil, pagination, err
			}

			resMap = append(resMap, map[string]interface{}{
				"id":    prom.ID,
				"type":  "promotion",
				"code":  prom.Code,
				"title": prom.PromoTitle,
				"image": prom.PromoImage,
				"promoType": viewmodel.PromoTypeVM{
					Label: prom.PromoType.String,
					Value: prom.PromoType.String,
				},
				"displayLocation":   prom.DisplayLocation.String,
				"start":             prom.PromoStart.String,
				"end":               prom.PromoEnd.String,
				"indefiniteEndDate": prom.IndefiniteEndDate.Int64,
				"detail":            prom.PromoDetail.String,
				"company": viewmodel.PromoCompanyVM{
					Code:        companyPromotion.Code,
					CompanyName: companyPromotion.CompanyName,
				},
				"createdAt": prom.CreatedAt.String,
			})
		}

		return resMap, pagination, err
	}

}

// GetAllDownload ...
func (uc UC) GetAllDownload() ([]viewmodel.DownloadVM, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.DownloadOp.GetAllDownload(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]viewmodel.DownloadVM, 0)
	for _, a := range data {

		if err != nil {
			return nil, pagination, err
		}

		resu := a.Result
		arrResu := strings.Split(resu, "|")

		viewResult := make([]map[string]interface{}, 0)

		if a.Type.String == "product" {

			for i := range arrResu {
				dataProduct, err := model.ProductOp.GetDetailProduct(uc.DB, arrResu[i])
				if err != nil {
					return nil, pagination, err
				}

				category, err := model.CategoryOp.GetByCode(uc.DB, dataProduct.ProductCategory)
				if err != nil {
					return nil, pagination, err
				}

				company, err := model.CompanyOp.GetDetailCompany(uc.DB, dataProduct.CompanyCode)
				if err != nil {
					return nil, pagination, err
				}

				viewResult = append(viewResult, map[string]interface{}{
					"id":           dataProduct.ID,
					"code":         dataProduct.Code,
					"name":         dataProduct.ProductName,
					"image":        dataProduct.ProductImage,
					"description":  dataProduct.ProductDescription.String,
					"targetMarket": dataProduct.TargetMarket.String,
					"category": viewmodel.ProductCategoryVM{
						Code:         category.Code,
						CategoryName: category.CategoryName,
					},
					"price":   dataProduct.Price,
					"variant": dataProduct.Variant.String,
					"notes":   dataProduct.Notes.String,
					"company": viewmodel.ProductCompanyVM{
						Code:        company.Code,
						CompanyName: company.CompanyName,
					},
					"createdAt": dataProduct.CreatedAt.String,
				})

			}
		} else {
			for i := range arrResu {
				dataPromotion, err := model.PromotionOp.GetDetailPromotion(uc.DB, arrResu[i])
				if err != nil {
					return nil, pagination, err
				}

				company, err := model.CompanyOp.GetDetailCompany(uc.DB, dataPromotion.CompanyCode)
				if err != nil {
					return nil, pagination, err
				}

				viewResult = append(viewResult, map[string]interface{}{
					"id":    dataPromotion.ID,
					"code":  dataPromotion.Code,
					"title": dataPromotion.PromoTitle,
					"image": dataPromotion.PromoImage,
					"type": viewmodel.PromoTypeVM{
						Label: dataPromotion.PromoType.String,
						Value: dataPromotion.PromoType.String,
					},
					"displayLocation":   dataPromotion.DisplayLocation.String,
					"start":             dataPromotion.PromoStart.String,
					"end":               dataPromotion.PromoEnd.String,
					"indefiniteEndDate": dataPromotion.IndefiniteEndDate.Int64,
					"detail":            dataPromotion.PromoDetail.String,
					"createdAt":         dataPromotion.CreatedAt.String,
					"company": viewmodel.PromoCompanyVM{
						Code:        company.Code,
						CompanyName: company.CompanyName,
					},
				})

			}

		}

		resMap = append(resMap, viewmodel.DownloadVM{
			ID:                         a.ID,
			Code:                       a.Code,
			DownloadOn:                 a.DownloadOn.String,
			Type:                       a.Type.String,
			NumberOfProductOrPromotion: a.NumberOfProductOrPromotion.String,
			Date: viewmodel.DateVM{
				Start: a.Start.String,
				End:   a.End.String,
			},
			URLRef:    a.URLRef.String,
			Result:    viewResult,
			CreatedAt: a.CreatedAt.String,
		})
	}

	return resMap, pagination, err
}

// AddDownload ...
func (uc UC) AddDownload(
	code string,
	downloadOn string,
	types string,
	numberOfProductOrPromotion int,
	start string,
	end string,
	urlRef string,
	result string,

) (int64, error) {

	dt, err := model.DownloadOp.StoreDownload(uc.DB, code, downloadOn, types, numberOfProductOrPromotion, start, end, urlRef, result, time.Now().UTC())
	return dt, err
}

// DeleteDownload ...
func (uc UC) DeleteDownload(code string) ([]*model.DownloadEntity, error) {

	dt, err := model.DownloadOp.DeleteDownload(uc.DB, code, time.Now().UTC())
	return dt, err
}
