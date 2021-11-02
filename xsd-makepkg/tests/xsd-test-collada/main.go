package main

import (
	"encoding/xml"
	"log"
	"path/filepath"

	"github.com/wangadong/go-xsd/xsd-makepkg/tests"

	collada14 "github.com/wangadong/go-xsd-pkg/khronos.org/files/collada_schema_1_4_go"
	collada15 "github.com/wangadong/go-xsd-pkg/khronos.org/files/collada_schema_1_5_go"
)

type Col14Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	collada14.TxsdCollada
}

type Col15Doc struct {
	XMLName xml.Name `xml:"COLLADA"`
	collada15.TxsdCollada
}

func main() {
	var (
		col14DirBasePath  = filepath.Join(tests.GetExecutableDir(), "1.4.1")
		col14MakeEmptyDoc = func() interface{} { return &Col14Doc{} }
		col15DirBasePath  = filepath.Join(tests.GetExecutableDir(), "1.5")
		col15MakeEmptyDoc = func() interface{} { return &Col15Doc{} }
	)
	if false {
		tests.OnDocLoaded = func(doc interface{}) {
			if c14, ok := doc.(*Col14Doc); ok {
				log.Print("ISC14")
				for _, camLib := range c14.CamerasLibraries {
					log.Print("CAMLIB")
					for _, cam := range camLib.Cameras {
						log.Printf("CAM aspect: %#v\n", cam.Optics.TechniqueCommon.Perspective.AspectRatio)
					}
				}
			}
		}
	}
	tests.TestViaRemarshal(col14DirBasePath, col14MakeEmptyDoc)
	tests.TestViaRemarshal(col15DirBasePath, col15MakeEmptyDoc)
}
