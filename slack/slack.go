package slack

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const SlackAPIToken string = "xoxp-315598160006-314865968613-314739060532-5ede8aca6ad11fcfc9cf3bdded48b9c7"

var SlackAPI *slack.Client
var SlackChannelID string

func Init() {
	SlackAPI = slack.New(SlackAPIToken)
	// SlackAPI.SetDebug(true)

	channels, err := SlackAPI.GetChannels(false)
	if err != nil {
		log.Error(err)
	}

	for _, channel := range channels {
		if channel.Name == "trades" {
			SlackChannelID = channel.ID
		}
	}

	_, _, err = SlackAPI.PostMessage(SlackChannelID, "Bot is online.", slack.PostMessageParameters{})
	if err != nil {
		log.Error(err)
	}
}

func PostTrade(action string, limit decimal.Decimal, quantity decimal.Decimal, base string, market string) {
	message := fmt.Sprintf("%s %s/%s: %s at %s", action, market, base, quantity.String(), limit.String())
	SlackAPI.PostMessage(SlackChannelID, message, slack.PostMessageParameters{})
}
