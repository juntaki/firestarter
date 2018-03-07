package infrastructure

import (
	"reflect"
	"sync"
	"testing"

	"github.com/juntaki/firestarter/domain"
	"go.uber.org/zap"
)

func TestNewConfigRepositoryImpl(t *testing.T) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("logger initialize failed")
	}
	logger := zapLogger.Sugar()

	tests := []struct {
		name string
		want *ConfigRepositoryImpl
	}{
		{
			name: "create",
			want: &ConfigRepositoryImpl{
				currentConfig: make(map[string]*SaveConfig),
				mutex:         &sync.RWMutex{},
				loaded:        false,
				configFile:    "config/config.json",
				logger:        logger,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigRepositoryImpl(logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigRepositoryImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigRepositoryImpl_GetConfigList(t *testing.T) {
	type fields struct {
		currentConfig map[string]*SaveConfig
		mutex         *sync.RWMutex
		loaded        bool
		configFile    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.ConfigMap
		wantErr bool
	}{
		{
			name: "get from json",
			fields: fields{
				currentConfig: make(map[string]*SaveConfig),
				mutex:         &sync.RWMutex{},
				loaded:        false,
				configFile:    "test/config.json",
			},
			want: func() domain.ConfigMap {
				configMap := domain.ConfigMap{}
				configMap["ba8oiiei1gbjr0ucqbo0"] = &domain.Config{
					Title:              "Test",
					Channels:           []string{"bottest"},
					RegexpString:       "^deploy$",
					Actions:            []string{"master", "branch"},
					CallbackID:         "ba8oiiei1gbjr0ucqbo0",
					Confirm:            true,
					URLTemplateString:  "http://httpbin.org/post?test={{index .matched 1}}",
					BodyTemplateString: "{ value : {{index .matched 1}} }",
					TextTemplateString: "Deploy app",
					Secrets:            map[string]string{},
				}
				configMap["ba8oiiei1gbjr0ucqbo0"].Hydrate()
				return configMap
			}(),
			wantErr: false,
		},
		{
			name: "get from memory",
			fields: fields{
				currentConfig: make(map[string]*SaveConfig),
				mutex:         &sync.RWMutex{},
				loaded:        true,
				configFile:    "test/config.json",
			},
			want:    domain.ConfigMap{},
			wantErr: false,
		},
		{
			name: "get from file (not found)",
			fields: fields{
				currentConfig: make(map[string]*SaveConfig),
				mutex:         &sync.RWMutex{},
				loaded:        false,
				configFile:    "test/config_not_found.json",
			},
			want:    domain.ConfigMap{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConfigRepositoryImpl{
				currentConfig: tt.fields.currentConfig,
				mutex:         tt.fields.mutex,
				loaded:        tt.fields.loaded,
				configFile:    tt.fields.configFile,
			}
			got, err := c.GetConfigList()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigRepositoryImpl.GetConfigList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigRepositoryImpl.GetConfigList() = %v, want %v", got, tt.want)
			}
		})
	}
}
