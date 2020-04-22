package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// ReportFirebaseEntity ...
type ReportFirebaseEntity struct {
	ID          uint           `db:"id" json:"id"`
	URL         string         `db:"url" json:"url"`
	JourneyCode string         `db:"journey_code" json:"journeyCode"`
	CreatedAt   sql.NullString `db:"created_at" json:"createdAt"`
}

type reportFirebaseOp struct{}

// ReportFirebaseOp ...
var ReportFirebaseOp = &reportFirebaseOp{}

// GetByJourneyCode ...
func (op *reportFirebaseOp) GetByJourneyCode(db *sqlx.DB, code string) ([]ReportFirebaseEntity, error) {
	var err error

	res := []ReportFirebaseEntity{}
	err = db.Select(&res, "SELECT id, url, journey_code, created_at FROM report_firebase WHERE journey_code = ? ", code)

	return res, err
}
