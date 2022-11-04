package constants

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	ChannelId, _         = strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	ChannelName          = os.Getenv("CHANNEL_NAME")
	Token                = os.Getenv("TOKEN")
	NumImagesResponse, _ = strconv.Atoi(os.Getenv("NUM_IMAGES"))

	DatabaseURL = os.Getenv("DB_URL")

	data, _             = ioutil.ReadFile("data/images/welcome.jpeg")
	WelcomeImage        = tgbotapi.FileBytes{Name: "data/images/welcome.jpeg", Bytes: data}
	WelcomeMassage      = "Таро — это инструмент для глубокого познания себя и мира вокруг, но важно понимать, что не все так просто. Есть много нюансов от неправильно поставленного вопроса до просто неподходящего времени.\nЭтот бот написан в развлекательных целях, если же есть желание сделать правильный расклад, лучше обратиться к профессионалу(Смотри <a href=\"https://t.me/" + ChannelName + "\">канал</a>).\nЧто хочешь спросить?🙃"
	RegistrationMassage = "Подпишись на <a href=\"https://t.me/" + ChannelName + "\">" + ChannelName + "</a>, чтобы пользоваться ботом😉"
)
