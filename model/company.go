package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CompanyEntity ...
type CompanyEntity struct {
	ID            uint           `db:"id" json:"id"`
	Code          string         `db:"code" json:"code"`
	CompanyName   string         `db:"company_name" json:"companyName"`
	Logo          string         `db:"logo" json:"logo"`
	Description   string         `db:"description" json:"description"`
	Website       string         `db:"website" json:"website"`
	Established   string         `db:"established" json:"established"`
	NoOfEmployees int            `db:"no_of_employees" json:"noOfEmployees"`
	Strength      string         `db:"strength" json:"strength"`
	Weakness      string         `db:"weakness" json:"weakness"`
	CreatedAt     sql.NullString `db:"created_at" json:"createdAt"`
	UpdatedAt     sql.NullString `db:"updated_at" json:"updatedAt"`
	DeletedAt     sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type companyOp struct{}

// CompanyOp ...
var CompanyOp = &companyOp{}

// GetAllCompany ...
func (op *companyOp) GetAllCompany(db *sqlx.DB) ([]CompanyEntity, error) {
	var (
		err error
	)

	res := []CompanyEntity{}

	native := "SELECT id, code, company_name, logo, description, website, no_of_employees, strength, weakness, created_at, deleted_at FROM company "

	sql := native
	err = db.Select(&res, sql)

	return res, err
}
