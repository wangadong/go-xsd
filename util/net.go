package util

import (
	"io"
	"net/http"
)

//	Downloads a remote file at the specified (`net/http`-compatible) `srcFileUrl` to the specified `dstFilePath`.
func DownloadFile(srcFileUrl, dstFilePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(srcFileUrl); err == nil {
		defer rc.Close()
		SaveToFile(rc, dstFilePath)
	}
	return
}

//	Opens a remote file at the specified (`net/http`-compatible) `srcFileUrl` and returns its `io.ReadCloser`.
func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = new(http.Client).Get(srcFileUrl); (err == nil) && (resp != nil) {
		src = resp.Body
	}
	return
}
