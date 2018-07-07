package message

import (
	"strconv"
	"fmt"
)

type ReplyMode interface {
	fmt.Stringer
}

type Message struct {
	ChatID                string
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
	DisableNotification   bool
	ReplyToMessageID      int64
	ReplyMarkup           ReplyMode
}

func (m Message) GetMapValues() map[string]string {
	mapValues := map[string]string {
		"chat_id":    m.ChatID,
		"text":       m.Text,
	}
	if m.ParseMode != "" {
		mapValues["parse_mode"] = m.ParseMode
	}
	if m.DisableWebPagePreview {
		mapValues["disable_web_page_preview"] = strconv.FormatBool(m.DisableWebPagePreview)
	}
	if m.DisableNotification {
		mapValues["disable_notification"] = strconv.FormatBool(m.DisableNotification)
	}
	if m.ReplyToMessageID != 0 {
		mapValues["reply_to_message_id"] = strconv.FormatInt(m.ReplyToMessageID, 10)
	}
	if m.ReplyMarkup != nil {
		mapValues["reply_markup"] = m.ReplyMarkup.String()
	}
	return mapValues
}
