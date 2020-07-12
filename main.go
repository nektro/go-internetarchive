package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nektro/internetarchive/pkg/cmd"

	"github.com/PuerkitoBio/goquery"
	"github.com/nektro/go-util/ansi/style"
)

// Version takes in version string from build_all.sh
var Version = "vMASTER"

func main() {
	dieOnError(cmd.Root.Execute())
}

func dieOnError(err error, args ...string) {
	if err != nil {
		logError(err.Error())
		for _, item := range args {
			logError(item)
		}
		os.Exit(1)
	}
}

func logError(err string) {
	fmt.Print(style.FgRed)
	log.Println(err)
	fmt.Print(style.ResetFgColor)
}

func assert(b bool, msg string) {
	if !b {
		dieOnError(errors.New(msg))
	}
}

func getDoc(urlS string, hdrs map[string]string) (*goquery.Document, []byte, bool) {
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
