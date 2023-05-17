package html2image

import (
	"bytes"
	"log"

	"golang.org/x/net/html"
)

func Html2Image(htmlBytes []byte, fontPath string) ([]byte, error) {
	htmlIoReader := bytes.NewReader(htmlBytes)
	htmlNode, err := html.Parse(htmlIoReader)
	if err != nil {
		log.Fatal(err)
	}

	body, styleList := GetBodyStyle(htmlNode)

	var styleString []string
	for _, value := range styleList {
		styleString = append(styleString, value.FirstChild.Data)
	}
	tagStyleList := ParseStyle(styleString, fontPath)

	parsedBodyDom := GetHtmlDom(body, tagStyleList)

	return bodyDom2Image(parsedBodyDom)
}
