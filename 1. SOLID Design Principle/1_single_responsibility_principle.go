package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCounnt = 0

// Journal ...
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// AddEntry ...
func (j *Journal) AddEntry(text string) int {
	entryCounnt++
	entry := fmt.Sprintf("%d: %s", entryCounnt, text)
	j.entries = append(j.entries, entry)
	return entryCounnt
}

// RemoveEntry ..
func (j *Journal) RemoveEntry(index int) {
	//
}

// separation of concerns

// Save ...
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

// Load ...
func (j *Journal) Load(filename string) {

}

// LoadfromWeb ...
func (j *Journal) LoadfromWeb(url *url.URL) {

}

// LineSeparator ...
var LineSeparator = "\n"

// SaveToFile ...
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// Persistence ...
type Persistence struct {
	lineSeparator string
}

// SaveToFile ...
func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// separate function
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}
