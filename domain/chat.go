package domain

type Channels []string

type ChatRepository interface {
	GetChannels() (Channels, error)
}
