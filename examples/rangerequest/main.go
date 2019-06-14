package main

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"

	bufra "github.com/avvmoto/buf-readerat"
	"github.com/snabb/httpreaderat"
)

func main() {
	// req, err := http.NewRequest("GET", "https://dl.google.com/go/go1.10.windows-amd64.zip", nil)
	// https://webhook.site/8b236e4f-b35c-49ac-95ac-29c3367be2fa
	req, err := http.NewRequest("GET", "https://webhook.site/8b236e4f-b35c-49ac-95ac-29c3367be2fa", nil)
	if err != nil {
		log.Fatal(err)
	}

	htrdr, err := httpreaderat.New(nil, req, nil) // client *http.Client, req *http.Request, bs Store
	if err != nil {
		log.Fatal(err)
	}
	bhtrdr := bufra.NewBufReaderAt(htrdr, 1024*1024)

	// https://golang.org/pkg/archive/zip/#NewReader
	zrdr, err := zip.NewReader(bhtrdr, htrdr.Size())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range zrdr.File {
		fmt.Println(f.Name)
	}
}
