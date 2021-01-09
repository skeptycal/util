package anansi

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.Debug("anansi logging started")
}

func BR() (n int, err error) {
	return fmt.Println("")
}

func HR() (n int, err error) {
	// todo fix this
	return fmt.Println("")
}

// TableFormat defines cli table formatting options
type TableFormat struct {
	TopHeader    string
	BottomHeader string
	MiddleHeader string
	// SideBorder      string
	// DebugHeader     string
	// AlternateBorder bool
	// AlternateColor  bool
	// HeaderColor     ansiCodes
	// RowColors       []ansiCodes
	// BorderColor     ansiCodes
}

// Tokens defines the ANSI colors used for 'pretty' code output
type Tokens map[string]Attribute

// New returns a new Tokens object
func (t *Tokens) New(m map[string]Attribute) Tokens {
	var defaultTokens = make(Tokens)

	defaultTokens = Tokens{
		"colorKeywords":    FgCyan,
		"colorOperator":    FgRed,
		"colorIntegers":    FgMagenta,
		"colorFloats":      FgMagenta,
		"colorStrings":     FgHiGreen,
		"colorComments":    FgGreen,
		"colorPunctuation": FgHiBlue,
		"colorTypes":       FgYellow,
	}
	if m == nil {
		return defaultTokens
	}

	return m
}

// CliConfig defines the configuration for CLI output
type CliConfig struct {
	ScreenWidth int
	UseColor    bool
	TableFormat
	Tokens
}

// New returns a new CliConfig object
// todo finish this
// func (c *CliConfig) New() *CliConfig {
// 	return &CliConfig{
// 		ScreenWidth: 79,
// 		UseColor:    true,
// 		TableFormat: TableFormat.New(),
// 		Tokens:      Tokens.New(),
// 	}
// }

// VerboseLevel defines the level of visual feedback in the cli terminal
type VerboseLevel uint

const (
	verboseQuiet VerboseLevel = iota
	verboseMinimal
	verboseStandard
	verboseVerbose
	verboseDebug
	verboseAll
)

// StartTimer returns a function that will return the elapsed time
func StartTimer(name string) func() {
	/*
		    func example() {
			   stop := StartTimer("myTimer")
			   defer stop()

			   time.Sleep(1*time.Second)
	*/

	t := time.Now()
	log.Info("Timer started:", name)
	return func() {
		d := time.Now().Sub(t)
		log.Info(name, "took", d)
	}
}
