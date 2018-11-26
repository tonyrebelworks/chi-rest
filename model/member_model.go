package model

import (
	"chi-rest/model/entity"
	"errors"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

// MemberModelInterface ...
type MemberModelInterface interface {
	All(cursor int) ([]entity.MemberEntity, error)
	Create(me *entity.MemberEntity) error
	FindOneBy(field string, val interface{}) (entity.MemberEntity, error)
}

type memberModel struct {
	DB     *mgo.Session
	DBname string
}

// NewMemberModel new instance of member model.
func NewMemberModel(db *mgo.Session, dbname string) MemberModelInterface {
	return &memberModel{db, dbname}
}

func (m *memberModel) All(cursor int) ([]entity.MemberEntity, error) {
	var err error

	members := []entity.MemberEntity{}
	ds := m.DB.Copy()
	defer ds.Close()

	coll := ds.DB(m.DBname).C("members")
	err = coll.Find(bson.M{}).All(&members)

	return members, err
}

func (m *memberModel) Create(me *entity.MemberEntity) error {
	var err error

	ds := m.DB.Copy()
	defer ds.Close()

	coll := ds.DB(m.DBname).C("members")

	index := mgo.Index{Key: []string{"phone"}, Unique: true}
	err = coll.EnsureIndex(index)

	index = mgo.Index{Key: []string{"member_code"}, Unique: true}
	err = coll.EnsureIndex(index)

	index = mgo.Index{Key: []string{"email"}, Unique: true}
	err = coll.EnsureIndex(index)

	if err = coll.Insert(&me); err != nil {
		if mgo.IsDup(err) {
			return errors.New("Duplicate entry")
		}
	}

	return err
}

func (m *memberModel) FindOneBy(field string, val interface{}) (entity.MemberEntity, error) {
	var err error

	ds := m.DB.Copy()
	defer ds.Close()
	coll := ds.DB(m.DBname).C("members")

	member := entity.MemberEntity{}
	err = coll.Find(bson.M{field: val}).One(&member)

	return member, err
}
