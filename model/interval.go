package model

import (
	"github.com/jmoiron/sqlx"
)

// IntervalEntity ...
type IntervalEntity struct {
	ID            uint `db:"id" json:"id"`
	TimePerSecond int  `db:"time_per_second" json:"timePerSecond"`
}

type intervalOp struct{}

// IntervalOp ...
var IntervalOp = &intervalOp{}

// GetInterval ...
func (op *intervalOp) GetInterval(db *sqlx.DB) ([]IntervalEntity, error) {
	var (
		err error
	)

	res := []IntervalEntity{}

	native := "SELECT id, time_per_second FROM interval_time "
	err = db.Select(&res, native)

	return res, err
}

// UpdateInterval ...
func (op *intervalOp) UpdateInterval(
	db *sqlx.DB,
	timePerSecond int,

) (int64, error) {

	var sql = "UPDATE interval_time SET time_per_second = ? WHERE id = 1"

	_, err := db.Exec(sql, timePerSecond)
	if err != nil {
		return 0, err
	}

	return 0, err
}
