package domain

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"

	"net/url"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/go-playground/validator.v9"
)

var SercretValueMask = "<SecretValue>"

type ConfigRepository interface {
	GetConfigList() (ConfigMap, error)
	GetConfig(ID string) (*Config, error)
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
	TextTemplateString string   `validate:"required"`
	RegexpString       string   `validate:"required"`
	Actions            []string `validate:"unique"`
	CallbackID         string   // should be unique
	Confirm            bool
	URLTemplateString  string `validate:"required"`
	BodyTemplateString string
	Type               string
	Secrets            map[string]string

	Regexp       *regexp.Regexp
	URLTemplate  *template.Template
	BodyTemplate *template.Template
	TextTemplate *template.Template
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

	_, err = template.New("text").Parse(config.TextTemplateString)
	if err != nil {
		sl.ReportError(config.TextTemplateString, "TextTemplateString", "", "", "")
	}

	_, err = regexp.Compile(config.RegexpString)
	if err != nil {
		sl.ReportError(config.RegexpString, "RegexpString", "", "", "")
	}
}

func (c *Config) ExecSecretValueMask(raw string) string {
	result := raw
	for _, v := range c.Secrets {
		result = strings.Replace(result, v, SercretValueMask, -1)
	}
	return result
}

func (c *Config) TextCompile(matched []string) (string, error) {
	textBuf := new(bytes.Buffer)
	err := c.TextTemplate.Execute(textBuf, map[string]interface{}{"matched": matched})
	if err != nil {
		return "", errors.Wrap(err, "Text template failed")
	}
	return textBuf.String(), nil
}

func (c *Config) URLCompile(value string, matched []string, secrets map[string]string) (string, error) {
	urlBuf := new(bytes.Buffer)
	err := c.URLTemplate.Execute(urlBuf, map[string]interface{}{"value": value, "matched": matched, "secrets": secrets})
	if err != nil {
		return "", errors.Wrap(err, "URL template failed")
	}

	parsedURL, err := url.Parse(urlBuf.String())
	if err != nil {
		return "", errors.Wrap(err, "URL parse failed")
	}
	return parsedURL.String(), nil
}

func (c *Config) BodyCompile(value string, matched []string, secrets map[string]string) (string, error) {
	bodyBuf := new(bytes.Buffer)
	err := c.BodyTemplate.Execute(bodyBuf, map[string]interface{}{"value": value, "matched": matched, "secrets": secrets})
	if err != nil {
		return "", errors.Wrap(err, "Body template failed")
	}
	return bodyBuf.String(), nil
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
	c.TextTemplate =
		template.Must(template.New(c.CallbackID + "text").Parse(c.TextTemplateString))
	c.Regexp = regexp.MustCompile(c.RegexpString)
}

func (c *Config) Mask() {
	for k := range c.Secrets {
		c.Secrets[k] = SercretValueMask
	}
}
