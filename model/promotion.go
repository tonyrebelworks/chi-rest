package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// PromotionEntity ...
type PromotionEntity struct {
	ID                uint           `db:"id" json:"id"`
	Code              string         `db:"code" json:"code"`
	CompanyCode       string         `db:"company_code" json:"companyCode"`
	PromoTitle        string         `db:"promo_title" json:"promoTitle"`
	PromoImage        string         `db:"promo_image" json:"promoImage"`
	PromoType         sql.NullString `db:"promo_type" json:"promoType"`
	DisplayLocation   sql.NullString `db:"display_location" json:"displayLocation"`
	PromoStart        sql.NullString `db:"promo_start" json:"promoStart"`
	PromoEnd          sql.NullString `db:"promo_end" json:"promoEnd"`
	IndefiniteEndDate sql.NullInt64  `db:"indefinite_end_date" json:"indefiniteEndDate"`
	PromoDetail       sql.NullString `db:"promo_detail" json:"promoDetail"`
	CreatedAt         sql.NullString `db:"created_at" json:"createdAt"`
	UpdatedAt         sql.NullString `db:"updated_at" json:"updatedAt"`
	DeletedAt         sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type promotionOp struct{}

// PromotionOp ...
var PromotionOp = &promotionOp{}

// GetAllPromotion ...
func (op *promotionOp) GetAllPromotion(db *sqlx.DB) ([]PromotionEntity, error) {
	var (
		err error
	)

	res := []PromotionEntity{}

	native := "SELECT id, code, company_code, promo_title, promo_image, promo_type, display_location, promo_start, promo_end, indefinite_end_date, promo_detail , created_at, deleted_at FROM company_promotion WHERE deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql)
	// fmt.Print(err)
	return res, err
}

// GetDetailPromotion ...
func (op *promotionOp) GetDetailPromotion(db *sqlx.DB, code string) (PromotionEntity, error) {
	var err error

	res := PromotionEntity{}
	err = db.Get(&res, "SELECT id, code, company_code, promo_title, promo_image, promo_type, display_location, promo_start, promo_end, indefinite_end_date, promo_detail  FROM company_promotion WHERE code = ? LIMIT 1", code)

	return res, err
}

// GetByCompanyCode ...
func (op *promotionOp) GetByCompanyCode(db *sqlx.DB, code string) ([]PromotionEntity, error) {
	var err error

	res := []PromotionEntity{}
	err = db.Select(&res, "SELECT * FROM company_promotion WHERE company_code = ? AND deleted_at IS NULL", code)

	return res, err
}

// StorePromotion ...
func (op *promotionOp) StorePromotion(
	db *sqlx.DB,
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
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO company_promotion (code, company_code, promo_title, promo_image, promo_type, display_location, promo_start, promo_end, indefinite_end_date, promo_detail , created_at) VALUES ( ?,?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, code, companyCode, promoTitle, promoImage, promoType, displayLocation, promoStart, promoEnd, indefiniteEndDate, promoDetail, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// UpdatePromotion ...
func (op *promotionOp) UpdatePromotion(
	db *sqlx.DB,
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
	changedAt time.Time,

) (int64, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "UPDATE company_promotion SET  company_code = ?, promo_title = ?, promo_image = ?, promo_type = ?, display_location = ?, promo_start = ?, promo_end = ?, indefinite_end_date = ?, promo_detail = ?, updated_at = ? WHERE code = ?"

	_, err := db.Exec(sql, companyCode, promoTitle, promoImage, promoType, displayLocation, promoStart, promoEnd, indefiniteEndDate, promoDetail, updatedAt, code)
	if err != nil {
		return 0, err
	}

	return 0, err
}

// DeletePromotion ...
func (op *promotionOp) DeletePromotion(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*PromotionEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*PromotionEntity{}
	sql := "UPDATE company_promotion SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}

// GetPromotionByDate ...
func (op *promotionOp) GetPromotionByDate(db *sqlx.DB, startFrom string, endTo string) ([]PromotionEntity, error) {
	var err error

	res := []PromotionEntity{}
	err = db.Select(&res, "SELECT * FROM company_promotion WHERE deleted_at IS NULL AND created_at BETWEEN ? AND ?", startFrom, endTo)

	return res, err
}
