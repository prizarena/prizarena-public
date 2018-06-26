package patrans

func RegisterTranslations(appTranslations map[string]map[string]string) {
	for k, v := range TRANS {
		if _, ok := appTranslations[k]; ok {
			panic("duplicate")
		}
		appTranslations[k] = v
	}
}

var TRANS = map[string]map[string]string{
	NewTournamentButton: {
		"en-US": "⚔ New tournament",
		"ru-RU": "⚔ Новый турнир",
	},
	TournamentsButton: {
		"en-US": "📰 Tournaments",
		"ru-RU": "📰 Турниры",
	},
	TournamentsIntro: {
		"en-US": `⚔ <b>Tournaments</b> by @prizarena_bot

<code>Tournament</code> is a series of plays between multiple players with score counting.

🏆 The winner is the one who defeats more opponents then any other participant.

💵 From time to time there are sponsored tournaments with cash prizes. Get notified by subscribing to @prizarena channel.

🆓 It's <b>free</b> to enter and to play in any tournament.

🎫 Though some tournaments are exclusive and/or by invite only.

📣 Creating a sponsored tournament exclusive to your subscribers only is a great way to promote your Telegram-channel.

🤞 Try to win and good luck!`,
		"ru-RU": `⚔ <b>Турниры</b> от @prizarena_bot

<code>Турнир</code> это серия матчей между нисколькими игроками с подсчётом побед.

🏆 Победителем турнира становится тот кто победит наибольшее количество соперников.

💵 Время от времени проводятся спосорские турниры с денежными призами. Узнавайте о таких турнирах подписавшись на канал @prizarena.

🆓 Участие в турнирах абсолютно <b>бесплатно</b>.

🎫 Правда некоторые турниры могут быть эксклюзивными для подписчиков канала или присоедениться можно только по приглашению.

📣 Создание спосируемого турнира экслюзивного для ваших подписчиков отличный способ рекламы вашего Телеграм-канала.

🤞 Попробуйте выиграть и удачи!`,
	},
}
