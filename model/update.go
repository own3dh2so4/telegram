package model

type Response struct {
	Ok bool `json:"ok"`
}

type UpdateResponse struct {
	Response
	Result []struct {
		UpdateID    int `json:"update_id"`
		ChannelPost struct {
			MessageID int `json:"message_id"`
			Chat      struct {
				ID    int64  `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
			} `json:"chat"`
			Date            int `json:"date"`
			ForwardFromChat struct {
				ID       int64  `json:"id"`
				Title    string `json:"title"`
				Username string `json:"username"`
				Type     string `json:"channel"`
			}
			ForwardFromMessageID int64  `json:"forward_from_message_id"`
			ForwardDate          int64  `json:"forward_date"`
			Text                 string `json:"text"`
			Entities             []struct {
				Offset int64  `json:"offset"`
				Length int64  `json:"length"`
				Type   string `json:"type"`
				Url    string `json:"url"`
			} `json:"entities"`
		} `json:"channel_post"`
	} `json:"result"`
}

type SendMessageResponse struct {
	Response
	Result struct {
		MessageID int `json:"message_id"`
		Chat      struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}
