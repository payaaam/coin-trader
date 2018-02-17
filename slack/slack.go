package slack

import (
	"fmt"
	"strings"

	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/utils"

	"github.com/shopspring/decimal"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const slackAPIToken string = "xoxp-315598160006-314865968613-314739060532-5ede8aca6ad11fcfc9cf3bdded48b9c7"

type SlackLoggerInterface interface {
	Init(channelName string)
	PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string)
}

// Logger logs to a specified Slack channel
type SlackLogger struct {
	client    *slack.Client
	channelID string
}

// NewSlackLogger creates a new Slack logger
func NewSlackLogger(slackToken string) *SlackLogger {
	return &SlackLogger{
		client: slack.New(slackToken),
	}
}

// Init initializes the Slack logger to the specific channel
func (s *SlackLogger) Init(channelName string) {
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
func (s *SlackLogger) PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string) {
	var emoji string
	var message string
	if action == orders.SellOrder {
		// Calculate profit/loss
		oneHundo := utils.StringToDecimal("100")
		buyPrice := utils.StringToDecimal("0.01").Div(quantity)
		profitLoss := limit.Sub(buyPrice).Div(buyPrice).Mul(oneHundo).Round(3)
		if profitLoss.Sign() > 0 {
			emoji = ":white_check_mark:"
		} else {
			emoji = ":x:"
		}
		message = fmt.Sprintf("%s *%s %s/%s* @ %s (%s%%)", emoji, action, strings.ToUpper(market), strings.ToUpper(base), limit.String(), profitLoss.String())
	} else {
		emoji := ":new:"
		message = fmt.Sprintf("%s *%s %s/%s* @ %s", emoji, action, strings.ToUpper(market), strings.ToUpper(base), limit.String())
	}
	s.client.PostMessage(s.channelID, message, slack.PostMessageParameters{})
}
