package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	err := driver.Start()
  
	if err != nil {
		log.Fatal(err)
	}

	page, err := driver.NewPage()
	err = page.Navigate("https://qiita.com/login?redirect_to=%2F")
  
	if err != nil {
		log.Fatal(err)
		log.Println(err)
		return
	}

	loginform := page.FindByID("identity")
	err = loginform.Fill("***") // your mail address
  
	if err != nil {
		log.Fatal(err)
		log.Println(err)
	}

	passform := page.FindByID("password")
	err = passform.Fill("***") // your password
	err = page.FindByXPath("/html/body/div[1]/div/div/div/div[2]/div[2]/form/div[3]/input").Click()

	if err != nil {
		log.Fatal(err)
		log.Println(err)
	}

	err = page.Navigate("https://qiita.com/")
  
	if err != nil {
		log.Fatal(err)
		log.Println(err)
	}

	searchform := page.FindByClass("st-Header_searchInput")

	err = searchform.Fill("社会人１年目に何をしたかまとめてみた")
  
	if err != nil {
		log.Fatal(err)
	}

	err = page.Navigate("https://qiita.com/search?q=%E7%A4%BE%E4%BC%9A%E4%BA%BA%EF%BC%91%E5%B9%B4%E7%9B%AE%E3%81%AB%E4%BD%95%E3%82%92%E3%81%97%E3%81%9F%E3%81%8B%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%81%9F")
	
  if err != nil {
		log.Fatal(err)
	}

	err = page.FindByLink("社会人１年目に何をしたかまとめてみた").Click()
	
  if err != nil {
		log.Fatal(err)
	}

	err = page.FindByClass("it-Footer_like").Click()
	
  if err != nil {
		log.Fatal(err)
	}
  
}