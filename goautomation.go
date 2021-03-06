package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
	"github.com/season-three/hiromae_week_5/secret"
)

func main() {
	fmt.Println(secret.MyPassword())

	url := "https://www.linkedin.com/login/ja?fromSignIn=true&trk=guest_homepage-basic_nav-header-signin"
	driver := agouti.ChromeDriver()

	//ドライバの起動＋スタート
	err := driver.Start()
	if err != nil {
		fmt.Println("Failed to start driver")
	}
	//defer driver.Stop()

	//クロームの起動＋ページを開く
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		fmt.Println("Failed to open new page")
	}

	//ログインページ開く
	err = page.Navigate(url)

	if err != nil {
		fmt.Println("Failed to open login page")
	}

	page.Session().SetImplicitWait(10)

	//もしログイン画面が出てきたら

	//usernameとpasswordのIDを取得
	username := page.FindByID("username")
	password := page.FindByID("password")

	//usernameとpasswordを入力
	username.Fill(secrets.Myusername())
	password.Fill(secrets.Mypassword())

	//サインインボタンを押す
	err = page.FindByButton("サインイン").Submit()
	if err != nil {
		fmt.Println("Failed to login")
	}

	//つながり申請する（終わるまで）

	//処理終了後、3秒間ブラウザを表示
	time.Sleep(3 * time.Second)

}
