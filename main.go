package main

import "os"

func main() {
	emuCleaner := cleaner{}

	args := os.Args[1:]

	if len(args) < 1 {
		emuCleaner.printInstructions("No directory set!")
		return
	}

	dir := args[0]

	if len(args) > 1 {
		for _, arg := range os.Args[2:] {
			if arg == "-t" {
				emuCleaner.cleanTranslations()
			} else if arg == "-p" {
				emuCleaner.cleanPirates()
			} else if arg == "-h" {
				emuCleaner.cleanHacks()
			} else if arg == "-d" {
				emuCleaner.dryRunOnly()
			}
		}
	}

	emuCleaner.cleanDir(dir)
}
