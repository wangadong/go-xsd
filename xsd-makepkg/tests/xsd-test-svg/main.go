package main

import (
	"encoding/xml"

	"github.com/wangadong/go-xsd/xsd-makepkg/tests"

	svg "github.com/wangadong/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
)

type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func main() {
	var (
		dirBasePath  = tests.GetExecutableDir()
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
