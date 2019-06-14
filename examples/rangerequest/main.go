package main

import (
	"archive/zip"
	"fmt"
	"net/http"

	bufra "github.com/avvmoto/buf-readerat"
	"github.com/snabb/httpreaderat"
)

func main() {
	req, _ := http.NewRequest("GET", "https://dl.google.com/go/go1.10.windows-amd64.zip", nil)

	htrdr, err := httpreaderat.New(nil, req, nil)
	if err != nil {
		panic(err)
	}
	bhtrdr := bufra.NewBufReaderAt(htrdr, 1024*1024)

	zrdr, err := zip.NewReader(bhtrdr, htrdr.Size())
	if err != nil {
		panic(err)
	}
	for _, f := range zrdr.File {
		fmt.Println(f.Name)
	}
}
