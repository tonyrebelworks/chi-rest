package usecase

import (
	"fmt"
	"time"

	"chi-rest/lib/mysql"
	"chi-rest/lib/utils"

	"github.com/andboson/carbon"
)

// UC default usecase dependencies
type UC struct {
	DB     *mysql.Info
	Config utils.Config
}

// GetData ...
func (uc UC) GetData() error {
	type res struct {
		Name  string `db:"name"`
		Email string `db:"email"`
	}
	r := []res{}
	err := uc.DB.Connect().Select(&r, "SELECT name, email FROM borrowers LIMIT 10")
	defer uc.DB.Close()

	fmt.Println(r)

	return err
}

func today() time.Time {
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// now := time.Now().In(loc)

	// return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 00, time.UTC)
	cb := carbon.Now()
	cb.SetTZ("UTC")

	return cb.Time
}
