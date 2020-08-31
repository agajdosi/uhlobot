package main

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/uhlobot/parlamentni"

	"github.com/chromedp/chromedp"
	"github.com/pelletier/go-toml"
)

func main() {
	config, err := toml.LoadFile("bots.toml")
	if err != nil {
		fmt.Println(err)
	}

	bots := config.Get("bot").([]*toml.Tree)
	for _, bot := range bots {
		fmt.Println(bot.Get("mail"))
		fmt.Println(bot.Get("parlamentni.password"))

		username := bot.Get("mail").(string)
		password := bot.Get("parlamentni.password").(string)

		url := "https://www.parlamentnilisty.cz/arena/monitor/Ovlivnovani-deti-vyber-pohlavi-ideologie-na-skolach-a-diktat-mensin-Detsky-psychiatr-vypustil-zasadni-varovani-Rodice-pozor-635311/diskuse"

		ctx, cancel := createBrowser()
		parlamentni.Login(ctx, cancel, url, username, password)
		//parlamentni.Comment(ctx, cancel, url, "Nesouhlas", "Tomu se mi nechce moc věřit, zní to dost vyumělkovaně.")

		fmt.Println("success!")
		time.Sleep(time.Second * 400)
	}

}

func createBrowser() (*context.Context, *context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		//chromedp.UserDataDir("config"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)

	return &ctx, &cancel
}
