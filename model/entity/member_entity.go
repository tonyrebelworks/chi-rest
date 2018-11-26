package entity

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// MemberEntity ...
type MemberEntity struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Code      string        `bson:"member_code"`
	Fname     string        `bson:"fname"`
	Lname     string        `bson:"lname"`
	Email     string        `bson:"email"`
	Phone     string        `bson:"phone"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
