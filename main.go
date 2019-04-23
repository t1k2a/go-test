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
		log.Println(err)
	}

	err := page.Navigate("https://achieve-t1k2a.c9users.io/blogs");
	if err != nil {
		log.Println(err)
	}

	err := page.FindByLink("Open the App").Click()
	if err != nil {
		log.Println(err)
	}

	err : = page.Screenshot("/tmp/chetome_qiita.jpg")
	if err != nil {
		log.Println(err)
	}

}