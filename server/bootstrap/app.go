package bootstrap

import (
	"chi-rest/lib/mysql"
	"chi-rest/lib/utils"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	idTranslations "gopkg.in/go-playground/validator.v9/translations/id"
)

// App instance of the skeleton
type App struct {
	R         *chi.Mux
	DB        *mysql.Info
	Config    utils.Config
	Validator *Validator
}

// SetupDB setup database connection
func SetupDB(config utils.Config) *mysql.Info {
	db := mysql.Info{
		Host: config.GetString("database.mysql.host"),
		DB:   config.GetString("database.mysql.db"),
		User: config.GetString("database.mysql.user"),
		Pass: config.GetString("database.mysql.pass"),
		Loc:  time.UTC,
	}

	return &db
}

// Validator set validator instance
type Validator struct {
	Driver     *validator.Validate
	Uni        *ut.UniversalTranslator
	Translator ut.Translator
}

// SetupValidator create new instance of validator driver
func SetupValidator(config utils.Config) *Validator {
	en := en.New()
	id := id.New()
	uni := ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver := validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	var trans ut.Translator
	switch config.GetString("app.locale") {
	case "en":
		trans = transEN
	case "id":
		trans = transID
	}

	return &Validator{Driver: validatorDriver, Uni: uni, Translator: trans}
}
