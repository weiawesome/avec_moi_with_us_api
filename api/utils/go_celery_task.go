package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const ContentEncoding = "utf-8"
const BodyEncoding = "base64"
const ContentType = "application/json"
const CeleryRedisKey = "celery"

type CeleryTask struct {
	Body            string      `json:"body"`
	ContentEncoding string      `json:"content-encoding"`
	ContentType     string      `json:"content-type"`
	Headers         interface{} `json:"headers"`
	Properties      interface{} `json:"properties"`
}

func (c *CeleryTask) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

type CeleryBody struct {
	Position []CeleryBodyPosition
	Args     interface{}
	Config   CeleryBodyConfig
}

func (c *CeleryBody) ToJsonBase64() (string, error) {
	data := []interface{}{c.Position, c.Args, c.Config}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return "", err
	}
	jsonString := string(jsonData)
	jsonData = []byte(jsonString)
	base64Data := base64.StdEncoding.EncodeToString(jsonData)
	return base64Data, nil
}

type CeleryBodyPosition struct{}

type CeleryOcrBodyArgs struct {
	Id             string `json:"id"`
	File           string `json:"file"`
	Prompt         string `json:"prompt"`
	ApiKey         string `json:"api_key"`
	AccessToken    string `json:"access_token"`
	KeyApiKey      string `json:"key_api_key"`
	KeyAccessToken string `json:"key_access_token"`
}
type CeleryOcrTextBodyArgs struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Prompt         string `json:"prompt"`
	ApiKey         string `json:"api_key"`
	AccessToken    string `json:"access_token"`
	KeyApiKey      string `json:"key_api_key"`
	KeyAccessToken string `json:"key_access_token"`
}
type CeleryAsrBodyArgs struct {
	Id             string `json:"id"`
	File           string `json:"file"`
	Prompt         string `json:"prompt"`
	ApiKey         string `json:"api_key"`
	AccessToken    string `json:"access_token"`
	KeyApiKey      string `json:"key_api_key"`
	KeyAccessToken string `json:"key_access_token"`
}
type CeleryNlpEditOcrBodyArgs struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Prompt         string `json:"prompt"`
	ApiKey         string `json:"api_key"`
	AccessToken    string `json:"access_token"`
	KeyApiKey      string `json:"key_api_key"`
	KeyAccessToken string `json:"key_access_token"`
}
type CeleryNlpEditAsrBodyArgs struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Prompt         string `json:"prompt"`
	ApiKey         string `json:"api_key"`
	AccessToken    string `json:"access_token"`
	KeyApiKey      string `json:"key_api_key"`
	KeyAccessToken string `json:"key_access_token"`
}
type CeleryMailBodyArgs struct {
	Mail string `json:"mail"`
}

type CeleryBodyConfig struct {
	Callbacks interface{} `json:"callbacks"`
	Chain     interface{} `json:"chain"`
	Chord     interface{} `json:"chord"`
}

type CeleryHeaders struct {
	Lang         string      `json:"lang"`
	Task         string      `json:"task"`
	Id           string      `json:"id"`
	Shadow       interface{} `json:"shadow"`
	Eta          interface{} `json:"eta"`
	Expires      interface{} `json:"expires"`
	Group        interface{} `json:"group"`
	GroupIndex   interface{} `json:"group_index"`
	Retries      int         `json:"retries"`
	RootId       string      `json:"root_id"`
	ParentId     interface{} `json:"parent_id"`
	Origin       string      `json:"origin"`
	IgnoreResult bool        `json:"ignore_result"`
}

type CeleryProperties struct {
	CorrelationId string                    `json:"correlation_id"`
	ReplyTo       string                    `json:"reply_to"`
	DeliveryMode  int                       `json:"delivery_mode"`
	DeliveryInfo  CeleryPropertyDeliverInfo `json:"delivery_info"`
	Priority      int                       `json:"priority"`
	BodyEncoding  string                    `json:"body_encoding"`
	DeliveryTag   string                    `json:"delivery_tag"`
}
type CeleryPropertyDeliverInfo struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
}
