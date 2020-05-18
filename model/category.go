package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// CategoryEntity ...
type CategoryEntity struct {
	ID           uint           `db:"id" json:"id"`
	Code         string         `db:"code" json:"code"`
	CategoryName string         `db:"category_name" json:"categoryName"`
	CreatedAt    sql.NullString `db:"created_at" json:"createdAt"`
	DeletedAt    sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type categoryOp struct{}

// CategoryOp ...
var CategoryOp = &categoryOp{}

// GetAllCategory ...
func (op *categoryOp) GetAllCategory(db *sqlx.DB) ([]CategoryEntity, error) {
	var (
		err error
	)

	res := []CategoryEntity{}

	native := "SELECT id, code, category_name, created_at, deleted_at FROM category WHERE deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql)
	// fmt.Print(err)
	return res, err
}

// GetByCode ...
func (op *categoryOp) GetByCode(db *sqlx.DB, code string) (CategoryEntity, error) {
	var err error

	res := CategoryEntity{}
	err = db.Get(&res, "SELECT * FROM category WHERE code = ? LIMIT 1", code)

	return res, err
}

// StoreCategory ...
func (op *categoryOp) StoreCategory(
	db *sqlx.DB,
	code string,
	categoryName string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO category (code, category_name , created_at) VALUES ( ?,?,?)"
	res, err := db.Exec(sql, code, categoryName, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// DeleteCategory ...
func (op *categoryOp) DeleteCategory(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*CategoryEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*CategoryEntity{}
	sql := "UPDATE category SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}
