package parlamentni

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/uhlobot/browser"
	"github.com/agajdosi/uhlobot/generate"
	"github.com/agajdosi/uhlobot/seznam"
	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
)

//Register registers a new account.
func Register(ctx *context.Context, cancel *context.CancelFunc, bot string) {
	password := generate.Password(10)
	mail := viper.GetString(bot + ".mail")
	name := viper.GetString(bot + ".name")
	surname := viper.GetString(bot + ".surname")
	born := viper.GetString(bot + ".born")
	city := viper.GetString(bot + ".city")
	nickname := generate.Nickname()
	nickname = name + nickname + surname

	sex := viper.GetString(bot + ".sex")
	if sex == "M" {
		sex = "Muž"
	} else {
		sex = "Žena"
	}

	chromedp.Run(*ctx,
		chromedp.Navigate("https://www.parlamentnilisty.cz/profily-sprava/ProfileRegistration.aspx"),
		chromedp.Sleep(time.Second*5),
		chromedp.WaitVisible(`#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep1ChooseType_ascx_txtEmail_I`, chromedp.ByQuery),
		//email
		chromedp.SendKeys(`#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep1ChooseType_ascx_txtEmail_I`, mail, chromedp.ByQuery),
		//password
		chromedp.SendKeys(`#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep1ChooseType_ascx_ucPassword_txtPassword1_I`, password, chromedp.ByQuery),
		//password verification
		chromedp.SendKeys(`#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep1ChooseType_ascx_ucPassword_txtPassword2_I`, password, chromedp.ByQuery),
		chromedp.Click("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep1ChooseType_ascx_chbOperationgConditions_I", chromedp.ByQuery),
		chromedp.Click("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_btnNextStep_CD", chromedp.ByQuery),

		chromedp.Sleep(time.Second*5),
		chromedp.WaitVisible("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep2CitizenPersonalInfo_ascx_txtNickname_I", chromedp.ByQuery),
		//nickname
		chromedp.SendKeys("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep2CitizenPersonalInfo_ascx_txtNickname_I", nickname, chromedp.ByQuery),
		//datum narozeni
		chromedp.SendKeys("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep2CitizenPersonalInfo_ascx_deBirthdate_I", born, chromedp.ByQuery),
		//pohlavi
		chromedp.SendKeys("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep2CitizenPersonalInfo_ascx_ddlGender_I", sex, chromedp.ByQuery),
		//mesto
		chromedp.SendKeys("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_RegistrationStep2CitizenPersonalInfo_ascx_ddlCity_I", city, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.Click("#ctl00_ctl00_cphBody_cphBody_ProfileRegistrationControl1_WizardDialogControl1_panelBody_panelBodyPlaceHolder_btnNextStep_CD", chromedp.ByQuery),

		chromedp.Sleep(time.Second*2),
	)

	fmt.Println("account registered:", mail, password)
}

//RegisterAll registers all bots who are not yet registered.
func RegisterAll() {
	bots := viper.GetStringSlice("bots")
	for _, bot := range bots {
		if viper.Get(bot+".parlamentni.password") != nil {
			continue
		}

		ctx, cancel := browser.CreateBrowser()
		Register(ctx, cancel, bot)
		seznam.Login()

		chromedp.Cancel(*ctx)
		fmt.Println("success!")
	}
}

//Login logins a user into the web.
func Login(ctx *context.Context, cancel *context.CancelFunc, bot string) error {

	username := viper.GetString(bot + ".mail")
	password := viper.GetString(bot + ".parlamentni.password")

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
