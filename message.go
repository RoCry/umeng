package umeng

import (
	"time"
)

const (
	UMENG_MESSAGE_TYPE_UNICAST        = "unicast"
	UMENG_MESSAGE_TYPE_LISTCAST       = "listcast"
	UMENG_MESSAGE_TYPE_FILECAST       = "filecast"
	UMENG_MESSAGE_TYPE_BROADCAST      = "broadcast"
	UMENG_MESSAGE_TYPE_GROUPCAST      = "groupcast"
	UMENG_MESSAGE_TYPE_CUSTOMIZEDCAST = "customizedcast"
)

////////////////////////////////////////////////////////////////////////////////
// UMengMessage
type UMengMessage struct {
	AppKey         string 			        `json:"appkey"`
	AppSecret      string                   `json:"-"`
	Timestamp      int64 	                `json:"timestamp"`
	Type           string                   `json:"type"`
	DeviceTokens   string                   `json:"device_tokens,omitempty"`
	AliasType      string                   `json:"alias_type,omitempty"`
	Alias          string                   `json:"alias,omitempty"`
	FileId         string                   `json:"file_id,omitempty"`
	ProductionMode string                   `json:"production_mode,omitempty"`
	Description    string                   `json:"description,omitempty"`
	ThirdPartyId   string                   `json:"thirdparty_id,omitempty"`
	Payload        map[string]interface{}   `json:"payload"`

	// default push
	MiPush     string   `json:"mipush,omitempty"`
	MiActivity string `json:"mi_activity,omitempty"`
}

//设定 当设备离线时转为系统下发
func (message *UMengMessage)SetMipush(defaultPush ,activity string) {
	message.MiPush = defaultPush
	message.MiActivity = activity
}

func NewUMengMessage(appKey, appSecret, mType string, production bool) *UMengMessage {
	var message = &UMengMessage{}
	message.AppKey    = appKey
	message.AppSecret = appSecret
	message.Type      = mType

	if production {
		message.ProductionMode = "true"
	} else {
		message.ProductionMode = "false"
	}

	message.Timestamp = time.Now().Unix()
	return message
}

////////////////////////////////////////////////////////////////////////////////
// UMengiOSPayload
type UMengiOSPayload map[string]interface{}

func NewUMengiOSPayload() UMengiOSPayload {
	var payload = UMengiOSPayload{}
	return payload
}

func (this UMengiOSPayload) updateAPSValue(key string, value interface{}) {
	if aps, ok := this["aps"].(map[string]interface{}); !ok {
		aps = make(map[string]interface{})
		aps[key] = value
		this["aps"] = aps
	} else {
		aps[key] = value
	}
}

func (this UMengiOSPayload) SetAlert(alert string) {
	this.updateAPSValue("alert", alert)
}

func (this UMengiOSPayload) SetBadge(badge int) {
	this.updateAPSValue("badge", badge)
}

func (this UMengiOSPayload) SetSound(sound string) {
	this.updateAPSValue("sound", sound)
}

func (this UMengiOSPayload) SetContentAvailable(contentAvailable string) {
	this.updateAPSValue("content-available", contentAvailable)
}

func (this UMengiOSPayload) SetCategory(category string) {
	this.updateAPSValue("category", category)
}

func (this UMengiOSPayload) Set(key, value string) {
	this[key] = value
}


// UMengAndroidpayload
type UMengAndroidPayload map[string]interface{}

const (
	UMENG_ANDROID_DISPLAY_TYPE_OF_NOTIFICATION = "notification"
	UMENG_ANDROID_DISPLAY_TYPE_OF_MESSAGE      = "message"
)

const  (
	UMENG_ANDROID_OPEN_TYPE_OF_APP      = "go_app"
	UMENG_ANDROID_OPEN_TYPE_OF_URL      = "go_url"
	UMENG_ANDROID_OPEN_TYPE_OF_ACTIVITY = "go_activity"
	UMENG_ANDROID_OPEN_TYPE_OF_CUSTOM   = "go_custom"
)

func NewUMengAndroidPayload() UMengAndroidPayload {
	var payload = UMengAndroidPayload{}
	payload.SetBuilderId(0)
	payload.SetDisplayType(UMENG_ANDROID_DISPLAY_TYPE_OF_NOTIFICATION)
	payload.SetAfterOpen(UMENG_ANDROID_OPEN_TYPE_OF_APP, "")
	return payload
}

func (this UMengAndroidPayload) SetDisplayType(displayType string) {
	this["display_type"] = displayType
}

func (this UMengAndroidPayload) updateBodyValue(key string, value interface{}) {
	if body, ok := this["body"].(map[string]interface{}); !ok {
		body = make(map[string]interface{})
		body[key] = value
		this["body"] = body
	} else {
		body[key] = value
	}
}

func (this UMengAndroidPayload) SetTicker(ticker string) {
	this.updateBodyValue("ticker", ticker)
}

func (this UMengAndroidPayload) SetTitle(title string) {
	this.updateBodyValue("title", title)
}

func (this UMengAndroidPayload) SetText(text string) {
	this.updateBodyValue("text", text)
}

func (this UMengAndroidPayload) SetIcon(icon string) {
	this.updateBodyValue("icon", icon)
}

func (this UMengAndroidPayload) SetLargeIcon(icon string) {
	this.updateBodyValue("largeIcon", icon)
}

func (this UMengAndroidPayload) SetImage(image string) {
	this.updateBodyValue("img", image)
}

func (this UMengAndroidPayload) SetSound(sound string) {
	this.updateBodyValue("sound", sound)
}

func (this UMengAndroidPayload) SetBuilderId(builderId int) {
	this.updateBodyValue("builder_id", builderId)
}

func (this UMengAndroidPayload) SetPlayVibrate(vibrate bool) {
	if vibrate {
		this.updateBodyValue("play_vibrate", "true")
	} else {
		this.updateBodyValue("play_vibrate", "false")
	}
}

func (this UMengAndroidPayload) SetPlayLights(lights bool) {
	if lights {
		this.updateBodyValue("play_lights", "true")
	} else {
		this.updateBodyValue("play_lights", "false")
	}
}

func (this UMengAndroidPayload) SetPlaySound(sound bool) {
	if sound {
		this.updateBodyValue("play_sound", "true")
	} else {
		this.updateBodyValue("play_sound", "false")
	}
}

func (this UMengAndroidPayload) SetCustom(custom string) {
	this.updateBodyValue("custom", custom)
}

func (this UMengAndroidPayload) SetAfterOpen(key string, value string) {
	this.updateBodyValue("after_open", key)
	if key == UMENG_ANDROID_OPEN_TYPE_OF_URL {
		this.updateBodyValue("url", value)
	} else if key == UMENG_ANDROID_OPEN_TYPE_OF_ACTIVITY {
		this.updateBodyValue("activity", value)
	}
}

func (this UMengAndroidPayload) SetExtra(key string, value interface{}) {
	if extra, ok := this["extra"].(map[string]interface{}); !ok {
		extra = make(map[string]interface{})
		extra[key] = value
		this["extra"] = extra
	} else {
		extra[key] = value
	}
}

