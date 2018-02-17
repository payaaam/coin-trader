package slack

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const slackAPIToken string = "xoxp-315598160006-314865968613-314739060532-5ede8aca6ad11fcfc9cf3bdded48b9c7"

// Logger logs to a specified Slack channel
type Logger struct {
	client    *slack.Client
	channelID string
}

// NewLogger creates a new Slack logger
func NewLogger(slackToken string) *Logger {
	return &Logger{
		client: slack.New(slackToken),
	}
}

// Init initializes the Slack logger to the specific channel
func (s *Logger) Init(channelName string) {
	// s.client.SetDebug(true)

	channels, err := s.client.GetChannels(false)
	if err != nil {
		log.Error(err)
	}

	for _, channel := range channels {
		if channel.Name == channelName {
			s.channelID = channel.ID
		}
	}

	_, _, err = s.client.PostMessage(s.channelID, "Bot is online.", slack.PostMessageParameters{})
	if err != nil {
		log.Error(err)
	}
}

// PostTrade logs a trade to Slack
func (s *Logger) PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string) {
	message := fmt.Sprintf("%s %s/%s: %s at %s", action, market, base, quantity.String(), limit.String())
	s.client.PostMessage(s.channelID, message, slack.PostMessageParameters{})
}
