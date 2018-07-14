package patrans

import "github.com/strongo/emoji/go/emoji"

func RegisterTranslations(appTranslations map[string]map[string]string) {
	for k, v := range TRANS {
		if _, ok := appTranslations[k]; ok {
			panic("duplicate")
		}
		appTranslations[k] = v
	}
}

var TRANS = map[string]map[string]string{
	"": {
		"en-US": "",
		"ru-RU": "",
		"fa-IR": "",
		"es-ES": "",
		"de-DE": "",
		"it-IT": "",
	},
	SinglePlayer: {
		"en-US": emoji.RobotFace + " Single-player (AI)",
		"ru-RU": emoji.RobotFace + " Играть одному (ИИ)",
		"fa-IR": emoji.RobotFace + "تک نفره (AI)",
		"es-ES": emoji.RobotFace + " Un solo jugador (AI)",
		"de-DE": emoji.RobotFace + " Einzelspieler (AI)",
		"it-IT": emoji.RobotFace + " Giocatore singolo (AI)",
		"fr-FR": emoji.RobotFace + " Un joueur (AI)",
	},
	// NewPlayWithAI: {
	// 	"en-US": emoji.RobotFace + " Play against AI",
	// 	"ru-RU": emoji.RobotFace + " Играть против компьютера",
	// },
	MultiPlayer: {
		"en-US": emoji.CrossedSwords + " Multi-player",
		"ru-RU": emoji.CrossedSwords + " Играть с противником",
		"fa-IR": emoji.CrossedSwords + " چند نفره",
		"es-ES": emoji.CrossedSwords + " Multi jugador",
		"de-DE": emoji.CrossedSwords + " Mehrspieler",
		"it-IT": emoji.CrossedSwords + " Multi-player",
		"fr-FR": emoji.CrossedSwords + " Multi-joueurs",
	},
	ChallengeFriendCommandText: {
		"en-US": "🤺 Challenge Telegram friend",
		"ru-RU": "🤺 Новая игра в Telegram",
		"fa-IR": "🤺 چالش Telegram دوست",
		"es-ES": "🤺 Desafía a amigo de Telegram",
		"de-DE": "🤺 Herausforderung Telegrammfreund",
		"it-IT": "🤺 Sfida l'amico di Telegram",
		"fr-FR": "🤺 Défi Télégramme ami",
	},
	MainMenuButton: {
		"en-US": emoji.BACKArrow + " Main menu",
		"ru-RU": emoji.BACKArrow + " Главное меню",
		"fa-IR": emoji.BACKArrow + " منوی اصلی",
		"es-ES": emoji.BACKArrow + " Menú principal",
		"de-DE": emoji.BACKArrow + " Hauptmenü",
		"it-IT": emoji.BACKArrow + " Menu principale",
		"fr-FR": emoji.BACKArrow + " Menu principal",
	},
	NewTournamentButton: {
		"en-US": "⚔ New tournament",
		"ru-RU": "⚔ Новый турнир",
		"fa-IR": "⚔ مسابقات جدید",
		"es-ES": "⚔ Nuevo torneo",
		"de-DE": "⚔ Neues Turnier",
		"it-IT": "⚔ Nuovo torneo",
		"fr-FR": "⚔ Nouveau tournoi",
	},
	TournamentsButton: {
		"en-US": "📰 Tournaments",
		"ru-RU": "📰 Турниры",
		"fa-IR": "📰 مسابقات",
		"es-ES": "📰 Torneos",
		"de-DE": "📰 Turniere",
		"it-IT": "📰 Tornei",
		"fr-FR": "📰 Tournois",
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

<code>Турнир</code> это серия матчей между несколькими игроками с подсчётом побед.

🏆 Победителем турнира становится тот кто победит наибольшее количество соперников.

💵 Время от времени проводятся спонсорские турниры с денежными призами. Узнавайте о таких турнирах подписавшись на канал @prizarena.

🆓 Участие в турнирах абсолютно <b>бесплатно</b>.

🎫 Правда некоторые турниры могут быть эксклюзивными для подписчиков канала или присоедениться можно только по приглашению.

📣 Создание спонсируемого турнира экслюзивного для ваших подписчиков отличный способ рекламы вашего Телеграм-канала.

🤞 Попробуйте выиграть и удачи!`,
		"fa-IR": `⚔ <b>مسابقات</b> از @prizarena_bot

<code> مسابقات </code> یک سری از مسابقات بین شمارش چند بازیکن است.

🏆 برنده مسابقات این است که برنده ترین رقبا است.

orship مسابقات حمایتی با جوایز نقدی از زمان به زمان برگزار می شود. با عضویت در کانالprizarena در مورد این تورنمنت ها اطلاعات کسب کنید.

🆓 مشارکت در مسابقات کاملا <b> رایگان </b> است.

🎫 درست است که بعضی از مسابقات می توانند برای مشترکان کانال منحصر به فرد باشند و یا فقط می توانند با دعوت نامه پیوست شوند.

📣 ایجاد یک تورنمنت حمایت شده برای مشترکین شما یک راه عالی برای تبلیغ کانال Telegram شما است.

🤞 سعی کنید برنده شوید و موفق باشید`,
		"it-IT": `⚔ <b>Tornei</b> di @prizarena_bot

<code>Tournament</code> è una serie di partite tra diversi giocatori che contano le vincite.

🏆 Il vincitore del torneo è colui che vince il maggior numero di rivali.

💵 I tornei di sponsorizzazione con premi in denaro sono tenuti di volta in volta. Scopri di più su tali tornei iscrivendoti al canale @prizarena.

🆓 La partecipazione ai tornei è assolutamente <b>gratuita</b>.

🎫 È vero, alcuni tornei possono essere esclusivi per gli abbonati del canale o possono essere uniti solo su invito.

📣 Creare un torneo sponsorizzato in esclusiva per i tuoi abbonati è un modo eccellente per pubblicizzare il tuo canale Telegram.

🤞 Prova a vincere e buona fortuna!`,
		"es-ES": `⚔ <b>Torneos</b> de @prizarena_bot

<code>Tournament</code> es una serie de partidos entre varios jugadores que cuentan victorias.

🏆 El ganador del torneo es el que gana la mayor cantidad de rivales.

Tour Torneos de patrocinio con premios en efectivo se llevan a cabo de vez en cuando. Obtenga información sobre dichos torneos suscribiéndose al canal @prizarena.

🆓 La participación en torneos es absolutamente <b>gratuita</b>.

🎫 Es cierto que algunos torneos pueden ser exclusivos para los suscriptores del canal o solo se pueden unir mediante invitación.

📣 Crear un torneo patrocinado exclusivo para sus suscriptores es una excelente manera de publicitar su canal Telegram.

🤞 ¡Intenta ganar y buena suerte!`,

		"de-DE": `⚔ <b>Turniere</b> von @prizarena_bot

<code> Tournament </code> ist eine Serie von Spielen zwischen mehreren Spielern, die Gewinne zählen.

🏆 Der Gewinner des Turniers ist derjenige, der die meisten Rivalen gewinnt.

💵 Von Zeit zu Zeit finden Sponsorturniere mit Geldpreisen statt. Erfahren Sie mehr über solche Turniere, indem Sie den @prizarena-Kanal abonnieren.

🆓 Die Teilnahme an Turnieren ist absolut <b>frei</b>.

🎫 Es stimmt, dass einige Turniere exklusiv für Abonnenten des Kanals sein können oder nur durch Einladung verbunden werden können.

📣 Das Erstellen eines gesponserten Turniers exklusiv für Ihre Abonnenten ist eine hervorragende Möglichkeit, Ihren Telegramm-Kanal zu bewerben.

🤞 Versuchen Sie zu gewinnen und viel Glück!`,
		"fr-FR": `⚔ <b>Tournois</b> par @prizarena_bot

<code>Tournoi</code> est une série de parties entre plusieurs joueurs qui comptent des points.

🏆 Le gagnant est celui qui bat plus d'adversaires que tout autre participant.

💵 De temps en temps, il y a des tournois sponsorisés avec des prix en argent. Recevez une notification en vous abonnant à la chaîne @prizarena.

🆓 C'est <b>gratuit </b> d'entrer et de jouer dans n'importe quel tournoi.

🎫 Certains tournois sont exclusifs et / ou sur invitation seulement.

📣 La création d'un tournoi sponsorisé réservé exclusivement à vos abonnés est un excellent moyen de promouvoir votre chaîne Telegram.

🤞 Essayez de gagner et bonne chance!`,
	},
}
