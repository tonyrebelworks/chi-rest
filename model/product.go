package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// ProductEntity ...
type ProductEntity struct {
	ID                 uint           `db:"id" json:"id"`
	Code               string         `db:"code" json:"code"`
	CompanyCode        string         ` db:"company_code" json:"companyCode"`
	ProductName        string         ` db:"product_name" json:"productName"`
	ProductImage       string         ` db:"product_image" json:"productImage"`
	ProductDescription sql.NullString ` db:"product_description" json:"productDescription"`
	TargetMarket       sql.NullString ` db:"target_market" json:"targetMarket"`
	ProductCategory    string         ` db:"product_category" json:"productCategory"`
	Price              int            ` db:"price" json:"price"`
	Variant            sql.NullString ` db:"variant" json:"variant"`
	Notes              sql.NullString ` db:"notes" json:"notes"`
	CreatedAt          sql.NullString `db:"created_at" json:"createdAt"`
	UpdatedAt          sql.NullString `db:"updated_at" json:"updatedAt"`
	DeletedAt          sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type productOp struct{}

// ProductOp ...
var ProductOp = &productOp{}

// GetAllProduct ...
func (op *productOp) GetAllProduct(db *sqlx.DB) ([]ProductEntity, error) {
	var (
		err error
	)

	res := []ProductEntity{}

	native := "SELECT id, code, company_code, product_name, product_image, product_description, target_market, product_category, price, variant, notes, created_at, deleted_at FROM company_product WHERE deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql)
	// fmt.Print(err)
	return res, err
}

// GetDetailProduct ...
func (op *productOp) GetDetailProduct(db *sqlx.DB, code string) (ProductEntity, error) {
	var err error

	res := ProductEntity{}
	err = db.Get(&res, "SELECT id, code, company_code, product_name, product_image, product_description, target_market, product_category, price, variant, notes, created_at, deleted_at FROM company_product WHERE code = ? LIMIT 1", code)

	return res, err
}

// GetByCompanyCode ...
func (op *productOp) GetByCompanyCode(db *sqlx.DB, code string) ([]ProductEntity, error) {
	var err error

	res := []ProductEntity{}
	err = db.Select(&res, "SELECT id, code, company_code, product_name, product_image, product_description, target_market, product_category, price, variant, notes, created_at, deleted_at FROM company_product WHERE company_code = ? AND deleted_at IS NULL", code)

	return res, err
}

// StoreProduct ...
func (op *productOp) StoreProduct(
	db *sqlx.DB,
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
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO company_product (code, company_code, product_name, product_image, product_description, target_market, product_category, price, variant, notes, created_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, code, companyCode, productName, productImage, productDescription, targetMarket, productCategory, price, variant, notes, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// UpdateProduct ...
func (op *productOp) UpdateProduct(
	db *sqlx.DB,
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
	changedAt time.Time,

) (int64, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "UPDATE company_product SET  company_code = ?, product_name = ?, product_image = ?, product_description = ?, target_market = ?, product_category = ?, price = ?, variant = ?, notes = ?, updated_at = ? WHERE code = ?"

	_, err := db.Exec(sql, companyCode, productName, productImage, productDescription, targetMarket, productCategory, price, variant, notes, updatedAt, code)
	if err != nil {
		return 0, err
	}

	return 0, err
}

// DeleteProduct ...
func (op *productOp) DeleteProduct(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*ProductEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*ProductEntity{}
	sql := "UPDATE company_product SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}

// GetProductByDate ...
func (op *productOp) GetProductByDate(db *sqlx.DB, startFrom string, endTo string) ([]ProductEntity, error) {
	var err error

	res := []ProductEntity{}
	err = db.Select(&res, "SELECT id, code, company_code, product_name, product_image, product_description, target_market, product_category, price, variant, notes, created_at, deleted_at FROM company_product WHERE deleted_at IS NULL AND created_at BETWEEN ? AND ?", startFrom, endTo)

	return res, err
}
