# Shadow Ban Bot
Marks messages from selected contacts as read silently.

Needs Telegram Premium.

## Run
```bash
go run main.go --token=[BOT_TOKEN]
```

## Connect bot
1. Go to Telegram Settings -> Telegram Business -> ChatBots
2. Add the bot
3. Configure the list of contacts you want to hide
4. Enable "Mark Messages As Read" bot permission

> ⚠️ The bot automatically marks messages as read for all selected contacts.
Make sure you only add users whose messages you really want to hide.