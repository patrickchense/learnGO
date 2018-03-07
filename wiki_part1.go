package main

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	out, err := os.OpenFile(p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer out.Close()

	ow := bufio.NewWriter(out)
	ow.Write(p.Body)
	ow.Flush()

}

func (this *Page) save1() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Body))
}

/*
# Page
## Section1
This is section1.
 */