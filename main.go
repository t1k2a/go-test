package main

import(
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Faild to start river:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	if err := page.Navigate("https://achieve-t1k2a.c9users.io/blogs"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	page.Screenshot("/tmp/chetome_qiita.jpg")
	}

