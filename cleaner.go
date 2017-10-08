package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type cleaner struct {
	dir    string
	dryRun bool

	validRegex             *regexp.Regexp
	invalidRegex           *regexp.Regexp
	removeHacks            bool
	removeHackRegex        *regexp.Regexp
	removeTranslations     bool
	removeTranslationRegex *regexp.Regexp
	removePirates          bool
	removePirateRegex      *regexp.Regexp

	totalCleaned int
	totalFiles   int
}

func (c *cleaner) cleanDir(dir string) {
	c.dir = dir
	files, err := ioutil.ReadDir(c.dir)
	c.totalCleaned = 0
	c.totalFiles = 0

	if err != nil {
		errorString := "Failed opening directory: " + err.Error()
		c.printInstructions(errorString)
		return
	}

	c.validRegex = regexp.MustCompile(`\[.*?!.*?\]`)
	c.invalidRegex = regexp.MustCompile(`\[(b|o)\d?\d?\]`)

	for _, file := range files {
		filename := file.Name()
		if filename[0:1] == "." {
			continue
		}

		c.validateFile(filename)
	}

	fmt.Printf("Removed %d out of %d total files.\n", c.totalCleaned, c.totalFiles)
}

func (c *cleaner) printInstructions(err string) {
	if err != "" {
		fmt.Println("Error: " + err)
	}

	fmt.Println("Usage: emuromcleanup.exe <directory>")
	fmt.Println("Optional arguments:\n\t-d\t\tDry run only, don't delete files\n\t-t\t\tRemove translations\n\t-h\t\tRemove hacks\n\t-p\t\tRemove pirated versions")
}

func (c *cleaner) validateFile(file string) {

	c.totalFiles++

	if c.validRegex.MatchString(file) {
		return
	}

	shouldClean := false

	if c.invalidRegex.MatchString(file) {
		shouldClean = true
	}

	if c.removeHacks == true && c.removeHackRegex.MatchString(file) {
		shouldClean = true
	}

	if c.removePirates == true && c.removePirateRegex.MatchString(file) {
		shouldClean = true
	}

	if c.removeTranslations == true && c.removeTranslationRegex.MatchString(file) {
		shouldClean = true
	}

	if shouldClean {
		c.totalCleaned++
		fmt.Println("rm " + c.dir + "/" + file)

		if c.dryRun == false {
			os.Remove(c.dir + "/" + file)
		}
	}
}

func (c *cleaner) cleanTranslations() {
	c.removeTranslations = true
	c.removeTranslationRegex = regexp.MustCompile(`\[T(\+|\-).*?\]`)
}

func (c *cleaner) cleanPirates() {
	c.removePirates = true
	c.removePirateRegex = regexp.MustCompile(`\[p.*?\]`)
}

func (c *cleaner) cleanHacks() {
	c.removeHacks = true
	c.removeHackRegex = regexp.MustCompile(`\[h.*?\]`)
}

func (c *cleaner) dryRunOnly() {
	c.dryRun = true
}
