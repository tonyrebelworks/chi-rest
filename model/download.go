package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// DownloadEntity ...
type DownloadEntity struct {
	ID                         uint           `db:"id" json:"id"`
	Code                       string         `db:"code" json:"code"`
	DownloadOn                 sql.NullString `db:"download_on" json:"DownloadOn"`
	Type                       sql.NullString `db:"type" json:"type"`
	NumberOfProductOrPromotion sql.NullString `db:"number_of_product_or_promotion" json:"numberOfProductOrPromotion"`
	Start                      sql.NullString `db:"start" json:"start"`
	End                        sql.NullString `db:"end" json:"end"`
	URLRef                     sql.NullString `db:"url_ref" json:"urlRef"`
	Result                     string         `db:"result" json:"result"`
	CreatedAt                  sql.NullString `db:"created_at" json:"createdAt"`
	DeletedAt                  sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type downloadOp struct{}

// DownloadOp ...
var DownloadOp = &downloadOp{}

// GetAllDownload ...
func (op *downloadOp) GetAllDownload(db *sqlx.DB) ([]DownloadEntity, error) {
	var (
		err error
	)

	res := []DownloadEntity{}

	native := "SELECT id, code, download_on, type, number_of_product_or_promotion, start, end, url_ref, result, created_at, deleted_at FROM download_activity WHERE deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql)
	return res, err
}

// GetByCode ...
func (op *downloadOp) GetByCode(db *sqlx.DB, code string) (DownloadEntity, error) {
	var err error

	res := DownloadEntity{}
	err = db.Get(&res, "SELECT * FROM download_activity WHERE code = ? LIMIT 1", code)

	return res, err
}

// StoreDownload ...
func (op *downloadOp) StoreDownload(
	db *sqlx.DB,
	code string,
	downloadOn string,
	types string,
	numberOfProductOrPromotion int,
	start string,
	end string,
	urlRef string,
	result string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO download_activity (code, download_on, type, number_of_product_or_promotion, start, end, url_ref, result , created_at) VALUES ( ?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, code, downloadOn, types, numberOfProductOrPromotion, start, end, urlRef, result, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// DeleteDownload ...
func (op *downloadOp) DeleteDownload(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*DownloadEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*DownloadEntity{}
	sql := "UPDATE download_activity SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}
