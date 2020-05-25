package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// TrackingTimeEntity ...
type TrackingTimeEntity struct {
	ID          uint           `db:"id" json:"id"`
	JourneyCode sql.NullString `db:"journey_code" json:"journeyCode"`
	ReportID    string         `db:"report_id" json:"reportID"`
	Latitude    string         `db:"latitude" json:"latitude"`
	Longitude   string         `db:"longitude" json:"longitude"`
	CreatedAt   sql.NullString `db:"created_at" json:"createdAt"`
	DeletedAt   sql.NullString `db:"deleted_at" json:"deletedAt"`
}

type trackingTimeOp struct{}

// TrackingTimeOp ...
var TrackingTimeOp = &trackingTimeOp{}

// GetByJourneyCode ...
func (op *trackingTimeOp) GetByJourneyCode(db *sqlx.DB, userCode string) ([]TrackingTimeEntity, error) {
	var err error

	res := []TrackingTimeEntity{}
	err = db.Select(&res, "SELECT id, journey_code, report_id, latitude, longitude, created_at, deleted_at FROM time_tracking WHERE  report_id = ? ", userCode)

	// fmt.Println(err)
	return res, err
}

// Store ...
func (op *trackingTimeOp) Store(
	db *sqlx.DB,
	reportJourneyID string,
	userCode string,
	latitude string,
	longitude string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")
	// userCode := "5qFKQb4kNJVFGsDBTp1NVrKojn12"
	var sql = "INSERT INTO time_tracking (report_id, latitude, longitude, created_at) VALUES ( ?, ?, ?, ?)"
	res, err := db.Exec(sql, reportJourneyID, latitude, longitude, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lID, nil
}
