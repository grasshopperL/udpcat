package main

import (
	"dumpackets/config"
	"dumpackets/logg"
	"dumpackets/ppackets"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ws := &config.WriteConfig{Writing: false}
	writing := make(chan struct{}, 1)
	writing <- struct{}{}

	// read data form live device
	handle, err := pcap.OpenLive(config.Config.Server.DeviceName, config.Config.Server.SnapLen,
		config.Config.Server.Promisc, config.Config.Server.Timeout)
	if err != nil {
		logg.Elogger.Err(err).Msg("监听流量失败")
	}
	defer handle.Close()
	err = handle.SetBPFFilter(config.Config.Server.BpfFilter)
	if err != nil {
		logg.Elogger.Fatal()
	}

	packets := gopacket.NewPacketSource(handle, handle.LinkType())

	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	// if writing file, use channel to block ticker
	for range ticker.C {
		<-writing
		ppackets.PacketCatch(wg, packets, ws, writing)
		fmt.Println("aa")

	}
	wg.Wait()
}
