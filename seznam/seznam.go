package seznam

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/agajdosi/uhlobot/generate"
	"github.com/chromedp/chromedp"
)

func Login(ctx *context.Context, cancel *context.CancelFunc, bot string) error {
	//TBD
}

//Register creates new account on seznam.cz
func Register(ctx *context.Context, cancel *context.CancelFunc) error {
	rand.Seed(time.Now().UnixNano())
	username := generate.Username(6)
	password := generate.Password(12)
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
