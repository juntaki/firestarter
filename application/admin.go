package application

import (
	"context"

	"github.com/juntaki/firestarter/domain"
	"github.com/juntaki/firestarter/proto"
)

type AdminAPI struct {
	ConfigRepository domain.ConfigRepository
}

func (a *AdminAPI) GetConfig(ctx context.Context, request *proto.GetConfigListRequest) (*proto.ConfigList, error) {
	config, err := a.ConfigRepository.GetConfig()
	if err != nil {
		return &proto.ConfigList{}, err
	}

	result := &proto.ConfigList{}
	for _, v := range *config {
		result.Config = append(result.Config, a.configToPbConfig(v))
	}
	return result, nil
}

func (a *AdminAPI) SetConfig(ctx context.Context, pbconfig *proto.Config) (*proto.SetConfigResponse, error) {
	config := a.pbConfigToConfig(pbconfig)
	err := a.ConfigRepository.SetConfig(config)
	return &proto.SetConfigResponse{}, err
}

func (a *AdminAPI) DeleteConfig(ctx context.Context, pbconfig *proto.Config) (*proto.DeleteConfigResponse, error) {
	config := a.pbConfigToConfig(pbconfig)
	err := a.ConfigRepository.DeleteConfig(config)
	return &proto.DeleteConfigResponse{}, err
}

// Mapper
func (a *AdminAPI) pbConfigToConfig(pbconfig *proto.Config) *domain.Config {
	return &domain.Config{
		CallbackID:         pbconfig.CallbackID,
		Channels:           pbconfig.Channels,
		Text:               pbconfig.Text,
		Actions:            pbconfig.Actions,
		URLTemplateString:  pbconfig.URLTemplate,
		BodyTemplateString: pbconfig.BodyTemplate,
		Confirm:            pbconfig.Confirm,
	}
}

func (a *AdminAPI) configToPbConfig(config *domain.Config) *proto.Config {
	return &proto.Config{
		CallbackID:   config.CallbackID,
		Channels:     config.Channels,
		Text:         config.Text,
		Actions:      config.Actions,
		URLTemplate:  config.URLTemplateString,
		BodyTemplate: config.BodyTemplateString,
		Confirm:      config.Confirm,
	}
}
