package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/nektro/internetarchive/pkg/cmd"
	. "github.com/nektro/internetarchive/pkg/util"

	"github.com/spf13/cobra"
)

type iaMeta struct {
	A string   `xml:"identifier,omitempty" json:"identifier,omitempty"`
	G string   `xml:"title,omitempty" json:"title,omitempty"`
	B string   `xml:"mediatype,omitempty" json:"mediatype,omitempty"`
	C []string `xml:"collection,omitempty" json:"collection,omitempty"`
	D string   `xml:"description,omitempty" json:"description,omitempty"`
	H string   `xml:"publicdate,omitempty" json:"publicdate,omitempty"`
	I string   `xml:"uploader,omitempty" json:"uploader,omitempty"`
	J string   `xml:"addeddate,omitempty" json:"addeddate,omitempty"`
	L string   `xml:"backup_location,omitempty" json:"backup_location,omitempty"`
	//
	E string `xml:"scanner,omitempty" json:"scanner,omitempty"`
	F string `xml:"subject,omitempty" json:"subject,omitempty"`
	K string `xml:"curation,omitempty" json:"curation,omitempty"`
	M string `xml:"external_metadata_update,omitempty" json:"external_metadata_update,omitempty"`
	//
	N string   `xml:"creator,omitempty" json:"creator,omitempty"`
	O string   `xml:"hidden,omitempty" json:"hidden,omitempty"`
	P []string `xml:"updater,omitempty" json:"updater,omitempty"`
	Q []string `xml:"updatedate,omitempty" json:"updatedate,omitempty"`
	R string   `xml:"nav_order,omitempty" json:"nav_order,omitempty"`
	S string   `xml:"spotlight_identifier,omitempty" json:"spotlight_identifier,omitempty"`
	T string   `xml:"publisher,omitempty" json:"publisher,omitempty"`
	U string   `xml:"num_top_ba,omitempty" json:"num_top_ba,omitempty"`
	V []string `xml:"related_collection,omitempty" json:"related_collection,omitempty"`
	W string   `xml:"show_search_by_year,omitempty" json:"show_search_by_year,omitempty"`
}

var cmdMetadata = &cobra.Command{
	Use:   "metadata",
	Short: "retrieve metadata for items and collections",
	Run: func(c *cobra.Command, args []string) {
		Assert(len(args) > 0, "missing item identifier")
		name := args[0]
		_, bys, ok := GetDoc("https://archive.org/download/"+name+"/"+name+"_meta.xml", nil)
		if !ok {
			LogError("errer finding metadata for item: " + name)
			return
		}
		data := &iaMeta{}
		xml.Unmarshal(bys, data)
		jsn, _ := json.MarshalIndent(data, "", "    ")
		fmt.Println(string(jsn))
	},
}

func init() {
	cmd.Root.AddCommand(cmdMetadata)
}
