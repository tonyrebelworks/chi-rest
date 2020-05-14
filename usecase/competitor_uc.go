package usecase

import (
	"chi-rest/model"
	"chi-rest/usecase/viewmodel"
)

// GetAllCompany ...
func (uc UC) GetAllCompany() ([]map[string]interface{}, viewmodel.SimplePaginationVM, error) {
	var (
		pagination viewmodel.SimplePaginationVM
	)

	data, err := model.CompanyOp.GetAllCompany(uc.DB)

	if err != nil {
		return nil, pagination, err
	}

	resMap := make([]map[string]interface{}, 0)
	for _, a := range data {

		if err != nil {
			return nil, pagination, err
		}

		resMap = append(resMap, map[string]interface{}{
			"id":            a.ID,
			"code":          a.Code,
			"companyName":   a.CompanyName,
			"logo":          a.Logo,
			"description":   a.Description,
			"website":       a.Website,
			"established":   a.Established,
			"noOfEmployees": a.NoOfEmployees,
			"strength":      a.Strength,
			"weakness":      a.Weakness,
			"createdAt":     a.CreatedAt.String,
			"updatedAt":     a.UpdatedAt.String,
			"deletedAt":     a.DeletedAt.String,
		})
	}

	return resMap, pagination, err
}
