package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/juntaki/firestarter/domain"
	proto "github.com/juntaki/firestarter/proto"
	"github.com/pkg/errors"
)

type DummyConfigRepository struct {
	domain.ConfigRepository
	dummyGetConfigList func() (domain.ConfigMap, error)
	dummyGetConfig     func(ID string) (*domain.Config, error)
	dummySetConfig     func(*domain.Config) error
	dummyIsExist       func(ID string) (bool, error)
	dummyDeleteConfig  func(ID string) error
}

func (d *DummyConfigRepository) GetConfigList() (domain.ConfigMap, error) {
	return d.dummyGetConfigList()
}
func (d *DummyConfigRepository) GetConfig(ID string) (*domain.Config, error) {
	return d.dummyGetConfig(ID)
}
func (d *DummyConfigRepository) SetConfig(c *domain.Config) error {
	return d.dummySetConfig(c)
}
func (d *DummyConfigRepository) IsExist(ID string) (bool, error) {
	return d.dummyIsExist(ID)
}
func (d *DummyConfigRepository) DeleteConfig(ID string) error {
	return d.dummyDeleteConfig(ID)
}

type DummyChatRepository struct {
	domain.ChatRepository
	dummyGetChannels func() (domain.Channels, error)
}

func (d *DummyChatRepository) GetChannels() (domain.Channels, error) {
	return d.dummyGetChannels()
}

func TestAdminAPI_GetConfig(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		ctx     context.Context
		request *proto.GetConfigRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Config
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyGetConfig: func(ID string) (*domain.Config, error) {
						c := &domain.Config{
							CallbackID:         "callbackid",
							Title:              "title",
							Channels:           []string{"channel"},
							TextTemplateString: "text",
							RegexpString:       "regexp",
							Actions:            []string{""},
							Confirm:            true,
							URLTemplateString:  "url",
							BodyTemplateString: "body",
							Type:               "unused",
							Secrets: map[string]string{
								"key": "value",
							},
						}
						c.Hydrate()
						return c, nil
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx:     context.Background(),
				request: &proto.GetConfigRequest{},
			},
			want: &proto.Config{
				ID:           "callbackid",
				Title:        "title",
				Channels:     []string{"channel"},
				TextTemplate: "text",
				Regexp:       "regexp",
				Actions:      []string{""},
				Confirm:      true,
				URLTemplate:  "url",
				BodyTemplate: "body",
				Secrets: []*proto.Secret{
					&proto.Secret{
						Key:   "key",
						Value: "<SecretValue>",
					},
				},
			},
		},
		{
			name: "fail",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyGetConfig: func(ID string) (*domain.Config, error) {
						return nil, errors.New("failed")
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx:     context.Background(),
				request: &proto.GetConfigRequest{},
			},
			want:    &proto.Config{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			got, err := a.GetConfig(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminAPI.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_GetConfigList(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		ctx     context.Context
		request *proto.GetConfigListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ConfigList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			got, err := a.GetConfigList(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminAPI.GetConfigList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.GetConfigList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_SetConfig(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		ctx      context.Context
		pbconfig *proto.Config
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.SetConfigResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyIsExist: func(ID string) (bool, error) {
						return false, nil
					},
					dummySetConfig: func(c *domain.Config) error {
						return nil
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx: context.Background(),
				pbconfig: &proto.Config{
					ID:           "",
					Title:        "title",
					Channels:     []string{"channel"},
					TextTemplate: "text",
					Regexp:       "regexp",
					Actions:      []string{""},
					Confirm:      true,
					URLTemplate:  "url",
					BodyTemplate: "body",
					Secrets: []*proto.Secret{
						&proto.Secret{
							Key:   "key",
							Value: "<SecretValue>",
						},
					},
				},
			},
			want: &proto.SetConfigResponse{},
		},
		{
			name: "validation error: regexp",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyIsExist: func(ID string) (bool, error) {
						return false, nil
					},
					dummySetConfig: func(c *domain.Config) error {
						return nil
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx: context.Background(),
				pbconfig: &proto.Config{
					ID:           "",
					Title:        "title",
					Channels:     []string{"channel"},
					TextTemplate: "text",
					Regexp:       "regexp***",
					Actions:      []string{""},
					Confirm:      true,
					URLTemplate:  "url",
					BodyTemplate: "body",
					Secrets: []*proto.Secret{
						&proto.Secret{
							Key:   "key",
							Value: "<SecretValue>",
						},
					},
				},
			},
			want:    &proto.SetConfigResponse{},
			wantErr: true,
		},
		{
			name: "validation error: template",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyIsExist: func(ID string) (bool, error) {
						return false, nil
					},
					dummySetConfig: func(c *domain.Config) error {
						return nil
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx: context.Background(),
				pbconfig: &proto.Config{
					ID:           "",
					Title:        "title",
					Channels:     []string{"channel"},
					TextTemplate: "text",
					Regexp:       "regexp",
					Actions:      []string{""},
					Confirm:      true,
					URLTemplate:  "url",
					BodyTemplate: "body{{{{",
					Secrets: []*proto.Secret{
						&proto.Secret{
							Key:   "key",
							Value: "<SecretValue>",
						},
					},
				},
			},
			want:    &proto.SetConfigResponse{},
			wantErr: true,
		},
		{
			name: "validation error: not exist ID",
			fields: fields{
				ConfigRepository: &DummyConfigRepository{
					dummyIsExist: func(ID string) (bool, error) {
						return false, nil
					},
					dummySetConfig: func(c *domain.Config) error {
						return nil
					},
				},
				ChatRepository: &DummyChatRepository{},
				Validator:      domain.NewValidator(),
			},
			args: args{
				ctx: context.Background(),
				pbconfig: &proto.Config{
					ID:           "dummyid",
					Title:        "title",
					Channels:     []string{"channel"},
					TextTemplate: "text",
					Regexp:       "regexp",
					Actions:      []string{""},
					Confirm:      true,
					URLTemplate:  "url",
					BodyTemplate: "body",
					Secrets: []*proto.Secret{
						&proto.Secret{
							Key:   "key",
							Value: "<SecretValue>",
						},
					},
				},
			},
			want:    &proto.SetConfigResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			got, err := a.SetConfig(tt.args.ctx, tt.args.pbconfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminAPI.SetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.SetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_DeleteConfig(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		ctx context.Context
		r   *proto.DeleteConfigRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.DeleteConfigResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			got, err := a.DeleteConfig(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminAPI.DeleteConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.DeleteConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_GetChannels(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		ctx context.Context
		req *proto.GetChannelsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.Channels
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			got, err := a.GetChannels(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminAPI.GetChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.GetChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_pbConfigToConfig(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		pbconfig *proto.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			if got := a.pbConfigToConfig(tt.args.pbconfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.pbConfigToConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminAPI_configToPbConfig(t *testing.T) {
	type fields struct {
		ConfigRepository domain.ConfigRepository
		ChatRepository   domain.ChatRepository
		Validator        *domain.Validator
	}
	type args struct {
		config *domain.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *proto.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdminAPI{
				ConfigRepository: tt.fields.ConfigRepository,
				ChatRepository:   tt.fields.ChatRepository,
				Validator:        tt.fields.Validator,
			}
			if got := a.configToPbConfig(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminAPI.configToPbConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
