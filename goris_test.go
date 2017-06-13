package main

import (
	"testing"

	"github.com/tanaikech/goris/ris"
)

func TestImgFromURL(t *testing.T) {
	imageurl := "https://github.com/tanaikech/goris/blob/master/myavatar.png?raw=true"
	webpages := true
	results := ris.DefImg(webpages).ImgFromURL(imageurl)
	t.Log(results)
}
