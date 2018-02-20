package domain

import (
	"regexp"
	"text/template"
)

type ConfigRepository interface {
	GetConfig() (*ConfigMap, error)
	SetConfig(*Config) error
	DeleteConfig(*Config) error
}

type ConfigMap map[string]*Config

type Config struct {
	Regexp             *regexp.Regexp
	Channels           []string
	Text               string
	Actions            []string
	CallbackID         string
	Type               string
	URLTemplate        *template.Template
	URLTemplateString  string
	BodyTemplate       *template.Template
	BodyTemplateString string
	Confirm            bool
}

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
