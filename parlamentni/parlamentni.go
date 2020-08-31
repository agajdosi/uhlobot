package parlamentni

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

//Login logins a user into the web.
func Login(ctx *context.Context, cancel *context.CancelFunc, url, username, password string) error {
	err := chromedp.Run(*ctx,
		chromedp.Navigate("https://www.parlamentnilisty.cz/login.aspx"),
		chromedp.Sleep(time.Second*1),

		chromedp.WaitVisible(`#inputEmail`, chromedp.ByQuery),
		chromedp.SendKeys(`#inputEmail`, username, chromedp.ByQuery),
		chromedp.SendKeys(`#inputPassword`, password+"\n", chromedp.ByQuery),
		chromedp.Sleep(time.Second*1),
	)

	return err
}

//Comment places a comment into the discussion located on URL.
func Comment(ctx *context.Context, cancel *context.CancelFunc, url, title, comment string) error {
	err := chromedp.Run(*ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*1),

		chromedp.WaitVisible(`#_titleTextBox`, chromedp.ByQuery),
		chromedp.SendKeys(`#_titleTextBox`, title, chromedp.ByQuery),
		chromedp.SendKeys(`#MainContentPlaceHolder__textTextBox`, comment, chromedp.ByQuery),
		chromedp.Click(`input[type="submit"]`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*100),
	)

	return err
}
