package slack

import (
	"encoding/json"
	"errors"
	"github.com/ndphu/swd-commons/utils"
	"log"
	"net/url"
	"os"
)

var (
	emailLookupBaseUrl = "https://slack.com/api/users.lookupByEmail"
	postMessageUrl     = "https://slack.com/api/chat.postMessage"
	accessToken        = os.Getenv("SLACK_BOT_ACCESS_TOKEN")
)

func LookupUserIdByEmail(email string) (*User, error) {
	if code, payload, err := utils.GetWithAccessToken(emailLookupBaseUrl+url.QueryEscape(email), accessToken); err != nil {
		return nil, err
	} else {
		if code != 200 {
			return nil, errors.New("INVALID_STATUS_CODE")
		} else {
			sr := UserSearchRespond{}
			if err := json.Unmarshal(payload, &sr); err != nil {
				return nil, err
			}
			if !sr.OK {
				return nil, errors.New("SEARCH_FAIL")
			}
			return &sr.User, nil
		}
	}
}

func SendMessageToUser(userId, message string) error {
	mesg := ReplyMessage{
		Text:    message,
		Channel: userId,
		AsUser:  true,
	}
	status, data, err := utils.PostJsonModel(postMessageUrl, "Bearer "+accessToken, mesg)
	if err != nil {
		log.Println("[SLACKBOT]", "Fail to post reply message, server return status:", status)
		return err
	}
	if status != 200 {
		log.Println("[SLACKBOT]", "Fail to post reply message, server return status:", status)
		return errors.New("INVALID_STATUS_CODE")
	}
	log.Printf("[SLACKBOT] Message posted successfully: status:%d, resp:%s\n", status, string(data))
	return nil
}
