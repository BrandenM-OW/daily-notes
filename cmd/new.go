package cmd

import (
	"strings"
	"time"

	utils "github.com/BrandenM-OW/daily-notes/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("preserve", "p", false, "Preserve notes from last day")
	newCmd.Flags().IntP("days", "d", 0, "Create a note for a day in the future")
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new daily note",
	Long: `Create a new daily note. For example:

daily-notes new
daily-notes new -p // Preserves notes from last day
daily-notes new -d 1 // Creates a note for tomorrow

Adds the following file:
root/
└── notes
    └── month
        └── week
            └── day.md

The day.md file will be created with based on the template specified in the config.yaml file.

If the preserve flag is set, the last day.md file will be used as a template for the new day.md file.
	- All unchecked tasks will be copied over
	- All notes will be copied over that are not marked with the -p flag


`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		if days, _ := cmd.Flags().GetInt("days"); days != 0 {
			now = now.AddDate(0, 0, days)
		}

		start := now.AddDate(0, 0, -int(now.Weekday()))
		end := start.AddDate(0, 0, 6)
		today := now.Format("02")

		var month string = now.Month().String()
		var weekRange string = start.Format("02") + "-" + end.Format("02")
		var day string = today
		var path string = "notes/" + month + "/" + weekRange + "/" + day + ".md"

		if utils.FileExists(path) {
			return
		}

		template := utils.GetTemplate()

		Preserve, _ := cmd.Flags().GetBool("preserve")
		if Preserve {
			lastPath := viper.Get("last").(string)
			lastPath = strings.ToLower(lastPath)
			if utils.FileExists(lastPath) {
				yesterdayTasks := utils.GetTasks(lastPath)
				yesterdayNotes := utils.GetNotes(lastPath)

				template = utils.AppendTasks(template, yesterdayTasks)
				template = utils.AppendNotes(template, yesterdayNotes)
				utils.CreateFile(path, today, []byte(template), 0644)

			} else {
				utils.CreateFile(path, today, []byte(template), 0644)
			}
		} else {
			utils.CreateFile(path, today, []byte(template), 0644)
		}

		if Preserve {
			viper.Set("last", path)
			viper.WriteConfig()
		}
	},
}
