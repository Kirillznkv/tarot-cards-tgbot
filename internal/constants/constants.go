package constants

import (
	"os"
	"strconv"
)

var (
	Token                = os.Getenv("TOKEN")
	NumImagesResponse, _ = strconv.Atoi(os.Getenv("NUM_IMAGES"))
	WelcomeMassage       = "Hello, new friend!"
)
