package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ateyun-com/html2image"
	"golang.org/x/net/html"
	"log"
	"os"
)

func OutputImg() {
	htmlPath := "./example.html"
	fontPath := "./fonts"
	htmlBytes, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Fatal(err)
	}
	imgByte, err := html2image.Html2Image(htmlBytes, fontPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fh, err := os.Create("./generated.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = fh.Chmod(0755)
	_, _ = fh.Write(imgByte)
	_ = fh.Close()
}

func ExportJson() {
	htmlPath := "./example.html"
	fontPath := "./fonts"
	htmlBytes, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Fatal(err)
	}
	htmlIoReader := bytes.NewReader(htmlBytes)
	htmlNode, err := html.Parse(htmlIoReader)
	if err != nil {
		log.Fatal(err)
	}

	body, styleList := html2image.GetBodyStyle(htmlNode)

	var styleString []string
	for _, value := range styleList {
		styleString = append(styleString, value.FirstChild.Data)
	}
	tagStyleList := html2image.ParseStyle(styleString, fontPath)

	parsedBodyDom := html2image.GetHtmlDom(body, tagStyleList)

	jsonStr, err := json.MarshalIndent(parsedBodyDom, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fp, err := os.Create("./example.json")
	_ = fp.Chmod(0755)
	_, _ = fp.Write(jsonStr)
	_ = fp.Close()
}
