package bot

import (
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var LineBot *linebot.Client

func init() {
	var err error
	LineBot, err = linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		fmt.Println(err)
		panic("line bot error")
	}
}
