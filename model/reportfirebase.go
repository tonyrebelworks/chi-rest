package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// ReportFirebaseEntity ...
type ReportFirebaseEntity struct {
	ID         uint           `db:"id" json:"id"`
	ReportID   string         `db:"code" json:"reportID"`
	UserID     string         `db:"user_code" json:"userID"`
	JourneyID  string         `db:"journey_code" json:"journeyID"`
	URL        sql.NullString `db:"url" json:"url"`
	Start      sql.NullString `db:"start" json:"start"`
	End        sql.NullString `db:"end" json:"end"`
	Status     int            `db:"status" json:"status"`
	ReportDate sql.NullString `db:"report_date" json:"reportDate"`
	CreatedAt  sql.NullString `db:"created_at" json:"createdAt"`
	DeletedAt  sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type reportFirebaseOp struct{}

// ReportFirebaseOp ...
var ReportFirebaseOp = &reportFirebaseOp{}

// GetAllReportJourney ...
func (op *reportFirebaseOp) GetAllReportJourney(db *sqlx.DB, journeyID string) ([]ReportFirebaseEntity, error) {
	var (
		err error
	)

	res := []ReportFirebaseEntity{}

	native := "SELECT * FROM report_journey WHERE journey_code = ? AND deleted_at IS NULL"

	sql := native
	err = db.Select(&res, sql, journeyID)
	// fmt.Print(err)
	return res, err
}

// GetByJourneyCode ...
func (op *reportFirebaseOp) GetByJourneyCode(db *sqlx.DB, code string) ([]ReportFirebaseEntity, error) {
	var err error

	res := []ReportFirebaseEntity{}
	err = db.Select(&res, "SELECT * FROM report_journey WHERE journey_code = ?  AND deleted_at IS NULL ", code)
	fmt.Println(err)
	return res, err
}

// GetByParam ...
func (op *reportFirebaseOp) GetByParam(db *sqlx.DB, userID string, journeyID string, reportDate string) (ReportFirebaseEntity, error) {
	var err error

	res := ReportFirebaseEntity{}
	err = db.Get(&res, "SELECT *  FROM report_journey WHERE user_code = ? AND journey_code  = ? AND report_date = ?  AND deleted_at IS NULL", userID, journeyID, reportDate)
	// fmt.Println(err)
	return res, err
}

// GetByParamStarted ...
func (op *reportFirebaseOp) GetByParamStarted(db *sqlx.DB, userID string, reportDate string) (ReportFirebaseEntity, error) {
	var err error

	res := ReportFirebaseEntity{}
	err = db.Get(&res, "SELECT *  FROM report_journey WHERE user_code = ?  AND report_date = ? AND status = ? AND deleted_at IS NULL", userID, reportDate, "1")
	// fmt.Println(err)
	return res, err
}

// GetByReportID ...
func (op *reportFirebaseOp) GetByReportID(db *sqlx.DB, reportID string) (ReportFirebaseEntity, error) {
	var err error

	res := ReportFirebaseEntity{}
	err = db.Get(&res, "SELECT *  FROM report_journey WHERE code = ?  AND deleted_at IS NULL", reportID)
	return res, err
}

// Store ...
func (op *reportFirebaseOp) Store(
	db *sqlx.DB,
	code string,
	userID string, journeyID string, url string,
	start string, end string, status int, reportDate string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")
	end = reportDate + " 23:59:59"
	var sql = "INSERT INTO report_journey (code, user_code,journey_code, url, start, end, status, report_date, created_at) VALUES ( ?,?,?, ?,?,?,?, ?,?)"
	res, err := db.Exec(sql, code, userID, journeyID, url, start, end, "1", reportDate, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}

// GetReportID ...
func (op *reportFirebaseOp) GetReportID(db *sqlx.DB, userID string, journeyID string, reportDate string) (ReportFirebaseEntity, error) {
	var err error

	res := ReportFirebaseEntity{}
	err = db.Get(&res, "SELECT * FROM report_journey WHERE user_code = ? AND journey_code  = ? AND report_date = ?  AND deleted_at IS NULL", userID, journeyID, reportDate)

	return res, err
}

// Update ...
func (op *reportFirebaseOp) Update(
	db *sqlx.DB,
	code string,
	url string,
	end string,
) ([]*ReportFirebaseEntity, error) {

	// updatedAt := changedAt.Format("2006-01-02 15:04:05")
	r := []*ReportFirebaseEntity{}
	var sql string
	var err error

	if url != "" {
		sql = "UPDATE report_journey SET url = ? WHERE code = ?  "
		_, err = db.Exec(sql, url, code)
		return r, err
	}
	sql = "UPDATE report_journey SET end = ?, status = 0 WHERE code = ?  "
	_, err = db.Exec(sql, end, code)
	// fmt.Println(end)
	// fmt.Println(err)
	return r, err

}

// DeleteReportByID ...
func (op *reportFirebaseOp) DeleteReportByID(
	db *sqlx.DB,
	code string,
	changedAt time.Time,

) ([]*ReportFirebaseEntity, error) {

	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*ReportFirebaseEntity{}
	sql := "UPDATE report_journey SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}
