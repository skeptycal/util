// readyset - a CLI framework that implements traditional command line tools in a more efficient way.
//
// Notes: Developer Alpha
// Feature suggestions
//
//https://clig.dev/ (https://github.com/cli-guidelines/cli-guidelines) written by SquareSpace devs
// CLIG - Command Line Interface Guidelines
// An open-source guide to help you write better command-line programs, taking traditional UNIX principles and updating them for the modern day.
//
// have a full featured configuration management system to add, list, and store options, tokens, usernames, option sets, and configuration profiles.
/* Excerpts from clig.dev
GUI design, particularly in its early days, made heavy use of metaphor: desktops, files, folders, recycle bins. It made a lot of sense, because computers were still trying to bootstrap themselves into legitimacy. The ease of implementation of metaphors was one of the huge advantages GUIs wielded over CLIs. Ironically, though, the CLI has embodied an accidental metaphor all along: it’s a conversation.

Acknowledging the conversational nature of command-line interaction means you can bring relevant techniques to bear on its design. You can suggest possible corrections when user input is invalid, you can make the intermediate state clear when the user is going through a multi-step process, you can confirm for them that everything looks good before they do something scary.

(check out https://www.nngroup.com/articles/anti-mac-interface/)

Generate feature conversations between users  and the system by retrieving information from the system history or an API. This allows the generation of lists of commands based on what the user is trying to accomplish.

e.g.
Traditional:
ls -R *.py
"doesn't work"
ls --help
man ls
ls -X sorts by extension but does not filter or search ...
google stackoverflow "macOS ls files containing extension"
maybe find is the only way?
several posts and technical arguments later ...
find . -regex '.* /Robert\.\(h\|cpp\)$'
find . -name '*.jpg' -o -name '*.png' -print

find . -name '*.py' -print

one better way?
search terms in single quotes mean 'literal' search terms  (this is a common UNIX and POSIX tradition)
ls '*.py'
generates list ...

Does this really work? Do I really have 33 gig of pure python files?
du -chs *.py
33G	total

try ...
find . -name '*.py' -print | du -chs

space taken up by all python files in this directory (recursively)
find . -name '.venv' -type d | du -cks

lines of code in all of the python files ...
find . -name '*.py' -print | wc -l
33G	total

names of all python virtual environments named .venv (my personal preference)
find . -name '.venv' -type d

ducks in a row!
space occupied by all python virtual environments named .venv (my personal preference)  (run the looong find query again )
find . -name '.venv' -type d | du -cks # <-- the ducks ...
35 923 656K	total # that's a lot! I need to do this more often!

delete them (run the looong find query again )
find . -name '.venv' -type d -exec rm -rf {} +

add this to a cron job to run weekly ...
crontab -l # find the right file for my weekly non-admin jobs
...
@weekly cd ${HOME}/.dotfiles/easycron && ./cron_sunday.sh
...

cd ${HOME}/.dotfiles/easycron && code ./cron_sunday.sh
copy/paste that previous command that worked ...


how about .git repos that I don't really need? Many of these directories are
just things I use as reference and different versions of dependencies
find . -name '.git' -type d -exec du -chs {} +
23G	total

what about adding a keep tag file to folders I definitely want to keep the local .git folder for?
then ignore the .git folders that contain keep (a blank file named keep with no access permissions)
??

time it takes to search ???
# using locate database
time locate "*.py" >/dev/null
locate "*.py" > /dev/null  2.46s user 0.02s system 99% cpu 2.485 total

# using find
time find . -name "*.py" -type f >/dev/null
find . -name "*.py" -type f > /dev/null  2.14s user 21.75s system 71% cpu 33.612 total

# using golang
gofind *.py

*/

package readyset

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"golang.org/x/text/language"
	_ "golang.org/x/text/language"
)

var (
	Lang language.Tag
)

type readySettings struct {
	// - GNU compatible flags
	//
	// - Import flags
	//
	// - Create flags (something like modeling with clay ...)
	//
	// - String based flags specification
	//
	gnuflags string

	// ansi colors
	//
	// - strings.Buffer
	//
	// - automatic replacement of "color tags"
	//
	// - what is the best template tag??
	//
	// - convert to / from all common color formats
	//
	colors string

	// strings
	//
	// - parsers
	//
	// - markdown
	//
	// - pretty formatters
	//
	// - translators (python to go would be fun)
	//
	strings string

	// language
	//
	// todo - use golang.org/x/text ??
	// quote from source
	//
	// text is a repository of text-related packages related to internationalization
	// (i18n) and localization (l10n), such as character encodings, text
	// transformations, and locale-specific text handling.
	//
	// There is a 30 minute video, recorded on 2017-11-30, on the "State of
	// golang.org/x/text" at https://www.youtube.com/watch?v=uYrDrMEGu58
	//
	// reference: https://phrase.com/blog/posts/internationalization-i18n-go/
	//
	language string

	// data is configuration information for the data package.
	//
	// mysql database connection information
	//
	// - parser
	//
	// - inferStructure
	//
	// - cleanDataSet
	//
	// - translate
	//
	// - io.Reader implementations
	//
	// - io.Writer implementations
	//
	data string

	// math is the configuration information for the math package
	//
	// - data analysis
	//
	// - statistics
	//
	// - boolean algebra
	//
	// - linear algebra
	//
	// - vector analysis
	//
	// - predictions
	//
	// - solver

	// AI is the configuration information for the AI package
	//
	// - pattern recognition
	//
	// - prediction engine
	//
	// - identifier
	//
	// - ignore engine (maximize ignoring to reduce scope and increase precision)
	//
	ai string

	// regex is the configuration information, constants and patterns
	//
	// - collection of standard compiled regex patterns
	//
	// - simple regex builder (like lego bricks)
	//
	regex string

	// io - stdout, stderr, stdin functionality
	//
	// - colorized output
	//
	// - tee output
	//
	// - multiwriter
	//
	// - jsonwriter
	//
	// - toml, yaml, CSV
	//
	// - SQL writer
	//
	// - ghost logger
	//
	io string

	// logger is the configuration information for the logger package
	//
	// using mainly logrus
	//
	// - config file
	//
	// - io.writer
	//
	// - io.reader
	//
	// - regexLogReader("xxxxxx")
	//
	// - logFind("xxxxxx")
	//
	// - duplicate to CSV, JSON, SQL, etc.
	//
	logger string

	// shell command wrapper configuration
	//
	shell string

	// repl is a shell REPL written in Go
	//
	repl string
}

func NewReadySettings() (*readySettings, error) {
	s := &readySettings{}
	if s == nil {
		return nil, errors.New("readySettings could not be initialized")
	}
	return s, nil
}

func init() {
	RSG, err := NewReadySettings()
	if err != nil {
		log.Info(err)
	}
	Lang, err = language.Parse("en-US")
	if err != nil {
		log.Info(err)
	}
	log.Info("readySettings are initialized: %v", RSG)
}
