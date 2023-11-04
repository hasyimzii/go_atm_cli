package logs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func SetLogger() {
	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed to create log file")
		panic(err)
	}

	// multiWriter := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(file)

	logrus.SetLevel(logrus.InfoLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
