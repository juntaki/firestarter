package infrastructure

import (
	"github.com/juntaki/firestarter/domain"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

type ChatRepositorySlackImpl struct {
	API *slack.Client
}

func (c *ChatRepositorySlackImpl) GetChannels() (domain.Channels, error) {
	channels, err := c.API.GetChannels(true)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get channels")
	}

	ret := make(domain.Channels, len(channels))
	for i, c := range channels {
		ret[i] = c.Name
	}
	return ret, nil
}
