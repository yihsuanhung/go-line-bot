package settings

type LineBotInfo struct {
	ChannelSecret      string `mapstructure:"channel_secret"`
	ChannelAccessToken string `mapstructure:"channel_access_token"`
}
