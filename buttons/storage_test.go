package buttons

import (
	"testing"

	"go.uber.org/mock/gomock"
	tele "gopkg.in/telebot.v3"

	"github.com/stretchr/testify/assert"

	"telegram-bot/buttons/mocks"
)

type User struct {
	ID int64 `json:"id"`

	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	IsForum           bool     `json:"is_forum"`
	Username          string   `json:"username"`
	LanguageCode      string   `json:"language_code"`
	IsBot             bool     `json:"is_bot"`
	IsPremium         bool     `json:"is_premium"`
	AddedToMenu       bool     `json:"added_to_attachment_menu"`
	Usernames         []string `json:"active_usernames"`
	CustomEmojiStatus string   `json:"emoji_status_custom_emoji_id"`

	// Returns only in getMe
	CanJoinGroups   bool `json:"can_join_groups"`
	CanReadMessages bool `json:"can_read_all_group_messages"`
	SupportsInline  bool `json:"supports_inline_queries"`
}

func TestStorage(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)

	r := 1.9283

	mockStorage.EXPECT().AddRate(gomock.Any(), "Вася", r).Return(nil).Times(1)

	mockReate := mocks.NewMockRate(ctrl)

	mockReate.EXPECT().Get(gomock.Any()).Return(r, nil).Times(1)

	telemock := mocks.NewMockTeleContex(ctrl)
	telemock.EXPECT().Send(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
	telemock.EXPECT().Sender().Return(&tele.User{
		ID:                0,
		FirstName:         "",
		LastName:          "",
		IsForum:           false,
		Username:          "Вася",
		LanguageCode:      "",
		IsBot:             false,
		IsPremium:         false,
		AddedToMenu:       false,
		Usernames:         nil,
		CustomEmojiStatus: "",
		CanJoinGroups:     false,
		CanReadMessages:   false,
		SupportsInline:    false,
	}).Times(2)

	h := handlers{
		storage: mockStorage,
		apiRate: mockReate,
	}

	err := h.startHandler(telemock)

	assert.NoError(t, err)
}
