package etree

import (
	"fmt"
	"github.com/beevik/etree"
	"testing"
)

func TestName(t *testing.T) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("bookstore.xml"); err != nil {
		panic(err)
	}

	root := doc.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}

	for _, t := range doc.FindElements("//book[@category='WEB']/title") {
		fmt.Println("Title:", t.Text())
	}

	for _, e := range doc.FindElements("./bookstore/book[1]/*") {
		fmt.Printf("%s: %s\n", e.Tag, e.Text())
	}

	path := etree.MustCompilePath("./bookstore/book[p:price='49.99']/title")
	for _, e := range doc.FindElementsPath(path) {
		fmt.Println(e.Text())
	}
}
