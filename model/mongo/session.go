package mongo

import (
	"time"

	"github.com/globalsign/mgo"
)

// Info ...
type Info struct {
	Host string
	Db   string
	User string
	Pass string
}

// Connect make a connection to mongo database.
func (i *Info) Connect() (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     []string{i.Host},
		Timeout:   60 * time.Second,
		Database:  i.Db,
		Username:  i.User,
		Password:  i.Pass,
		Source:    "admin",
		Mechanism: "SCRAM-SHA-1",
	}

	sess, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}

	sess.SetMode(mgo.Monotonic, true)

	return sess, err
}
