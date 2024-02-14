package model

type NotifyBaseRequest struct {
	UserId     int64  `json:"user_id,omitempty"`
	CategoryId int64  `json:"cat_id,omitempty"`
	FcmToken   string `json:"fcm_token" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Message    string `json:"message" validate:"required"`
}

type ComposePushMessage struct {
	ApplicationId string            `json:"application_id,omitempty"`
	Action        string            `json:"action,omitempty"`
	Image         string            `json:"image,omitempty"`
	Icon          string            `json:"icon,omitempty"`
	Data          map[string]string `json:"data,omitempty"`
	Priority      string            `json:"priority,omitempty"`
	Destination   string            `json:"destination,omitempty"`
	Title         string            `json:"title,omitempty"`
	Message       string            `json:"message,omitempty"`
	CategoryId    int64             `json:"cat_id,omitempty"`
	UserId        int64             `json:"user_id,omitempty"`
	Url           string            `json:"url,omitempty"`
	CustomData    map[string]string `json:"custom_data,omitempty"`
	AppId         string            `json:"app_id,omitempty"`
	TemplateId    string            `json:"template_id,omitempty"`
}
