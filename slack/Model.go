package slack

import "time"

var (
	SlackMessageColorSuccess = "#36a64f"
)

type BaseRequest struct {
	Type string `json:"type"`
}

type ChallengeRequest struct {
	BaseRequest
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type EventCallbackRequest struct {
	BaseRequest
	Event       Event     `json:"event"`
	EventId     string    `json:"event_id"`
	EventTime   time.Time `json:"event_time"`
	AuthedUsers []string  `json:"authed_users"`
	TeamId      string    `json:"team_id"`
	ApiAppId    string    `json:"api_app_id"`
	Token       string    `json:"token"`
}

type Event struct {
	ClientMessageId string `json:"client_message_id"`
	Subtype         string `json:"subtype"`
	Type            string `json:"type"`
	Text            string `json:"text"`
	User            string `json:"user"`
	Ts              string `json:"ts"`
	Channel         string `json:"channel"`
	EventTs         string `json:"event_ts"`
	ChannelType     string `json:"channel_type"`
}

type ReplyMessage struct {
	Text        string       `json:"text"`
	Channel     string       `json:"channel"`
	AsUser      bool         `json:"as_user"`
	Attachments []Attachment `json:"attachments"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	AuthorName string `json:"author_name"`
	Color      string `json:"color"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Footer string `json:"footer"`
}

type UserInviteResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

type UserSearchResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	User  User   `json:"user"`
}

type User struct {
	Id       string `json:"id"`
	TeamId   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	RealName string `json:"real_name"`
}
