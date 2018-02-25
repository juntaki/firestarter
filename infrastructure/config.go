package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"syscall"

	"github.com/juntaki/firestarter/domain"
	"github.com/pkg/errors"
)

type SaveConfig struct {
	Title              string
	Channels           []string
	Text               string
	RegexpString       string
	Actions            []string
	CallbackID         string
	Confirm            bool
	URLTemplateString  string
	BodyTemplateString string
	Type               string
}

type ConfigRepositoryImpl struct {
	currentConfig map[string]*SaveConfig
	mutex         sync.RWMutex
	loaded        bool
}

func NewConfigRepositoryImpl() *ConfigRepositoryImpl {
	return &ConfigRepositoryImpl{
		currentConfig: make(map[string]*SaveConfig),
		mutex:         sync.RWMutex{},
		loaded:        false,
	}
}

func (c *ConfigRepositoryImpl) GetConfig() (domain.ConfigMap, error) {
	if !c.loaded {
		c.mutex.Lock()
		bytes, err := ioutil.ReadFile("config.json")
		if err != nil {
			pathErr := err.(*os.PathError)
			errno := pathErr.Err.(syscall.Errno)
			if errno == syscall.ENOENT {
				// File not found, return empty config on memory once.
				// SetConfig will make new file.
				c.mutex.Unlock()
				return domain.ConfigMap{}, nil
			}
			// Something happen.
			c.mutex.Unlock()
			return nil, errors.Wrap(err, "Failed to read config file")
		}

		if err := json.Unmarshal(bytes, &c.currentConfig); err != nil {
			c.mutex.Unlock()
			return nil, errors.Wrap(err, "Config is invalid json")
		}
		c.loaded = true
		c.mutex.Unlock()
	}

	c.mutex.RLock()
	defer c.mutex.RUnlock()
	ret := domain.ConfigMap{}
	for _, config := range c.currentConfig {
		ret[config.CallbackID] = c.saveConfigToConfig(config)
	}

	return ret, nil
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

func (c *ConfigRepositoryImpl) loadConfigIfNeeded() error {
	if !c.loaded {
		if _, err := c.GetConfig(); err != nil {
			return errors.Wrap(err, "Load config")
		}
	}
	return nil
}

func (c *ConfigRepositoryImpl) SetConfig(config *domain.Config) error {
	err := c.loadConfigIfNeeded()
	if err != nil {
		return errors.Wrap(err, "Load config on SetConfig")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Update on memory
	bak := c.currentConfig[config.CallbackID]
	c.currentConfig[config.CallbackID] = c.configToSaveConfig(config)

	// Write it to file
	err = c.saveConfig()
	if err != nil {
		// rollback
		c.currentConfig[config.CallbackID] = bak
		return err
	}

	return nil
}

func (c *ConfigRepositoryImpl) IsExist(ID string) (bool, error) {
	config, err := c.GetConfig()
	if err != nil {
		return false, err
	}
	_, ok := config[ID]
	return ok, nil
}

func (c *ConfigRepositoryImpl) DeleteConfig(ID string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.currentConfig, ID)
	err := c.saveConfig()
	if err != nil {
		c.currentConfig = nil
		return err
	}
	return nil
}

// Mapper
func (c *ConfigRepositoryImpl) saveConfigToConfig(saveconfig *SaveConfig) *domain.Config {
	config := &domain.Config{
		Title:              saveconfig.Title,
		CallbackID:         saveconfig.CallbackID,
		Channels:           saveconfig.Channels,
		RegexpString:       saveconfig.RegexpString,
		Text:               saveconfig.Text,
		Actions:            saveconfig.Actions,
		URLTemplateString:  saveconfig.URLTemplateString,
		BodyTemplateString: saveconfig.BodyTemplateString,
		Confirm:            saveconfig.Confirm,
	}
	config.Hydrate()
	return config
}

func (c *ConfigRepositoryImpl) configToSaveConfig(config *domain.Config) *SaveConfig {
	return &SaveConfig{
		Title:              config.Title,
		CallbackID:         config.CallbackID,
		Channels:           config.Channels,
		RegexpString:       config.RegexpString,
		Text:               config.Text,
		Actions:            config.Actions,
		URLTemplateString:  config.URLTemplateString,
		BodyTemplateString: config.BodyTemplateString,
		Confirm:            config.Confirm,
	}
}
