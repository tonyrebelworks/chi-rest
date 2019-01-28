package usecase

import (
	"chi-rest/lib"
	"chi-rest/model"
	"chi-rest/model/entity"
	"chi-rest/server/request"
	"strings"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
	"golang.org/x/crypto/bcrypt"
)

// memberUsecase add dependencies of usecases
type memberUsecase struct {
	UC
}

// MemberUsercaseInterface ...
type MemberUsercaseInterface interface {
	FindByID(value string) (entity.MemberEntity, error)
	FindByPhone(value string) (entity.MemberEntity, error)
	Register(req request.RegisterRequest) error
}

// NewMemberUsecase ...
func NewMemberUsecase(db *mgo.Session, cfg lib.Config) MemberUsercaseInterface {
	return &memberUsecase{UC{db, cfg}}
}

func (uc *memberUsecase) FindByID(value string) (entity.MemberEntity, error) {
	var (
		err error
		me  entity.MemberEntity
	)

	mod := model.NewMemberModel(uc.db, uc.cfg.GetString("database.mongo.db"))
	me, err = mod.FindOneBy("_id", bson.ObjectIdHex(value))

	return me, err
}

// FindByPhone ...
func (uc *memberUsecase) FindByPhone(value string) (entity.MemberEntity, error) {
	var (
		err error
		me  entity.MemberEntity
	)

	mod := model.NewMemberModel(uc.db, uc.cfg.GetString("database.mongo.db"))
	me, err = mod.FindOneBy("phone", value)
	me.CreatedAt = me.CreatedAt.UTC()
	return me, err
}

// Register ...
func (uc *memberUsecase) Register(req request.RegisterRequest) error {
	var err error

	// 1. need a validation here!

	password := []byte(string(req.Password))
	password, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	me := &entity.MemberEntity{
		Code:      strings.ToUpper(lib.RandStringBytesMaskImprSrc(10)),
		Fname:     req.Fname,
		Lname:     req.Lname,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(password),
		CreatedAt: today(),
	}

	m := model.NewMemberModel(uc.db, uc.cfg.GetString("database.mongo.db"))
	err = m.Create(me)

	return err
}
