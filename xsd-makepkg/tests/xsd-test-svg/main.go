package main

import (
	"encoding/xml"

	"github.com/wangadong/goxsd/xsd-makepkg/tests"

	udevgo "github.com/wangadong/goxsd/goutil/dev/go"

	svg "github.com/wangadong/goxsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
)

type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-svg")
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
