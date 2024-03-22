package logg

import (
	"log"
	"os"

	"dumpackets/config"
	"github.com/rs/zerolog"
)

var (
	Elogger zerolog.Logger
)

func init() {
	efile, err := os.OpenFile(config.Config.Log.ErrorLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("打开error日志文件失败")
	}
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000 +0800"
	Elogger = zerolog.New(efile).With().Timestamp().Logger().Level(1)
}
