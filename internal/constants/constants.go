package constants

import (
	"os"
	"strconv"
)

var (
	ChannelId, _         = strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	ChannelName          = os.Getenv("CHANNEL_NAME")
	Token                = os.Getenv("TOKEN")
	NumImagesResponse, _ = strconv.Atoi(os.Getenv("NUM_IMAGES"))

	DatabaseURL = os.Getenv("DB_URL")

	WelcomeMassage      = "Hello, new friend!"
	RegistrationMassage = "Подпишись на <a href=\"https://t.me/" + ChannelName + "\">" + ChannelName + "</a>, чтобы пользоваться ботом😉"
)
