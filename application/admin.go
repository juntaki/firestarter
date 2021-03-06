package application

import (
	"context"

	"sort"

	"github.com/juntaki/firestarter/domain"
	proto "github.com/juntaki/firestarter/proto"
	"github.com/twitchtv/twirp"
)

type AdminAPI struct {
	ConfigRepository domain.ConfigRepository
	ChatRepository   domain.ChatRepository
	Validator        *domain.Validator
}

func NewAdminAPI(configRepository domain.ConfigRepository, chatRepository domain.ChatRepository) *AdminAPI {
	return &AdminAPI{
		ConfigRepository: configRepository,
		ChatRepository:   chatRepository,
		Validator:        domain.NewValidator(),
	}
}

func (a *AdminAPI) GetConfig(ctx context.Context, request *proto.GetConfigRequest) (*proto.Config, error) {
	config, err := a.ConfigRepository.GetConfig(request.ID)
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

	// sort by id
	keys := make([]string, 0)
	for k := range config {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := &proto.ConfigList{}
	for _, k := range keys {
		config[k].Mask()
		result.Config = append(result.Config, a.configToPbConfig(config[k]))
	}

	return result, nil
}

func (a *AdminAPI) SetConfig(ctx context.Context, pbconfig *proto.Config) (*proto.SetConfigResponse, error) {
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
	err := a.ConfigRepository.DeleteConfig(r.ID)
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
	config := &domain.Config{
		Title:              pbconfig.Title,
		CallbackID:         pbconfig.ID,
		Channels:           pbconfig.Channels,
		RegexpString:       pbconfig.Regexp,
		TextTemplateString: pbconfig.TextTemplate,
		Actions:            pbconfig.Actions,
		URLTemplateString:  pbconfig.URLTemplate,
		BodyTemplateString: pbconfig.BodyTemplate,
		Confirm:            pbconfig.Confirm,
		Secrets:            make(map[string]string),
	}

	for _, s := range pbconfig.Secrets {
		config.Secrets[s.Key] = s.Value
	}

	return config
}

func (a *AdminAPI) configToPbConfig(config *domain.Config) *proto.Config {
	pbconfig := &proto.Config{
		Title:        config.Title,
		ID:           config.CallbackID,
		Channels:     config.Channels,
		Regexp:       config.RegexpString,
		TextTemplate: config.TextTemplateString,
		Actions:      config.Actions,
		URLTemplate:  config.URLTemplateString,
		BodyTemplate: config.BodyTemplateString,
		Confirm:      config.Confirm,
		Secrets:      make([]*proto.Secret, 0),
	}

	for k, v := range config.Secrets {
		pbconfig.Secrets = append(pbconfig.Secrets, &proto.Secret{Key: k, Value: v})
	}
	return pbconfig
}
