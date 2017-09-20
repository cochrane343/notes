package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

/*
  Creates a new note file in the archive folder if there is none yet, prepopulates it with
	template content and opens the file in a text editor.
*/
func main() {
	const NotesDirectory = "../archive"
	const Editor = "atom"

	now := time.Now()
	fileName := NotesDirectory + "/note_" + now.Format("2006_01_02") + ".md"
	dateHeader := "# 📆 " + now.Format("02.01.2006")

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("No note for today found - creating one ...")

		templateData, err := ioutil.ReadFile("template.md")
		check(err)

		f, err := os.Create(fileName)
		check(err)
		defer f.Close()

		f.WriteString(dateHeader)
		f.WriteString("\n\n")
		f.WriteString("* 💻 Some work stuff ...")
		f.WriteString("\n\n\n\n")
		f.WriteString("* 🌱 ...")
		f.WriteString("\n\n")
		f.WriteString(string(templateData))
	}

	fmt.Println("Opening note for today ...")

	editorCommand := exec.Command(Editor, fileName)
	err := editorCommand.Start()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
