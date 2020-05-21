package handler

import (
	"chi-rest/services/journeyplan/request"
	"chi-rest/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	validator "gopkg.in/go-playground/validator.v9"
)

//GetAllCompany
func (h *Contract) GetAllCompany(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllCompany()

	h.SendSuccess(w, res, pagination)
	return
}

// GetDetailCompany ...
func (h *Contract) GetDetailCompany(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailCompany(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddCompany ...
func (h *Contract) AddCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddCompanyRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := xid.New().String()

	lastID, err := usecase.UC{h.App}.AddCompany(
		code,
		req.CompanyName,
		req.Logo,
		req.Description,
		req.Website,
		req.Established,
		req.NoOfEmployees,
		req.Strength,
		req.Weakness,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// UpdateCompany ...
func (h *Contract) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateCompanyRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateCompany(
		code,
		req.CompanyName,
		req.Logo,
		req.Description,
		req.Website,
		req.Established,
		req.NoOfEmployees,
		req.Strength,
		req.Weakness,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// DeleteCompany ...
func (h *Contract) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteCompany(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

//GetAllProduct ...
func (h *Contract) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllProduct()

	h.SendSuccess(w, res, pagination)
	return
}

// GetDetailProduct ...
func (h *Contract) GetDetailProduct(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailProduct(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddProduct ...
func (h *Contract) AddProduct(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddProductRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := xid.New().String()

	lastID, err := usecase.UC{h.App}.AddProduct(
		code,
		req.CompanyCode,
		req.ProductName,
		req.ProductImage,
		req.ProductDescription,
		req.TargetMarket,
		req.ProductCategory,
		req.Price,
		req.Variant,
		req.Notes,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// UpdateProduct ...
func (h *Contract) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdateProductRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdateProduct(
		code,
		req.CompanyCode,
		req.ProductName,
		req.ProductImage,
		req.ProductDescription,
		req.TargetMarket,
		req.ProductCategory,
		req.Price,
		req.Variant,
		req.Notes,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// DeleteProduct ...
func (h *Contract) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteProduct(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetAllPromotion ...
func (h *Contract) GetAllPromotion(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllPromotion()

	h.SendSuccess(w, res, pagination)
	return
}

// GetDetailPromotion ...
func (h *Contract) GetDetailPromotion(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.GetDetailPromotion(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// AddPromotion ...
func (h *Contract) AddPromotion(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddPromotionRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := xid.New().String()

	lastID, err := usecase.UC{h.App}.AddPromotion(
		code,
		req.CompanyCode,
		req.PromoTitle,
		req.PromoImage,
		req.PromoType,
		req.DisplayLocation,
		req.PromoStart,
		req.PromoEnd,
		req.IndefiniteEndDate,
		req.PromoDetail,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// UpdatePromotion ...
func (h *Contract) UpdatePromotion(w http.ResponseWriter, r *http.Request) {
	var err error

	req := request.UpdatePromotionRequest{}
	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := chi.URLParam(r, "code")
	mdl := usecase.UC{h.App}

	_, err = mdl.UpdatePromotion(
		code,
		req.CompanyCode,
		req.PromoTitle,
		req.PromoImage,
		req.PromoType,
		req.DisplayLocation,
		req.PromoStart,
		req.PromoEnd,
		req.IndefiniteEndDate,
		req.PromoDetail,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, nil)
	return
}

// DeletePromotion ...
func (h *Contract) DeletePromotion(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeletePromotion(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

// GetAllCategory ...
func (h *Contract) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllCategory()

	h.SendSuccess(w, res, pagination)
	return
}

// AddCategory ...
func (h *Contract) AddCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddCategoryRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	// if len(req.CategoryName) < 1 {
	// 	h.SendBadRequest(w, "Category name tidak boleh kosong.")
	// 	return
	// }

	code := xid.New().String()

	lastID, err := usecase.UC{h.App}.AddCategory(
		code,
		req.CategoryName,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// DeleteCategory ...
func (h *Contract) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteCategory(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}

//SearchExport ...
func (h *Contract) SearchExport(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	types := r.URL.Query().Get("types")
	if types == "" {
		h.SendBadRequest(w, "Invalid parameter")
		return
	}
	startFrom := r.URL.Query().Get("start_from")
	if types == "" {
		h.SendBadRequest(w, "Invalid parameter")
		return
	}
	endTo := r.URL.Query().Get("end_to")
	if types == "" {
		h.SendBadRequest(w, "Invalid parameter")
		return
	}

	filters := map[string]string{
		"types":      types,
		"start_from": startFrom,
		"end_to":     endTo,
	}

	res, pagination, err := usecase.UC{h.App}.SearchExport(filters)

	h.SendSuccess(w, res, pagination)
	return
}

// GetAllDownload ...
func (h *Contract) GetAllDownload(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, pagination, err := usecase.UC{h.App}.GetAllDownload()

	h.SendSuccess(w, res, pagination)
	return
}

// AddDownload ...
func (h *Contract) AddDownload(w http.ResponseWriter, r *http.Request) {
	var err error
	req := request.AddDownloadRequest{}

	if err = h.Bind(r, &req); err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}
	if err = h.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	code := xid.New().String()

	lastID, err := usecase.UC{h.App}.AddDownload(
		code,
		req.DownloadOn,
		req.Type,
		req.NumberOfProductOrPromotion,
		req.Start,
		req.End,
		req.URLRef,
		req.Result,
	)
	if err != nil {
		h.SendBadRequest(w, err.Error())
		return
	}

	h.SendSuccess(w, map[string]interface{}{}, lastID)
	return
}

// DeleteDownload ...
func (h *Contract) DeleteDownload(w http.ResponseWriter, r *http.Request) {

	code := chi.URLParam(r, "code")
	res, err := usecase.UC{h.App}.DeleteDownload(code)
	if err != nil {
		fmt.Println(err)
		return
	}

	h.SendSuccess(w, res, nil)
	return
}
