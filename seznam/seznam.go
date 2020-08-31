package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	ctx, cancel := createBrowser()

	register(ctx, cancel)

	fmt.Println("success!")
	time.Sleep(time.Second * 400)
}

func createBrowser() (*context.Context, *context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("user-agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0"),
		//chromedp.UserDataDir("config"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)

	return &ctx, &cancel
}

//Comment places a comment into the discussion located on URL.
func register(ctx *context.Context, cancel *context.CancelFunc) error {
	rand.Seed(time.Now().UnixNano())
	username := randUsername(6)
	password := randPassword(12)
	birth := strconv.Itoa(1970 + rand.Intn(30))

	fmt.Println(username, password, birth)

	err := chromedp.Run(*ctx,
		chromedp.Navigate("https://google.com/"),
		chromedp.Navigate("https://login.szn.cz/?service=homepage&return_url=https%3A%2F%2Fwww.seznam.cz%2F"),
		chromedp.WaitVisible("#login > div > form.login > footer > a:nth-child(2)", chromedp.ByQuery),
		chromedp.Click("#login > div > form.login > footer > a:nth-child(2)", chromedp.ByQuery),
		chromedp.WaitVisible(`#register-username`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.SendKeys(`#register-username`, username, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.SendKeys(`input[data-placeholder="register.password1"]`, password, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.SendKeys(`input[data-placeholder="register.password2"]`, password, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.SendKeys(`#register-year`, birth, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#register > div > form.main > fieldset > label:nth-child(1) > span`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`#register > div > form.main > label > span`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`button[data-locale="register.submit"]`, chromedp.ByQuery),
	)

	return err
}

func randPassword(n int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randUsername(n int) string {
	var letters = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	names := []string{
		"jiri",
		"jan",
		"honza",
		"petr",
		"daniel",
		"andrej",
		"adam",
		"filip",
		"pavel",
	}

	username := names[rand.Intn(len(names))] + string(b)

	return username
}
