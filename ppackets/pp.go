package ppackets

import (
	"dumpackets/config"
	"dumpackets/dfile"
	"dumpackets/logg"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"os"
	"sync"
)

// PacketCatch deal with different scene, including read packet err, no packet and begin to receive packet
func PacketCatch(wg *sync.WaitGroup, packets *gopacket.PacketSource, ws *config.WriteConfig, writing chan struct{}) {
	p, err := packets.NextPacket()

	// read packet err
	if err != nil {
		logg.Elogger.Err(err)
	}

	// no packets receive
	if p == nil {
		writing <- struct{}{}
		return
	}

	// begin to receive packet and write packet
	go WritePackets(packets.Packets(), ws, writing)
}

// WritePackets write packets form channel to file
func WritePackets(packets chan gopacket.Packet, ws *config.WriteConfig, writing chan struct{}) {
	file := dfile.OpenF()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logg.Elogger.Err(err)
		}
	}(file)

	pw := pcapgo.NewWriter(file)
	err := pw.WriteFileHeader(1024, layers.LinkTypeEthernet)
	if err != nil {
		logg.Elogger.Err(err)
	}

	for p := range packets {
		udpLayer := p.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			err := pw.WritePacket(p.Metadata().CaptureInfo, p.Data())
			if err != nil {
				logg.Elogger.Err(err)
			}
		}
	}

	// finish writing file
	writing <- struct{}{}
}
