package model

import (
	"database/sql"
	"time"

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

	native := "SELECT id, code, company_name, logo, description, website, no_of_employees, strength, weakness, created_at, deleted_at FROM company WHERE deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql)

	return res, err
}

// GetDetailCompany ...
func (op *companyOp) GetDetailCompany(db *sqlx.DB, code string) (CompanyEntity, error) {
	var err error

	res := CompanyEntity{}
	err = db.Get(&res, "SELECT * FROM company WHERE code = ? LIMIT 1", code)

	return res, err
}

// StoreCompany ...
func (op *companyOp) StoreCompany(
	db *sqlx.DB,
	code string,
	companyName string,
	logo string,
	description string,
	website string,
	established string,
	noOfEmployees int64,
	strength string,
	weakness string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO company (code, company_name, logo, description, website, established,no_of_employees, strength, weakness, created_at) VALUES ( ?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, code, companyName, logo, description, website, established, noOfEmployees, strength, weakness, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// UpdateCompany ...
func (op *companyOp) UpdateCompany(
	db *sqlx.DB,
	code string,
	companyName string,
	logo string,
	description string,
	website string,
	established string,
	noOfEmployees int64,
	strength string,
	weakness string,
	changedAt time.Time,

) (int64, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "UPDATE company SET company_name = ?, logo = ?, description = ?, website = ?, established = ?,no_of_employees = ?, strength = ?, weakness = ?, updated_at = ? WHERE code = ?"

	_, err := db.Exec(sql, companyName, logo, description, website, established, noOfEmployees, strength, weakness, updatedAt, code)
	if err != nil {
		return 0, err
	}

	return 0, err
}

// DeleteCompany ...
func (op *companyOp) DeleteCompany(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*CompanyEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*CompanyEntity{}
	sql := "UPDATE company SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}
