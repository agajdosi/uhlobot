package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	ctx, cancel := createBrowser()

	seznamLogin(ctx, cancel)

	fmt.Println("success!")
	time.Sleep(time.Second * 400)
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

func seznamLogin(ctx *context.Context, cancel *context.CancelFunc) {

	err := chromedp.Run(*ctx,
		chromedp.Navigate(`https://www.novinky.cz/diskuze/40334782`),
		chromedp.Sleep(time.Second*1),
		chromedp.WaitVisible(`div[data-dot="diskuse_akce_nazory/button_pridat_nazor_horni"]`, chromedp.ByQuery),
		chromedp.Click(`div[data-dot="diskuse_akce_nazory/button_pridat_nazor_horni"]`, chromedp.ByQuery),

		chromedp.Sleep(time.Second*1),
		chromedp.WaitVisible(`#login-username`, chromedp.ByQuery),
		chromedp.SendKeys(`#login-username`, "jan.trnkaa", chromedp.ByQuery),
		chromedp.SendKeys(`#login-password`, "j37bj4E!", chromedp.ByQuery),
		chromedp.Sleep(time.Second*3),
		chromedp.Click(`form.login > button`, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

}

/* func iframes() {
	var iframes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(`#pollrecpatcha`, chromedp.ByQuery),
		chromedp.WaitVisible(`.modal-dialog`, chromedp.ByQuery),
		chromedp.Nodes(`#pollrecpatcha > div > div > iframe`, &iframes, chromedp.ByQueryAll),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = chromedp.Run(ctx,
		chromedp.Click(`.recaptcha-checkbox-border`, chromedp.ByQuery, chromedp.FromNode(iframes[0])),
		chromedp.Click(`#solver-button`, chromedp.ByQuery, chromedp.FromNode(iframes[0])),
		chromedp.Sleep(4*time.Second),
		chromedp.Click(`.btn-blue`, chromedp.ByQuery, chromedp.FromNode(iframes[0])),
		chromedp.Sleep(40*time.Second),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("voted!")
} */
