package usecase

import (
	"time"

	"chi-rest/bootstrap"

	"github.com/andboson/carbon"
)

// UC default usecase dependencies
type UC struct {
	*bootstrap.App
}

func today() time.Time {
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)

	// return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 00, time.UTC)
	cb := carbon.Now()
	cb.SetTZ("UTC")

	return cb.Time
}
