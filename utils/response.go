package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

type Message struct {
	Success   bool        `json:"success"`
	ErrorCode string      `json:"error_code,omitempty"`
	Message   *string     `json:"message"`
	Code      string      `json:"code,omitempty"`
	Count     int         `json:"count,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	User      interface{} `json:"user,omitempty"`
	Token     string      `json:"token,omitempty"`
}

func SuccessResponse(res Message) Message {
	msg := findMessageWithCode("message", res.Code)

	return Message{
		Success: true,
		Message: &msg,
		Count:   res.Count,
		Data:    res.Data,
		User:    res.User,
		Token:   res.Token,
	}
}
func ErrorResponse(res Message) Message {
	msg := findMessageWithCode("errors", fmt.Sprint(res.ErrorCode))

	return Message{
		Success:   false,
		Message:   &msg,
		ErrorCode: res.ErrorCode,
	}
}

func findMessageWithCode(typeMsg string, code string) string {
	message_responses, _ := ioutil.ReadFile("./configs/message_responses.json")
	configValue := gjson.Get(string(message_responses), fmt.Sprintf("%s.%v", typeMsg, code))
	return configValue.String()
}
