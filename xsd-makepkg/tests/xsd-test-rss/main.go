package main

import (
	"encoding/xml"

	"github.com/wangadong/go-xsd/xsd-makepkg/tests"

	util "github.com/wangadong/go-xsd/util"

	rss "github.com/wangadong/go-xsd-pkg/thearchitect.co.uk/schemas/rss-2_0.xsd_go"
)

type RssDoc struct {
	XMLName xml.Name `xml:"rss"`
	rss.TxsdRss
}

func main() {
	var (
		dirBasePath  = util.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-rss")
		makeEmptyDoc = func() interface{} { return &RssDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
