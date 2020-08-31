package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	username := "jan.trnkaa@seznam.cz"
	password := "ah63SVTxNVGSwmP"
	url := "https://www.idnes.cz/zpravy/domaci/psychiatricky-ustav-reforma-psychiatricke-pece-ombudsman.A200817_155135_domaci_knn/diskuse"

	ctx, cancel := createBrowser()
	idnesLogin(ctx, cancel, url, username, password)

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

func idnesLogin(ctx *context.Context, cancel *context.CancelFunc, url, username, password string) {

	err := chromedp.Run(*ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*1),
		chromedp.WaitVisible(`a[class="webz-bg"]`, chromedp.ByQuery),
		chromedp.Click(`a[class="webz-bg"]`, chromedp.ByQuery),

		chromedp.Sleep(time.Second*1),
		chromedp.WaitVisible(`#prem_login`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="email"]`, username, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="password"]`, password+"\n", chromedp.ByQuery),
		chromedp.Sleep(time.Second*1),
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
