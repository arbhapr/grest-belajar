package app

import "grest.dev/grest"

// telegram returns a pointer to the telegramUtil instance (telegram).
// If telegram is not initialized, it creates a new telegramUtil instance, configures it, and assigns it to telegram.
// It ensures that only one instance of telegramUtil is created and reused.
func Telegram(message ...string) *telegramUtil {
	if telegram == nil {
		telegram = &telegramUtil{}
		telegram.configure()
		if len(message) > 0 {
			telegram.AddMessage(message[0])
		}
	}
	return telegram
}

// telegram is a pointer to a telegramUtil instance.
// It is used to store and access the singleton instance of telegramUtil.
var telegram *telegramUtil

// telegramUtil represents a utility to interact with telegram API.
// It embeds grest.Telegram, indicating that telegramUtil inherits from grest.Telegram.
type telegramUtil struct {
	grest.Telegram
}

// configure configures the telegram utility instance.
func (t *telegramUtil) configure() {
	t.BotToken = TELEGRAM_ALERT_TOKEN
	t.ChatID = TELEGRAM_ALERT_USER_ID
}
