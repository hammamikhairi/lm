package main

type (
	ExtMap  map[string]IconCol
	IconCol struct {
		Icon, Color string
	}
)

const (
	// Language / file type icons
	IconGo         = "\uE724" // dev-go
	IconReact      = "\uE7BA" // dev-react
	IconTypescript = "\uE8CA" // dev-typescript
	IconPython     = "\uE73C" // dev-python
	IconGoMod      = "\uEB29" // cod-package
	IconBash       = "\uE760" // dev-bash
	IconJSON       = "\uE80B" // dev-json
	IconSass       = "\uE74B" // dev-sass
	IconJS         = "\uE781" // dev-javascript
	IconMarkdown   = "\uE73E" // dev-markdown
	IconCSS        = "\uE749" // dev-css3
	IconYAML       = "\uE8EB" // dev-yaml
	IconText       = "\uEA7B" // cod-file
	IconHTML       = "\uE736" // dev-html5
	IconPHP        = "\uE73D" // dev-php
	IconLog        = "\uEB6A" // cod-three_bars
	IconImage      = "\uEAEA" // cod-file_media
	IconTmp        = "\uEB6A" // cod-three_bars
	IconJupyter    = "\uE80F" // dev-jupyter
	IconPDF        = "\uEAEB" // cod-file_pdf
	IconCSV        = "\uEBB7" // cod-table
	IconRust       = "\uE7A8" // dev-rust
	IconRuby       = "\uE739" // dev-ruby
	IconJava       = "\uE738" // dev-java
	IconC          = "\uE771" // dev-c_lang
	IconCpp        = "\uE7A3" // dev-cplusplus
	IconLua        = "\uE826" // dev-lua
	IconKotlin     = "\uE81B" // dev-kotlin
	IconSwift      = "\uE755" // dev-swift
	IconVue        = "\uE8DC" // dev-vuejs
	IconSvelte     = "\uE8B7" // dev-svelte
	IconToml       = "\uE6B2" // custom-toml
	IconXML        = "\uE8EA" // dev-xml
	IconSQL        = "\uE8B0" // dev-sqldeveloper
	IconTerraform  = "\uE8BD" // dev-terraform
	IconDocker     = "\uE7B0" // dev-docker
	IconArchive    = "\uEAEF" // cod-file_zip
	IconMusic      = "\uEC1B" // cod-music
	IconVideo      = "\uEAD9" // cod-device_camera_video
	IconLock       = "\uEA75" // cod-lock
	IconEnv        = "\uEAF8" // cod-gear
	IconHeader     = "\uEB63" // cod-symbol_misc
	IconMake       = "\uEB6D" // cod-tools
	IconINI        = "\uEB52" // cod-settings
	IconSQLite     = "\uE7C4" // dev-sqlite

	// Special icons
	IconFolder    = "\uE5FF" // custom-folder
	IconGit       = "\uE702" // dev-git
	IconGitIgnore = "\uE702" // dev-git
	IconIgnore    = " "
	IconUnknown   = "\uEB32" // cod-question
	IconTest      = "\uEA79" // cod-beaker
	IconLicense   = "\uEB12" // cod-law
	IconNPM       = "\uE71E" // dev-npm
)

var (
	ICONS ExtMap = ExtMap{
		// Go
		"go":  IconCol{IconGo, "00CDCD"},
		"mod": IconCol{IconGoMod, "83a598"},
		"sum": IconCol{IconGoMod, "83a598"},

		// JavaScript / TypeScript
		"js":  IconCol{IconJS, "FABD2F"},
		"jsx": IconCol{IconReact, "00CDCD"},
		"ts":  IconCol{IconTypescript, "00CDCD"},
		"tsx": IconCol{IconTypescript, "00CDCD"},

		// Web
		"html":   IconCol{IconHTML, "D65D0E"},
		"css":    IconCol{IconCSS, "98971A"},
		"scss":   IconCol{IconSass, "b16286"},
		"sass":   IconCol{IconSass, "b16286"},
		"vue":    IconCol{IconVue, "98971A"},
		"svelte": IconCol{IconSvelte, "D65D0E"},

		// Languages
		"py":    IconCol{IconPython, "548588"},
		"rs":    IconCol{IconRust, "D65D0E"},
		"rb":    IconCol{IconRuby, "CC241D"},
		"java":  IconCol{IconJava, "D65D0E"},
		"c":     IconCol{IconC, "83a598"},
		"cpp":   IconCol{IconCpp, "83a598"},
		"h":     IconCol{IconHeader, "83a598"},
		"lua":   IconCol{IconLua, "00CDCD"},
		"kt":    IconCol{IconKotlin, "D65D0E"},
		"swift": IconCol{IconSwift, "D65D0E"},
		"php":   IconCol{IconPHP, "B16286"},
		"sh":    IconCol{IconBash, "B16286"},

		// Config / Data
		"json":   IconCol{IconJSON, "d79921"},
		"yaml":   IconCol{IconYAML, "FE8019"},
		"yml":    IconCol{IconYAML, "FE8019"},
		"toml":   IconCol{IconToml, "FE8019"},
		"xml":    IconCol{IconXML, "D65D0E"},
		"ini":    IconCol{IconINI, "928374"},
		"env":    IconCol{IconEnv, "d79921"},
		"sql":    IconCol{IconSQL, "83a598"},
		"sqlite": IconCol{IconSQLite, "83a598"},
		"tf":     IconCol{IconTerraform, "B16286"},

		// Documents / Data
		"md":    IconCol{IconMarkdown, "98971A"},
		"txt":   IconCol{IconText, "EBDBB2"},
		"pdf":   IconCol{IconPDF, "EBDBB2"},
		"csv":   IconCol{IconCSV, "928374"},
		"ipynb": IconCol{IconJupyter, "b16286"},
		"log":   IconCol{IconLog, "EBDBB2"},

		// Images
		"jpg":  IconCol{IconImage, "B8BB26"},
		"jpeg": IconCol{IconImage, "B8BB26"},
		"png":  IconCol{IconImage, "B8BB26"},
		"gif":  IconCol{IconImage, "B8BB26"},
		"svg":  IconCol{IconImage, "B8BB26"},
		"webp": IconCol{IconImage, "B8BB26"},
		"ico":  IconCol{IconImage, "B8BB26"},

		// Media
		"mp3": IconCol{IconMusic, "B16286"},
		"mp4": IconCol{IconVideo, "B16286"},

		// Archives
		"zip": IconCol{IconArchive, "d79921"},
		"tar": IconCol{IconArchive, "d79921"},
		"gz":  IconCol{IconArchive, "d79921"},

		// Build / Container
		"lock": IconCol{IconLock, "928374"},
		"tmp":  IconCol{IconTmp, "CC241D"},

		// Special entries
		"DIR":       IconCol{IconFolder, "928374"},
		"gitignore": IconCol{IconGitIgnore, "D65D0E"},
		"git":       IconCol{IconGit, "D65D0E"},
		"ignore":    IconCol{IconIgnore, "D65D0E"},

		"404":  IconCol{IconUnknown, "000000"},
		"TEST": IconCol{IconTest, "D79921"},
	}

	SPECIAL_ICONS ExtMap = ExtMap{
		"md":           IconCol{IconMarkdown, "98971A"},
		"LICENSE":      IconCol{IconLicense, "CC241D"},
		"package.json": IconCol{IconNPM, "CC241D"},
		"Dockerfile":   IconCol{IconDocker, "00CDCD"},
		"Makefile":     IconCol{IconMake, "928374"},
	}
)
