package util

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	GoVersion      string
	GoVersionShort string
	GoPaths        []string

	Has_godoc        bool
	Has_gofmt        bool
	Has_goimports    bool
	Has_goreturns    bool
	Has_guru         bool
	Has_gorename     bool
	Has_godef        bool
	Has_gogetdoc     bool
	Has_gocode       bool
	Has_structlayout bool
	Has_godocdown    bool

	Has_golint      bool
	Has_checkvar    bool
	Has_checkalign  bool
	Has_checkstruct bool
	Has_errcheck    bool
	Has_ineffassign bool
	Has_interfacer  bool
	Has_unparam     bool
	Has_unindent    bool
	Has_unconvert   bool
	Has_maligned    bool
	Has_goconst     bool
	Has_gosimple    bool
	Has_unused      bool
	Has_staticcheck bool
	Has_deadcode    bool
)

//	Returns all paths listed in the `GOPATH` environment variable, for users who don't care about calling HasGoDevEnv.
func AllGoPaths() []string {
	if len(GoPaths) == 0 {
		GoPaths = filepath.SplitList(os.Getenv("GOPATH"))
	}
	return GoPaths
}

func DirPathToImportPath(dirpath string) string {
	for _, gopath := range AllGoPaths() {
		if strings.HasPrefix(dirpath, gopath) {
			return dirpath[len(filepath.Join(gopath, "src"))+1:]
		}
	}
	return ""
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src` sub-directory.
//	Example: `GopathSrc("tools", "importers", "sql")` yields `c:\gd\src\tools\importers\sql` if `$GOPATH` is `c:\gd`.
func GopathSrc(subDirNames ...string) (gps string) {
	gp := []string{"", "src"}
	for _, goPath := range AllGoPaths() { // in 99% of setups there's only 1 GOPATH, but hey..
		gp[0] = goPath
		if gps = filepath.Join(append(gp, subDirNames...)...); DirExists(gps) {
			break
		}
	}
	return
}

//	Returns the `path/filepath.Join`-ed full directory path for a specified `$GOPATH/src/github.com` sub-directory.
//	Example: `GopathSrcGithub("goutil", "num")` yields `c:\gd\src\github.com\goutil\num` if `$GOPATH` is `c:\gd`.
func GopathSrcGithub(gitHubName string, subDirNames ...string) string {
	return GopathSrc(append([]string{"github.com", gitHubName}, subDirNames...)...)
}
