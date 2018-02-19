package slack

import (
	"fmt"
	"strings"

	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/utils"

	"github.com/shopspring/decimal"

	"github.com/nlopes/slack"
	//log "github.com/sirupsen/logrus"
)

type SlackLoggerInterface interface {
	Init(channelName string) error
	PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string) error
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
func (s *SlackLogger) Init(channelName string) error {
	channels, err := s.client.GetChannels(false)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		if channel.Name == channelName {
			s.channelID = channel.ID
		}
	}

	_, _, err = s.client.PostMessage(s.channelID, "Bot is online.", slack.PostMessageParameters{})
	if err != nil {
		return err
	}

	return nil
}

// PostTrade logs a trade to Slack
func (s *SlackLogger) PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string) error {
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
		message = fmt.Sprintf("%s *`%s`* - *%s/%s* - %s @ %s - *(%s%%)*", emoji, strings.ToUpper(action), strings.ToUpper(market), strings.ToUpper(base), quantity.String(), limit.String(), profitLoss.String())
	} else {
		emoji := ":shopping_bags:"
		message = fmt.Sprintf("%s *`%s`* - *%s/%s* - %s @ %s", emoji, strings.ToUpper(action), strings.ToUpper(market), strings.ToUpper(base), quantity.String(), limit.String())
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			{
				Color:      "good",
				Pretext:    message,
				MarkdownIn: []string{"pretext"},
			},
		},
	}

	_, _, err := s.client.PostMessage(s.channelID, "", params)
	if err != nil {
		return err
	}

	return nil
}
