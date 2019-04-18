package main

import (
	// "Keys"
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		log.Fatal(err)
	}
	page, err := driver.NewPage(agouti.Browser("chrome"))
	err = page.Navigate("https://qiita.com/")
	if err != nil {
		log.Fatal(err)
	}
	searchform := page.FindByClass("st-Header_searchInput")
	err = searchform.Fill("社会人１年目に何をしたかまとめてみた")
	if err != nil {
		log.Fatal(err)
	}
	// btn := page.send_keys(Keys.ENTER)
}
