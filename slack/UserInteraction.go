package slack

import (
	"encoding/json"
	"errors"
	"github.com/ndphu/swd-commons/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	emailLookupBaseUrl = "https://slack.com/api/users.lookupByEmail"
	postMessageUrl     = "https://slack.com/api/chat.postMessage"
	emailInvitation    = "https://slack.com/api/users.admin.invite"
	botAccessToken     = os.Getenv("SLACK_BOT_ACCESS_TOKEN")
	appAccessToken     = os.Getenv("SLACK_OAUTH_ACCESS_TOKEN")
)

func SendSlackInvitation(email string) (error) {
	req, err := http.NewRequest("GET", emailInvitation, nil);
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("token", appAccessToken)
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	if code, payload, err := utils.MakeRequest(req); err != nil {
		return err
	} else {
		if code != 200 {
			log.Println("[SLACK]", "Fail to send invitation. Server response", code, string(payload))
			return errors.New("INVALID_STATUS_CODE_" + strconv.Itoa(code))
		} else {
			sr := UserInviteResponse{}
			if err := json.Unmarshal(payload, &sr); err != nil {
				return err
			}
			if !sr.OK {
				log.Println("[SLACK]", "Invite fail, server response:", string(payload))
				return errors.New(strings.ToUpper(sr.Error))
			}
			return nil
		}
	}
}

func LookupUserIdByEmail(email string) (*User, error) {
	req, err := http.NewRequest("GET", emailLookupBaseUrl, nil);
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("token", botAccessToken)
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	if code, payload, err := utils.MakeRequest(req); err != nil {
		return nil, err
	} else {
		if code != 200 {
			log.Println("[SLACK]", "LookupUserIdByEmail failed. Server response", code, string(payload))
			return nil, errors.New("INVALID_STATUS_CODE_" + strconv.Itoa(code))
		} else {
			sr := UserSearchResponse{}
			if err := json.Unmarshal(payload, &sr); err != nil {
				return nil, err
			}
			if !sr.OK {
				log.Println("[SLACK]", "LookupUserIdByEmail failed. Server response", code, string(payload))
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
	status, data, err := utils.PostJsonModel(postMessageUrl, "Bearer "+botAccessToken, mesg)
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
