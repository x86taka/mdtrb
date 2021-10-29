package main

import (
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
		s.ChannelMessageSend(m.ChannelID, "> "+ipstr+"\n "+m.Author.Mention()+"\n```"+string(out)+"```")
	}
}

func Mtr(ipstr string, s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := exec.Command("mtr", "-n", "-z", "-r", "-c", "1", ipstr)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		s.ChannelMessageSend(m.ChannelID, "> "+ipstr+"\n "+m.Author.Mention()+"\n```"+string(out)+"```")
	}
}
