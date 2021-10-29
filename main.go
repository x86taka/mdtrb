package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/x86taka/mdtrb/bgpview"
)

type TraceT struct {
	ip string
	s  *discordgo.Session
	m  *discordgo.MessageCreate
}

var (
	Token string
	r     = &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", "1.1.1.1:53")
		},
	}
	taskchan = make(chan TraceT, 2000)
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	for i := 0; i < 2; i++ {
		go TaskTrace()
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Type != discordgo.MessageTypeDefault {
		return
	}

	if m.Content == "" {
		return
	}
	//Multiline
	strs := strings.Split(m.Content, "\n")
	for _, str := range strs {
		if len(str) != 0 {
			runCmd(str, s, m)
		}
	}
}

func runCmd(content string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// ASoo -> Response ASName
	if strings.HasPrefix(content, "AS") {
		ASinfo := bgpview.GetASNName(strings.ReplaceAll(content, "AS", ""))
		s.ChannelMessageSend(m.ChannelID, ASinfo)
	}
	//Check Hostname Or Ip
	if _, err := net.ResolveIPAddr("ip", content); err == nil || net.ParseIP(content) != nil {
		tt := TraceT{
			content,
			s,
			m,
		}
		taskchan <- tt
	}
}

func TaskTrace() {
	for {
		task := <-taskchan
		Mtr(task.ip, task.s, task.m)
	}
}
