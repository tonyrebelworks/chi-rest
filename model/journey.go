package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// JourneyEntity ...
type JourneyEntity struct {
	ID              uint             `db:"id" json:"id"`
	Code            string           `db:"code" json:"code"`
	JourneyName     string           `db:"journey_name" json:"journeyName"`
	JourneySchedule string           `db:"journey_schedule" json:"journeySchedule"`
	Salesman        string           `db:"salesman" json:"assignedAuditor"`
	Sites           string           `db:"sites" json:"sites"`
	Questionnaires  string           `db:"questionnaires" json:"questionnaires"`
	Activity        string           `db:"activity" json:"activity"`
	Signatures      string           `db:"signatures" json:"signatures"`
	RequireSelfie   string           `db:"require_selfie" json:"requireSelfie"`
	EmailTo         string           `db:"email_to" json:"emailTargets"`
	StartJourney    sql.NullString   `db:"start_journey" json:"startJourney"`
	FinishJourney   sql.NullString   `db:"finish_journey" json:"finishJourney"`
	CreatedAt       sql.NullString   `db:"created_at" json:"createdAt"`
	UpdatedAt       sql.NullString   `db:"updated_at" json:"updatedAt"`
	DeletedAt       sql.NullString   `db:"deleted_at" json:"deletedAt"`
	AssignedAuditor []SalesmanEntity `json:"assignedAuditor2"`
}

// SalesmanEntity ...
type SalesmanEntity struct {
	UserID string `json:"userID"`
}

type journeyOp struct{}

// JourneyOp ...
var JourneyOp = &journeyOp{}

// GetAll ...
func (op *journeyOp) GetAll(db *sqlx.DB) ([]JourneyEntity, error) {

	activeQ := "WHERE deleted_at IS NULL"

	// sql := "SELECT id, code, journey_name, journey_schedule, salesman, sites, questionnaires, signatures, require_selfie, email_to, activity, start_journey, finish_journey, created_at, updated_at FROM journey_plan   WHERE deleted_at IS NULL LIMIT 10"

	res := []JourneyEntity{}
	err := db.Select(&res, "SELECT * FROM journey_plan "+activeQ)

	// fmt.Println(res)
	return res, err
}

// GetDetail ...
func (op *journeyOp) GetDetail(db *sqlx.DB, code string) ([]*JourneyEntity, error) {
	r := []*JourneyEntity{}
	sql := "SELECT id, code, journey_name, journey_schedule, salesman, sites, questionnaires, signatures, require_selfie, email_to, activity, start_journey, finish_journey, created_at, updated_at FROM journey_plan WHERE code = ? "
	// fmt.Println(sql)

	err := db.Select(&r, sql, code)
	return r, err
}

// Store ...
func (op *journeyOp) Store(
	db *sqlx.DB,
	code string,
	journeyName string,
	journeySchedule int64,
	salesman string,
	sites string,
	questionnaires string,
	signatures int64,
	requireSelfie int64,
	emailTo string,
	activity string,
	// startJourney string,
	// finishJourney string,
	changedAt time.Time,

) (int64, error) {

	createdAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "INSERT INTO journey_plan (code, journey_name, journey_schedule, salesman, sites, questionnaires, signatures, require_selfie, email_to, activity,created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(sql, code, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, createdAt)
	if err != nil {
		return 0, err
	}

	lID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lID, nil
}

// Update ...
func (op *journeyOp) Update(
	db *sqlx.DB,
	code string,
	journeyName string,
	journeySchedule int64,
	salesman string,
	sites string,
	questionnaires string,
	signatures int64,
	requireSelfie int64,
	emailTo string,
	activity string,
	// startJourney string,
	// finishJourney string,
	changedAt time.Time,

) (int64, error) {

	updatedAt := changedAt.Format("2006-01-02 15:04:05")

	var sql = "UPDATE journey_plan SET journey_name = ?,  journey_schedule = ?, salesman = ?, sites = ?, questionnaires = ?, signatures = ?, require_selfie = ?, email_to = ?, activity = ?, updated_at = ? WHERE code = ?"

	_, err := db.Exec(sql, journeyName, journeySchedule, salesman, sites, questionnaires, signatures, requireSelfie, emailTo, activity, updatedAt, code)
	if err != nil {
		return 0, err
	}

	return 0, err
}

// DeleteJourney ...
func (op *journeyOp) DeleteJourney(db *sqlx.DB, code string, changedAt time.Time) ([]*JourneyEntity, error) {
	deletedAt := changedAt.Format("2006-01-02 15:04:05")

	r := []*JourneyEntity{}
	sql := "UPDATE journey_plan SET deleted_at = ? WHERE code = ? "

	_, err := db.Exec(sql, deletedAt, code)
	return r, err
}
