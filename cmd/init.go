package cmd

import (
	"fmt"
	"strings"
	"time"

	utils "github.com/BrandenM-OW/daily-notes/tools"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new daily-notes directory",
	Long: `Initialize a new daily-notes directory. For example:
	
daily-notes init

Creates the following directory structure:
root/
├── config.yaml
├── notes
│  └── month
│    └── week
│      └── day.md
└── templates
	└── daily.md

`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CreateDir("notes", "notes", 0755)
		utils.CreateDir("templates", "templates", 0755)

		utils.CreateFile("config.yaml", "config.yaml", []byte(utils.ConfigTemplate), 0644)
		utils.CreateFile("templates/daily.md", "daily.md", []byte(utils.DefaultTemplate), 0644)

		month := time.Now().Month()
		utils.CreateDir("notes/"+strings.ToLower(month.String()), month.String(), 0755)

		now := time.Now()
		start := now.AddDate(0, 0, -int(now.Weekday()))
		end := start.AddDate(0, 0, 6)

		var weekRange string = start.Format("02") + "-" + end.Format("02")
		var path string = "notes/" + month.String() + "/" + weekRange + "/"
		utils.CreateDir(path, start.Format("02")+"-"+end.Format("02"), 0755)

		fmt.Println("daily-notes directory initialized")
	},
}
