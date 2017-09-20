package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
)

/*
	Prints a report generated from the most recent notes while skipping some entries.
	This script expects the notes to be formated with Markdown bullet points like so:

	# 📆 01.01.1970
	* 💻 Paired with @someone
	  * Did some tech stuff
	  * ⏳🎽 Start doing some thing [← skipped retro entry]
	  * 💫 Made progress on tech stuff

	* 💬 Some feedback from @someone [← skipped feedback entry]
	  * 🔄 Do some thing more regularly [← skipped sub entry of feedback entry]
*/
func main() {
	const NumberOfNotesToDisplay = 5
	const SkippedEntryMarkers = "🔖🏷💽💾🛠🔧📐📑📄📖⏳🏆💬📰"
	const NotesGlobExp = "../archive/*.md"

	isTopLevelEntryRegExp := regexp.MustCompile("^\\*")
	isSkippedLineRegExp := regexp.MustCompile("\\*\\ [" + SkippedEntryMarkers + "]")

	allNotes, err := filepath.Glob(NotesGlobExp)
	check(err)

	totalNumberOfNotes := len(allNotes)

	if totalNumberOfNotes == 0 {
		fmt.Println("No notes found in ", NotesGlobExp)
		return
	}

	mostRecentNotes := allNotes[int(math.Max(float64(totalNumberOfNotes-NumberOfNotesToDisplay), 0)):]

	isFirstNote := true

	for _, note := range mostRecentNotes {
		if !isFirstNote {
			fmt.Println()
			fmt.Println(" --- --- --- --- --- --- ")
			fmt.Println()
		}

		file, err := os.Open(note)
		check(err)
		defer file.Close()

		skipSubEntries := false

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if isTopLevelEntryRegExp.MatchString(line) {
				skipSubEntries = false
			}

			if isSkippedLineRegExp.MatchString(line) {
				if isTopLevelEntryRegExp.MatchString(line) {
					// Skipping all sub entries of the current entry
					skipSubEntries = true
				}
			} else {
				if !skipSubEntries {
					fmt.Println(line)
				}
			}
		}

		isFirstNote = false
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
