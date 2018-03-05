package application

import (
	"context"

	"github.com/juntaki/firestarter/domain"
	proto "github.com/juntaki/firestarter/proto"
	"github.com/k0kubun/pp"
	"github.com/twitchtv/twirp"
)

type AdminAPI struct {
	ConfigRepository domain.ConfigRepository
	ChatRepository   domain.ChatRepository
	Validator        *domain.Validator
}

func (a *AdminAPI) GetConfig(ctx context.Context, request *proto.GetConfigRequest) (*proto.Config, error) {
	config, err := a.ConfigRepository.GetConfig(request.CallbackID)
	if err != nil {
		return &proto.Config{}, err
	}

	config.Mask()
	return a.configToPbConfig(config), nil
}

func (a *AdminAPI) GetConfigList(ctx context.Context, request *proto.GetConfigListRequest) (*proto.ConfigList, error) {
	config, err := a.ConfigRepository.GetConfigList()
	if err != nil {
		return &proto.ConfigList{}, err
	}

	result := &proto.ConfigList{}
	for _, v := range config {
		v.Mask()
		result.Config = append(result.Config, a.configToPbConfig(v))
	}

	return result, nil
}

func (a *AdminAPI) SetConfig(ctx context.Context, pbconfig *proto.Config) (*proto.SetConfigResponse, error) {
	pp.Println(pbconfig)
	config := a.pbConfigToConfig(pbconfig)
	err := a.Validator.ValidateConfig(config)
	if err != nil {
		return &proto.SetConfigResponse{},
			twirp.InvalidArgumentError(
				"config",
				err.Error(),
			)
	}

	if config.CallbackID != "" {
		if exist, err := a.ConfigRepository.IsExist(config.CallbackID); err != nil {
			return &proto.SetConfigResponse{}, err
		} else if !exist {
			return &proto.SetConfigResponse{}, twirp.InvalidArgumentError(
				"config.id", "Malformed request")
		}
	}

	config.Hydrate()
	err = a.ConfigRepository.SetConfig(config)
	return &proto.SetConfigResponse{}, err
}

func (a *AdminAPI) DeleteConfig(ctx context.Context, r *proto.DeleteConfigRequest) (*proto.DeleteConfigResponse, error) {
	err := a.ConfigRepository.DeleteConfig(r.CallbackID)
	return &proto.DeleteConfigResponse{}, err
}

func (a *AdminAPI) GetChannels(ctx context.Context, req *proto.GetChannelsRequest) (*proto.Channels, error) {
	ch, err := a.ChatRepository.GetChannels()
	if err != nil {
		return &proto.Channels{}, err
	}
	return &proto.Channels{
		List: ch,
	}, nil
}

// Mapper
func (a *AdminAPI) pbConfigToConfig(pbconfig *proto.Config) *domain.Config {
	return &domain.Config{
		Title:              pbconfig.Title,
		CallbackID:         pbconfig.CallbackID,
		Channels:           pbconfig.Channels,
		RegexpString:       pbconfig.Regexp,
		Text:               pbconfig.Text,
		Actions:            pbconfig.Actions,
		URLTemplateString:  pbconfig.URLTemplate,
		BodyTemplateString: pbconfig.BodyTemplate,
		Confirm:            pbconfig.Confirm,
		Secrets:            pbconfig.Secrets,
	}
}

func (a *AdminAPI) configToPbConfig(config *domain.Config) *proto.Config {
	return &proto.Config{
		Title:        config.Title,
		CallbackID:   config.CallbackID,
		Channels:     config.Channels,
		Regexp:       config.RegexpString,
		Text:         config.Text,
		Actions:      config.Actions,
		URLTemplate:  config.URLTemplateString,
		BodyTemplate: config.BodyTemplateString,
		Confirm:      config.Confirm,
		Secrets:      config.Secrets,
	}
}
