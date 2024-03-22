package dfile

import (
	"os"
	"time"

	"dumpackets/logg"
)

func OpenF() *os.File {
	t := time.Now()
	fileTime := t.Format("200601021504")
	fileName := "logs/dump-" + fileTime + ".pcap"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logg.Elogger.Err(err)
	}
	return file
}
