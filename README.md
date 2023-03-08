# Go LINE Bot

這是一個使用 Go 語言串接 LINE Messaging API 與 MongoDB 的範例程式

把聊天機器人加入好友並開始對話後，對話內容便會存到 MongoDB 中，並提供完整 CRUD APIs 操作資料庫

![line-bot-demo](https://user-images.githubusercontent.com/58166555/223167931-160e3226-1fb9-47e7-8223-a575c14b172d.gif)

# 使用方式
1. 啟動 MongoDB 服務
```shell
docker run --name mongo -d -p 27017:27017 mongo:4.4
```

2. 在根目錄下新增 `/chat/app/setting.yml` ，並將 channel secret 與 channel access token 填入 （此程式會使用 Viper 讀取，secret 自 LINE Developer 取得）
```
channel_secret: "YOUR_CHANNEL_SECRET"
channel_access_token: "YOUR_CHANNEL_ACCESS_TOKEN"
```

3. 啟動程式
```shell
cd cmd/line_bot && go run main.go
```

4. 啟動 ngrok 反向代理，並將 URL 貼到 LINE Developer Console 中，註冊 Webhook

```shell
ngrok http 8080
```
5. 加入使用 secret 與 access token 的 channel 好友，開始傳送訊息

## 參考資料
- LINE Messaging API https://developers.line.biz/en/services/messaging-api/
