package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/nektro/go-util/ansi/style"
)

// DieOnError kills the procss if err is not nil
func DieOnError(err error, args ...string) {
	if err != nil {
		LogError(err.Error())
		for _, item := range args {
			LogError(item)
		}
		os.Exit(1)
	}
}

// LogError does that
func LogError(err string) {
	fmt.Print(style.FgRed)
	log.Println(err)
	fmt.Print(style.ResetFgColor)
}

// Assert calls DieOnError is false
func Assert(b bool, msg string) {
	if !b {
		DieOnError(errors.New(msg))
	}
}

// GetDoc fetch and html document and parses it
func GetDoc(urlS string, hdrs map[string]string) (*goquery.Document, []byte, bool) {
	req, err := http.NewRequest(http.MethodGet, urlS, nil)
	if err != nil {
		return nil, nil, false
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, false
	}
	if res.StatusCode >= 400 {
		return nil, nil, false
	}
	bys, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, bys, false
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bys))
	if err != nil {
		return doc, bys, false
	}
	return doc, bys, true
}