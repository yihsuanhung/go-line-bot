package bot

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/yihsuanhung/go-line-bot/pkg/settings"
)

var LineBot *linebot.Client

func InitLineBot(s *settings.LineBotInfo) {
	var err error
	LineBot, err = linebot.New(
		s.ChannelSecret,
		s.ChannelAccessToken,
	)
	if err != nil {
		fmt.Println(err)
		panic("line bot error")
	}
}
