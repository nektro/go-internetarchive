package download

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	. "github.com/nektro/internetarchive/pkg/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/nektro/go-util/arrays/stringsu"
	"github.com/nektro/go-util/mbpp"
	"github.com/nektro/go-util/util"
	"github.com/spf13/cobra"
)

var (
	onlyMeta bool
	dense    bool
	nSOrig   bool
	nSMeta   bool
	ySDerv   bool

	sources = []string{"original", "metadata"}
)

func init() {
	//
	Cmd.Flags().StringP("save-dir", "o", "./data", "")
	Cmd.Flags().BoolVar(&onlyMeta, "only-meta", false, "when enabled, only saves _meta.xml files")
	Cmd.Flags().BoolVar(&dense, "dense", false, "when enabled, stores items based on their creation date")
	Cmd.Flags().IntP("concurrency", "c", 10, "number of concurrent download jobs to run at once")
	Cmd.Flags().BoolVar(&nSOrig, "no-original", false, "when enabled, does not save items with a source of original")
	Cmd.Flags().BoolVar(&nSMeta, "no-metadata", false, "when enabled, does not save items with a source of metadata")
	Cmd.Flags().BoolVar(&ySDerv, "yes-derivative", false, "when enabled, does save items with a source of derivative")
}

// Cmd is the cobra.Command
var Cmd = &cobra.Command{
	Use:   "download",
	Short: "download an item or collection",
	Run: func(c *cobra.Command, args []string) {
		Assert(len(args) > 0, "missing item identifier")
		p, _ := c.Flags().GetString("save-dir")
		cc, _ := c.Flags().GetInt("concurrency")
		nso, _ := c.Flags().GetBool("no-original")
		nsm, _ := c.Flags().GetBool("no-metadata")
		ysd, _ := c.Flags().GetBool("yes-derivative")

		if nso {
			sources = stringsu.Remove(sources, "original")
		}
		if nsm {
			sources = stringsu.Remove(sources, "metadata")
		}
		if ysd {
			sources = append(sources, "derivative")
		}

		d, _ := filepath.Abs(p)
		mbpp.Init(cc)
		dlItem(d, args[0], nil)
		mbpp.Wait()
		time.Sleep(time.Second)
		log.Println(mbpp.GetCompletionMessage())
	},
}

func dlItem(dir, name string, b *mbpp.BarProxy) {
	mbpp.CreateJob("item: "+name, func(bar *mbpp.BarProxy) {
		val, _, ok := GetJSON("https://archive.org/metadata/"+name, nil)
		if !ok {
			return
		}
		mt := string(val.GetStringBytes("metadata", "mediatype"))
		if len(mt) == 0 {
			return
		}
		if mt == "collection" {
			go dlCollection(dir, name)
			return
		}
		ad := string(val.GetStringBytes("metadata", "addeddate"))
		ad = ad[:strings.IndexRune(ad, ' ')]
		ad = strings.ReplaceAll(ad, "-", "/")
		dir2 := dir
		if dense {
			dir2 += "/" + ad
		}
		dir2 += "/" + name
		if util.DoesDirectoryExist(dir2) {
			return
		}
		os.MkdirAll(dir2, os.ModePerm)
		wg := new(sync.WaitGroup)
		arr := val.GetArray("files")
		for _, item := range arr {
			n := string(item.GetStringBytes("name"))
			s := string(item.GetStringBytes("source"))
			if onlyMeta {
				if n != name+"_meta.xml" {
					continue
				}
				go saveTo(dir2, name, n, b)
				return
			}
			if !stringsu.Contains(sources, s) {
				continue
			}
			bar.AddToTotal(1)
			wg.Add(1)
			go func() {
				saveTo(dir2, name, n, bar)
				wg.Done()
			}()
		}
		wg.Wait()
		if b != nil {
			b.Increment(1)
		}
	})
}

func dlCollection(dir, name string) {
	mbpp.CreateJob("collection: "+name, func(bar *mbpp.BarProxy) {
		dat := map[string]string{"x-requested-with": "XMLHttpRequest"}
		for i := 1; true; i++ {
			doc, _, _ := GetDoc("https://archive.org/details/"+name+"?&page="+strconv.Itoa(i), dat)
			arr := doc.Find(".item-ia[data-id]")
			if arr.Length() == 1 {
				break
			}
			arr.Each(func(_ int, el *goquery.Selection) {
				n, _ := el.Attr("data-id")
				if n == "__mobile_header__" {
					return
				}
				bar.AddToTotal(1)
				if onlyMeta {
					go dlItem(dir, n, bar)
					return
				}
				dlItem(dir, n, bar)
			})
		}
	})
}

func saveTo(dir, item, file string, b *mbpp.BarProxy) {
	pathS := dir + "/" + file
	os.MkdirAll(filepath.Dir(pathS), os.ModePerm)
	urlS := "https://archive.org/download/" + item + "/" + file
	mbpp.CreateDownloadJob(urlS, pathS, b)
}
