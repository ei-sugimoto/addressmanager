package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Sit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 席数を減らす
	err := DegreaseSeatsCount()
	if err != nil {
		// 席数が0の場合はエラーメッセージを返す
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "確保できる席はありません。そのコマンドは実行できません。",
			},
		})
		if err != nil {
			log.Println("Error responding to /sit:", err)
		}
	}
	// ユーザー名を取得。ニックネームが設定されている場合はそれを使う
	userName := i.Member.Nick
	if userName == "" {
		userName = i.Member.User.Username
	}
	content := fmt.Sprintf("@here %s さんが席を確保しました。残りの席数は %d です。", userName, SeatsCount)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		log.Println("Error responding to /sit:", err)
	}
}

func Leave(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := IncreaseSeatsCount()
	if err != nil {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "最大席数を越えます。そのコマンドは実行できません。",
			},
		})
		if err != nil {
			log.Println("Error responding to /leave:", err)
		}
	}
	userName := i.Member.Nick

	content := fmt.Sprintf("@here %sさんが席を解放しました。残りの席数は %d です。", userName, SeatsCount)
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		log.Println("Error responding to /leave:", err)
	}
}

func Now(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := fmt.Sprintf("現在の席数は %d です。", SeatsCount)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		log.Println("Error responding to /now:", err)
	}
}
