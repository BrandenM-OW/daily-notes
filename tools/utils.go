package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const ConfigTemplate = `
template: daily.md
last: /path/to/last/day.md
`
const DefaultTemplate = `
Daily Notes
==========

## Tasks
- [ ] Task 1
- [ ] Task 2
- [ ] Task 3

## Notes
-p Note 1
- Note 2
- Note 3
`
const PreserveTemplate = `
Daily Notes
==========
`

func CreateDir(path string, name string, perm os.FileMode) {
	err := os.Mkdir(path, perm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name + " directory created")
	}
}

func CreateFile(path string, name string, data []byte, perm os.FileMode) {
	parts := strings.Split(path, "/")
	for i := 0; i < len(parts)-1; i++ {
		dir := strings.Join(parts[:i+1], "/")
		if !FileExists(dir) {
			CreateDir(strings.ToLower(dir), dir, 0755)
		}
	}

	err := os.WriteFile(path, data, perm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name + " file created")
	}
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetTemplate() string {
	templateName := viper.Get("template").(string)

	template, err := os.ReadFile("templates/" + templateName)
	if err != nil {
		fmt.Println(err)
	}

	return string(template)
}

func GetTasks(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var tasks []string
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, "- [ ] ") {
			tasks = append(tasks, line[6:])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return tasks
}

func GetNotes(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var notes []string
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, "-p") {
			notes = append(notes, line[2:])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return notes
}

func AppendTasks(template string, tasks []string) string {
	var builder strings.Builder
	foundTasks := false

	for _, line := range strings.Split(template, "\n") {
		if strings.HasPrefix(line, "## Tasks") {
			foundTasks = true
		}
		builder.WriteString(line)
		builder.WriteString("\n")
		if foundTasks {
			for _, task := range tasks {
				builder.WriteString("- [ ] ")
				builder.WriteString(task)
				builder.WriteString("\n")
			}
			foundTasks = false
		}
	}

	return builder.String()
}

func AppendNotes(template string, notes []string) string {
	var builder strings.Builder
	foundNotes := false

	for _, line := range strings.Split(template, "\n") {
		if strings.HasPrefix(line, "## Notes") {
			foundNotes = true
		}
		builder.WriteString(line)
		builder.WriteString("\n")
		if foundNotes {
			for _, note := range notes {
				builder.WriteString("-p")
				builder.WriteString(note)
				builder.WriteString("\n")
			}
			foundNotes = false
		}
	}

	return builder.String()
}
