// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	io "io"

	models "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	mock "github.com/stretchr/testify/mock"

	msteams "github.com/mattermost/mattermost-plugin-msteams-sync/server/msteams"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *Client) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrGetChatForUsers provides a mock function with given fields: usersIDs
func (_m *Client) CreateOrGetChatForUsers(usersIDs []string) (string, error) {
	ret := _m.Called(usersIDs)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(usersIDs)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(usersIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteChatMessage provides a mock function with given fields: chatID, msgID
func (_m *Client) DeleteChatMessage(chatID string, msgID string) error {
	ret := _m.Called(chatID, msgID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(chatID, msgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessage provides a mock function with given fields: teamID, channelID, parentID, msgID
func (_m *Client) DeleteMessage(teamID string, channelID string, parentID string, msgID string) error {
	ret := _m.Called(teamID, channelID, parentID, msgID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(teamID, channelID, parentID, msgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSubscription provides a mock function with given fields: subscriptionID
func (_m *Client) DeleteSubscription(subscriptionID string) error {
	ret := _m.Called(subscriptionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(subscriptionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetChannelInTeam provides a mock function with given fields: teamID, channelID
func (_m *Client) GetChannelInTeam(teamID string, channelID string) (*msteams.Channel, error) {
	ret := _m.Called(teamID, channelID)

	var r0 *msteams.Channel
	if rf, ok := ret.Get(0).(func(string, string) *msteams.Channel); ok {
		r0 = rf(teamID, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(teamID, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelsInTeam provides a mock function with given fields: teamID, filterQuery
func (_m *Client) GetChannelsInTeam(teamID string, filterQuery string) ([]*msteams.Channel, error) {
	ret := _m.Called(teamID, filterQuery)

	var r0 []*msteams.Channel
	if rf, ok := ret.Get(0).(func(string, string) []*msteams.Channel); ok {
		r0 = rf(teamID, filterQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msteams.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(teamID, filterQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChat provides a mock function with given fields: chatID
func (_m *Client) GetChat(chatID string) (*msteams.Chat, error) {
	ret := _m.Called(chatID)

	var r0 *msteams.Chat
	if rf, ok := ret.Get(0).(func(string) *msteams.Chat); ok {
		r0 = rf(chatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Chat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChatMessage provides a mock function with given fields: chatID, messageID
func (_m *Client) GetChatMessage(chatID string, messageID string) (*msteams.Message, error) {
	ret := _m.Called(chatID, messageID)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string) *msteams.Message); ok {
		r0 = rf(chatID, messageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(chatID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCodeSnippet provides a mock function with given fields: url
func (_m *Client) GetCodeSnippet(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFileContent provides a mock function with given fields: weburl
func (_m *Client) GetFileContent(weburl string) ([]byte, error) {
	ret := _m.Called(weburl)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(weburl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(weburl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMe provides a mock function with given fields:
func (_m *Client) GetMe() (*msteams.User, error) {
	ret := _m.Called()

	var r0 *msteams.User
	if rf, ok := ret.Get(0).(func() *msteams.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: teamID, channelID, messageID
func (_m *Client) GetMessage(teamID string, channelID string, messageID string) (*msteams.Message, error) {
	ret := _m.Called(teamID, channelID, messageID)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string, string) *msteams.Message); ok {
		r0 = rf(teamID, channelID, messageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(teamID, channelID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyID provides a mock function with given fields:
func (_m *Client) GetMyID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReply provides a mock function with given fields: teamID, channelID, messageID, replyID
func (_m *Client) GetReply(teamID string, channelID string, messageID string, replyID string) (*msteams.Message, error) {
	ret := _m.Called(teamID, channelID, messageID, replyID)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string, string, string) *msteams.Message); ok {
		r0 = rf(teamID, channelID, messageID, replyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(teamID, channelID, messageID, replyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeam provides a mock function with given fields: teamID
func (_m *Client) GetTeam(teamID string) (*msteams.Team, error) {
	ret := _m.Called(teamID)

	var r0 *msteams.Team
	if rf, ok := ret.Get(0).(func(string) *msteams.Team); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeams provides a mock function with given fields: filterQuery
func (_m *Client) GetTeams(filterQuery string) ([]*msteams.Team, error) {
	ret := _m.Called(filterQuery)

	var r0 []*msteams.Team
	if rf, ok := ret.Get(0).(func(string) []*msteams.Team); ok {
		r0 = rf(filterQuery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msteams.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filterQuery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userID
func (_m *Client) GetUser(userID string) (*msteams.User, error) {
	ret := _m.Called(userID)

	var r0 *msteams.User
	if rf, ok := ret.Get(0).(func(string) *msteams.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserAvatar provides a mock function with given fields: userID
func (_m *Client) GetUserAvatar(userID string) ([]byte, error) {
	ret := _m.Called(userID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListChannels provides a mock function with given fields: teamID
func (_m *Client) ListChannels(teamID string) ([]msteams.Channel, error) {
	ret := _m.Called(teamID)

	var r0 []msteams.Channel
	if rf, ok := ret.Get(0).(func(string) []msteams.Channel); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]msteams.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTeams provides a mock function with given fields:
func (_m *Client) ListTeams() ([]msteams.Team, error) {
	ret := _m.Called()

	var r0 []msteams.Team
	if rf, ok := ret.Get(0).(func() []msteams.Team); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]msteams.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields:
func (_m *Client) ListUsers() ([]msteams.User, error) {
	ret := _m.Called()

	var r0 []msteams.User
	if rf, ok := ret.Get(0).(func() []msteams.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]msteams.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSubscription provides a mock function with given fields: subscriptionID
func (_m *Client) RefreshSubscription(subscriptionID string) (*time.Time, error) {
	ret := _m.Called(subscriptionID)

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func(string) *time.Time); ok {
		r0 = rf(subscriptionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(subscriptionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendChat provides a mock function with given fields: chatID, parentID, message, mentions
func (_m *Client) SendChat(chatID string, parentID string, message string, mentions []models.ChatMessageMentionable) (*msteams.Message, error) {
	ret := _m.Called(chatID, parentID, message, mentions)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string, string, []models.ChatMessageMentionable) *msteams.Message); ok {
		r0 = rf(chatID, parentID, message, mentions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, []models.ChatMessageMentionable) error); ok {
		r1 = rf(chatID, parentID, message, mentions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: teamID, channelID, parentID, message
func (_m *Client) SendMessage(teamID string, channelID string, parentID string, message string) (*msteams.Message, error) {
	ret := _m.Called(teamID, channelID, parentID, message)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string, string, string) *msteams.Message); ok {
		r0 = rf(teamID, channelID, parentID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(teamID, channelID, parentID, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessageWithAttachments provides a mock function with given fields: teamID, channelID, parentID, message, attachments, mentions
func (_m *Client) SendMessageWithAttachments(teamID string, channelID string, parentID string, message string, attachments []*msteams.Attachment, mentions []models.ChatMessageMentionable) (*msteams.Message, error) {
	ret := _m.Called(teamID, channelID, parentID, message, attachments, mentions)

	var r0 *msteams.Message
	if rf, ok := ret.Get(0).(func(string, string, string, string, []*msteams.Attachment, []models.ChatMessageMentionable) *msteams.Message); ok {
		r0 = rf(teamID, channelID, parentID, message, attachments, mentions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, []*msteams.Attachment, []models.ChatMessageMentionable) error); ok {
		r1 = rf(teamID, channelID, parentID, message, attachments, mentions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetChatReaction provides a mock function with given fields: chatID, messageID, userID, emoji
func (_m *Client) SetChatReaction(chatID string, messageID string, userID string, emoji string) error {
	ret := _m.Called(chatID, messageID, userID, emoji)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(chatID, messageID, userID, emoji)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReaction provides a mock function with given fields: teamID, channelID, parentID, messageID, userID, emoji
func (_m *Client) SetReaction(teamID string, channelID string, parentID string, messageID string, userID string, emoji string) error {
	ret := _m.Called(teamID, channelID, parentID, messageID, userID, emoji)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) error); ok {
		r0 = rf(teamID, channelID, parentID, messageID, userID, emoji)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeToChannel provides a mock function with given fields: teamID, channelID, baseURL, webhookSecret
func (_m *Client) SubscribeToChannel(teamID string, channelID string, baseURL string, webhookSecret string) (*msteams.Subscription, error) {
	ret := _m.Called(teamID, channelID, baseURL, webhookSecret)

	var r0 *msteams.Subscription
	if rf, ok := ret.Get(0).(func(string, string, string, string) *msteams.Subscription); ok {
		r0 = rf(teamID, channelID, baseURL, webhookSecret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(teamID, channelID, baseURL, webhookSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToChannels provides a mock function with given fields: baseURL, webhookSecret, pay
func (_m *Client) SubscribeToChannels(baseURL string, webhookSecret string, pay bool) (*msteams.Subscription, error) {
	ret := _m.Called(baseURL, webhookSecret, pay)

	var r0 *msteams.Subscription
	if rf, ok := ret.Get(0).(func(string, string, bool) *msteams.Subscription); ok {
		r0 = rf(baseURL, webhookSecret, pay)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool) error); ok {
		r1 = rf(baseURL, webhookSecret, pay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToChats provides a mock function with given fields: baseURL, webhookSecret, pay
func (_m *Client) SubscribeToChats(baseURL string, webhookSecret string, pay bool) (*msteams.Subscription, error) {
	ret := _m.Called(baseURL, webhookSecret, pay)

	var r0 *msteams.Subscription
	if rf, ok := ret.Get(0).(func(string, string, bool) *msteams.Subscription); ok {
		r0 = rf(baseURL, webhookSecret, pay)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool) error); ok {
		r1 = rf(baseURL, webhookSecret, pay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToUserChats provides a mock function with given fields: user, baseURL, webhookSecret, pay
func (_m *Client) SubscribeToUserChats(user string, baseURL string, webhookSecret string, pay bool) (*msteams.Subscription, error) {
	ret := _m.Called(user, baseURL, webhookSecret, pay)

	var r0 *msteams.Subscription
	if rf, ok := ret.Get(0).(func(string, string, string, bool) *msteams.Subscription); ok {
		r0 = rf(user, baseURL, webhookSecret, pay)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, bool) error); ok {
		r1 = rf(user, baseURL, webhookSecret, pay)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnsetChatReaction provides a mock function with given fields: chatID, messageID, userID, emoji
func (_m *Client) UnsetChatReaction(chatID string, messageID string, userID string, emoji string) error {
	ret := _m.Called(chatID, messageID, userID, emoji)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(chatID, messageID, userID, emoji)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsetReaction provides a mock function with given fields: teamID, channelID, parentID, messageID, userID, emoji
func (_m *Client) UnsetReaction(teamID string, channelID string, parentID string, messageID string, userID string, emoji string) error {
	ret := _m.Called(teamID, channelID, parentID, messageID, userID, emoji)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) error); ok {
		r0 = rf(teamID, channelID, parentID, messageID, userID, emoji)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateChatMessage provides a mock function with given fields: chatID, msgID, message, mentions
func (_m *Client) UpdateChatMessage(chatID string, msgID string, message string, mentions []models.ChatMessageMentionable) error {
	ret := _m.Called(chatID, msgID, message, mentions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []models.ChatMessageMentionable) error); ok {
		r0 = rf(chatID, msgID, message, mentions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMessage provides a mock function with given fields: teamID, channelID, parentID, msgID, message, mentions
func (_m *Client) UpdateMessage(teamID string, channelID string, parentID string, msgID string, message string, mentions []models.ChatMessageMentionable) error {
	ret := _m.Called(teamID, channelID, parentID, msgID, message, mentions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, []models.ChatMessageMentionable) error); ok {
		r0 = rf(teamID, channelID, parentID, msgID, message, mentions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadFile provides a mock function with given fields: teamID, channelID, filename, filesize, mimeType, data
func (_m *Client) UploadFile(teamID string, channelID string, filename string, filesize int, mimeType string, data io.Reader) (*msteams.Attachment, error) {
	ret := _m.Called(teamID, channelID, filename, filesize, mimeType, data)

	var r0 *msteams.Attachment
	if rf, ok := ret.Get(0).(func(string, string, string, int, string, io.Reader) *msteams.Attachment); ok {
		r0 = rf(teamID, channelID, filename, filesize, mimeType, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msteams.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, string, io.Reader) error); ok {
		r1 = rf(teamID, channelID, filename, filesize, mimeType, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
