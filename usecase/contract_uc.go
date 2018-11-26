package usecase

import (
	"chi-rest/lib"
	"time"

	"github.com/globalsign/mgo"
)

// UC default usecase dependencies
type UC struct {
	db  *mgo.Session
	cfg lib.Config
}

func today() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 00, time.UTC)
}
