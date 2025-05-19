package config

import "os"

const (
	Gemini_Welcome_Reply_Key = "geminiWelcomeReply"
	Gemini_Key               = "geminiKey"
)

func GetGeminiWelcomeReply() (r string) {
	r = os.Getenv(Gemini_Welcome_Reply_Key)
	if r == "" {
		r = "我是 Gemini，开始聊天吧！"
	}
	return
}

func GetGeminiKey() string {
	return os.Getenv(Gemini_Key)
}
