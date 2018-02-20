package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"text/template"

	"github.com/juntaki/firestarter/domain"
	"github.com/pkg/errors"
)

type ConfigRepositoryImpl struct {
	currentConfig *domain.ConfigMap
}

func NewConfigRepositoryImpl() *ConfigRepositoryImpl {
	return &ConfigRepositoryImpl{}
}

func (c *ConfigRepositoryImpl) GetConfig() (*domain.ConfigMap, error) {
	if c.currentConfig == nil {
		bytes, err := ioutil.ReadFile("config.json")
		if err != nil {
			return nil, errors.Wrap(err, "Failed to read config file")
		}

		var configMap *domain.ConfigMap
		if err := json.Unmarshal(bytes, configMap); err != nil {
			return nil, errors.Wrap(err, "Config is invalid json")
		}

		for _, config := range *configMap {
			err := c.SetConfig(config)
			if err != nil {
				return nil, errors.Wrap(err, "Config set failed")
			}
		}
	}
	return c.currentConfig, nil
}

func (c *ConfigRepositoryImpl) saveConfig() error {
	bytes, err := json.Marshal(c.currentConfig)
	if err != nil {
		return errors.Wrap(err, "JSON marshal failed")
	}

	err = ioutil.WriteFile("config.json", bytes, 0600)
	if err != nil {
		return errors.Wrap(err, "Failed to write file")
	}

	return nil
}

func (c *ConfigRepositoryImpl) SetConfig(config *domain.Config) (err error) {
	config.BodyTemplate, err =
		template.New(config.CallbackID + "body").Parse(config.BodyTemplateString)
	if err != nil {
		return errors.Wrap(err, "failed to parse body template")
	}
	config.URLTemplate, err =
		template.New(config.CallbackID + "url").Parse(config.URLTemplateString)
	if err != nil {
		return errors.Wrap(err, "failed to parse url template")
	}
	(*c.currentConfig)[config.CallbackID] = config

	err = c.saveConfig()
	if err != nil {
		c.currentConfig = nil
		return err
	}
	return nil
}

func (c *ConfigRepositoryImpl) DeleteConfig(config *domain.Config) error {
	delete(*c.currentConfig, config.CallbackID)
	err := c.saveConfig()
	if err != nil {
		c.currentConfig = nil
		return err
	}
	return nil
}
