package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func Traceroute(ipstr string, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := exec.Command("traceroute", "-A", ipstr)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		content := fmt.Sprintf("> %s\n %s\n```%s```", ipstr, m.Author.Mention(), string(out))
		s.ChannelMessageSend(m.ChannelID, content)
	}
}

func Mtr(ipstr string, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := exec.Command("mtr", "-n", "-z", "-r", "-c", "1", ipstr)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		content := fmt.Sprintf("> %s\n %s\n```%s```", ipstr, m.Author.Mention(), string(out))
		s.ChannelMessageSend(m.ChannelID, content)
	}
}
