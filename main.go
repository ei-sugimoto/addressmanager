package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// ローカル環境でのみ.envを読み込む
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// 環境変数DISCORD_TOKENが設定されていない場合は終了
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatalf("DISCORD_TOKEN is not set")
	}

	// Discordクライアントを作成.
	// "Bot "をトークンの前に付けることでBotとしてログインする
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating Discord client: %s", err)
	}

	// インテントを設定
	client.Identify.Intents = discordgo.IntentsGuildMessages

	// インタラクションのハンドラを追加
	client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.ApplicationCommandData().Name {
		case "sit":
			Sit(s, i)
		case "leave":
			Leave(s, i)
		case "now":
			Now(s, i)
		}
	})

	// Readyイベントのハンドラを追加
	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		commands := []*discordgo.ApplicationCommand{
			{
				Name:        "sit",
				Description: "席を確保します",
			},
			{
				Name:        "leave",
				Description: "席を解放します",
			},
			{
				Name:        "now",
				Description: "現在の席数を表示します",
			},
		}

		// コマンドを登録
		for _, cmd := range commands {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
			if err != nil {
				fmt.Println("Error creating command:", err)
			}
		}
	})

	// Discordに接続
	err = client.Open()
	if err != nil {
		log.Fatalf("Error opening Discord connection: %s", err)
	}

	// 終了シグナルを受信するまで待機
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// Discordとの接続を閉じる
	err = client.Close()
	if err != nil {
		log.Fatalf("Error closing Discord client: %s", err)
	}

	fmt.Println("Discord client closed")

}
