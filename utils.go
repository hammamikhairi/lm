package main

import "fmt"

const (
	MEM_SPACE int = 8
)

type (
	ExtMap  map[string]IconCol
	IconCol struct {
		Icon, Color string
	}
)

var (
	dirPath string = ""
	termGap string = ""
	ICONS   ExtMap = ExtMap{
		"go":    IconCol{"", "00CDCD"},
		"jsx":   IconCol{"", "00CDCD"},
		"ts":    IconCol{"", "00CDCD"},
		"tsx":   IconCol{"", "00CDCD"},
		"py":    IconCol{"", "548588"},
		"mod":   IconCol{"", "83a598"},
		"sum":   IconCol{"", "83a598"},
		"sh":    IconCol{"", "B16286"},
		"json":  IconCol{"", "d79921"},
		"sass":  IconCol{"", "b16286"},
		"js":    IconCol{"", "FABD2F"},
		"md":    IconCol{"", "98971A"},
		"css":   IconCol{"", "98971A"},
		"yaml":  IconCol{"", "FE8019"},
		"yml":   IconCol{"", "FE8019"},
		"txt":   IconCol{"", "EBDBB2"},
		"html":  IconCol{"", "D65D0E"},
		"php":   IconCol{"", "B16286"},
		"log":   IconCol{"", "EBDBB2"},
		"jpg":   IconCol{"", "B8BB26"},
		"tmp":   IconCol{"", "CC241D"},
		"ipynb": IconCol{"", "b16286"},
		"pdf":   IconCol{"", "EBDBB2"},
		"csv":   IconCol{"", "928374"},

		"DIR":       IconCol{"", "928374"},
		"gitignore": IconCol{"", "D65D0E"},
		"git":       IconCol{"", "D65D0E"},
		"ignore":    IconCol{" ", "D65D0E"},

		"404":  IconCol{"", "000000"},
		"TEST": IconCol{"", "D79921"},
	}

	SPECIAL_ICONS ExtMap = ExtMap{
		"md":           IconCol{"", "98971A"},
		"LICENSE":      IconCol{"", "CC241D"},
		"package.json": IconCol{"", "CC241D"},
	}

	SIZES map[int]string = map[int]string{
		0: "",
		1: " K",
		2: " M",
		3: " G",
	}
)

const BLOC_SIZE float64 = 1024.0

func FormatSize(size float64, depth int) string {
	if size >= BLOC_SIZE {
		return FormatSize(float64(size)/BLOC_SIZE, depth+1)
	}

	return fmt.Sprintf("%.2f%s",
		size,
		SIZES[depth],
	)
}