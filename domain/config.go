package domain

import (
	"regexp"
	"text/template"

	"github.com/rs/xid"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigRepository interface {
	GetConfig() (ConfigMap, error)
	SetConfig(*Config) error
	IsExist(ID string) (bool, error)
	DeleteConfig(ID string) error
}

type ConfigMap map[string]*Config

func (q *ConfigMap) FindMatched(channel, text string) *Config {
	for _, config := range *q {
		for _, ch := range config.Channels {
			if ch == channel && config.Regexp.MatchString(text) {
				return config
			}
		}
	}
	return nil
}

func (q *ConfigMap) FindByCallbackID(callbackID string) *Config {
	return (*q)[callbackID]
}

type Config struct {
	Title              string   // for admin
	Channels           []string `validate:"unique,required,dive,required"`
	Text               string   `validate:"required"`
	RegexpString       string   `validate:"required"`
	Actions            []string `validate:"unique"`
	CallbackID         string   // should be unique
	Confirm            bool
	URLTemplateString  string `validate:"required"`
	BodyTemplateString string
	Type               string

	Regexp       *regexp.Regexp
	URLTemplate  *template.Template
	BodyTemplate *template.Template
}

func ConfigValidator(sl validator.StructLevel) {
	config := sl.Current().Interface().(Config)

	_, err := template.New("body").Parse(config.BodyTemplateString)
	if err != nil {
		sl.ReportError(config.BodyTemplateString, "BodyTemplateString", "", "", "")
	}

	_, err = template.New("url").Parse(config.URLTemplateString)
	if err != nil {
		sl.ReportError(config.URLTemplateString, "URLTemplateString", "", "", "")
	}

	_, err = regexp.Compile(config.RegexpString)
	if err != nil {
		sl.ReportError(config.RegexpString, "RegexpString", "", "", "")
	}
}

func (c *Config) Hydrate() {
	// Assign callback ID, new config
	if c.CallbackID == "" {
		c.CallbackID = xid.New().String()
	}

	// Title should be filled by something.
	if c.Title == "" {
		c.Title = c.CallbackID
	}

	// Compile from string
	c.BodyTemplate =
		template.Must(template.New(c.CallbackID + "body").Parse(c.BodyTemplateString))
	c.URLTemplate =
		template.Must(template.New(c.CallbackID + "url").Parse(c.URLTemplateString))
	c.Regexp = regexp.MustCompile(c.RegexpString)
}
