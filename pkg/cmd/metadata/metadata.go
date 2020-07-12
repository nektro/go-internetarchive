package metadata

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	. "github.com/nektro/internetarchive/pkg/util"

	"github.com/spf13/cobra"
)

func init() {
	//
}

// Cmd is the cobra.Command
var Cmd = &cobra.Command{
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
