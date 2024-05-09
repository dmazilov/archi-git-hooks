package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ArchimateDiagramModel struct {
	XMLName xml.Name `xml:"ArchimateDiagramModel"`
	Name    string   `xml:"name,attr"`
	Id      string   `xml:"id,attr"`
}

func modifiedDiagramsCollection() []string {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Command %s returned error: ", strings.Join(cmd.Args, " ")), err)
		if stderr.Len() > 0 {
			fmt.Println("Error message:", stderr.String())
		}
		os.Exit(1)
	}

	scanner := bufio.NewScanner(&stdout)
	diagramFilenamePattern := regexp.MustCompile(`ArchimateDiagramModel_([a-z0-9\-]+).xml`)

	var modifiedDiagrams []string
	diagramNumber := 0
	for scanner.Scan() {
		filename := scanner.Text()
		matches := diagramFilenamePattern.FindStringSubmatch(filename)
		if len(matches) > 1 {
			var diagramString string
			diagramNumber++
			diagramXmlContent, err := os.ReadFile(filename)
			if err != nil {
				if os.IsNotExist(err) {
					diagramString = fmt.Sprintf("%d) DELETED diagram [ %s ]", diagramNumber, filename)
				} else {
					fmt.Println("Couldn't read file: ", filename, err)
					os.Exit(1)
				}
			} else {
				var diagram ArchimateDiagramModel
				err = xml.Unmarshal(diagramXmlContent, &diagram)
				if err != nil {
					fmt.Println("Couldn't parse diagram xml: ", filename, err)
					os.Exit(1)
				}
				diagramString = fmt.Sprintf("%d) %s [ %s ]", diagramNumber, diagram.Name, diagram.Id)
			}
			modifiedDiagrams = append(modifiedDiagrams, diagramString)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while searching for Archi diagrams changes:", err)
		os.Exit(1)
	}

	return modifiedDiagrams
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Commit message file couldn't be found")
		os.Exit(1)
	}

	commitMsgFileName := os.Args[1]
	commitMsgFile, err := os.Open(commitMsgFileName)
	if err != nil {
		fmt.Println("Error while commit message file reading:", err)
		os.Exit(1)
	}
	defer commitMsgFile.Close()

	var commitMsgLines []string
	scanner := bufio.NewScanner(commitMsgFile)
	for scanner.Scan() {
		commitMsgLines = append(commitMsgLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while commit message file reading:", err)
		os.Exit(1)
	}

	modifiedDiagrams := modifiedDiagramsCollection()
	if len(modifiedDiagrams) > 0 {
		commitMsgLines = append(commitMsgLines, "\n\nThese Archimate diagrams were changed in current commit:\n")
		commitMsgLines = append(commitMsgLines, modifiedDiagrams...)
	}

	if err := os.WriteFile(commitMsgFileName, []byte(strings.Join(commitMsgLines, "\n")), 0644); err != nil {
		fmt.Println("Error while saving new commit message:", err)
		os.Exit(1)
	}
}
