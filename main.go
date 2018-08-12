package main

import (
	"log"
	"time"

	"github.com/3d0c/gmf"
	"github.com/imkira/go-libav/avformat"
)

func init() {
	avformat.RegisterAll()
}

func asyncCopyPackets() {
	ch := make(chan *gmf.Packet)
	files := []*VFile{
		&VFile{Name: "./1mp3.mp4"},
		&VFile{Name: "./2mp3.mp4"},
	}
	reader, _ := CreateReader(ch, files)
	// rtmp := "rtmp://live-prg.twitch.tv/live_129862765_H7988wWNq4m2kNaPPwnHkIxRKIsoDB"
	writer, _ := CreateWriter(ch, "./t.flv")
	// read
	log.Println("INFO: Reader Start Loop")
	go reader.StartLoop()
	// write
	log.Println("INFO: Writer Prepare")
	writer.Prepare()
	log.Println("INFO: Writer Start Loop")
	go writer.StartLoop()
	time.Sleep(time.Minute * 4)
}

func main() {
	asyncCopyPackets()
}
