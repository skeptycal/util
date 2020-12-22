// Copyright (c) 2020 Michael Treanor
// MIT License

/*
 Package ansi provides fast ansi escape sequence processing based on strings.Builder. The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
*/
package ansi

import (
	"fmt"
	"strconv"
)

type Ansi int8

func (a Ansi) String() string {
	return fmt.Sprintf("/x1b[%d;", a)
}

func MarshalAnsi(s string) ([]byte, error) {
	return nil, nil
}

// itoa converts the binary value n into an ascii byte slice.  Negative
// values produce an empty slice.
func itoa(n int) []byte {
	if n < 0 {
		return []byte{}
	}
	return []byte(strconv.Itoa(n))
}

const (
	Normal = iota
	Bold   // bold or increased intensity
	Faint  // faint, decreased intensity or second color
	Italics
	Underline
	Blink
	FastBlink
	Inverse
	Hidden
	Strikeout
	PrimaryFont
	AltFont1
	AltFont2
	AltFont3
	AltFont4
	AltFont5
	AltFont6
	AltFont7
	AltFont8
	AltFont9
	Gothic // fraktur
	DoubleUnderline
	NormalColor // normal color or normal intensity (neither bold nor faint)
	NotItalics  // not italicized, not fraktur
	NotUnderlined
	Steady     // not Blink or FastBlink
	Reserved26 // reserved for proportional spacing as specified in CCITT Recommendation T.61
	NotInverse // Positive
	NotHidden  // Revealed
	NotStrikeout
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Reserved38 // intended for setting character foreground color as specified in ISO 8613-6 [CCITT Recommendation T.416]
	Default    // default display color (implementation-defined)
	BlackBackground
	RedBackground
	GreenBackground
	YellowBackground
	BlueBackground
	MagentaBackground
	CyanBackground
	WhiteBackground
	Reserved48        // reserved for future standardization; intended for setting character background color as specified in ISO 8613-6 [CCITT Recommendation T.416]
	DefaultBackground // default background color (implementation-defined)
	Reserved50        // reserved for cancelling the effect of the rendering aspect established by parameter value 26
	Framed
	Encircled
	Overlined
	NotFramed // NotEncircled
	NotOverlined
	Reserved56
	Reserved57
	Reserved58
	Reserved59
	IdeogramUnderline       // ideogram underline or right side line
	IdeogramDoubleUnderline // ideogram double underline or double line on the right side
	IdeogramOverline        // ideogram overline or left side line
	IdeogramDoubleOverline  // ideogram double overline or double line on the left side
	IdeogramStress          // ideogram stress marking
	IdeogramCancel          // cancels the effect of the rendition aspects established by parameter values IdeogramUnderline to IdeogramStress

)

const (
	NormalText       = "\033[0m" // Turn off all attributes
	Reset            = "\033[0m" // alias of NormalText
	BlackText        = "\033[30m"
	RedText          = "\033[31m"
	GreenText        = "\033[32m"
	YellowText       = "\033[33m"
	BlueText         = "\033[34m"
	MagentaText      = "\033[35m"
	CyanText         = "\033[36m"
	WhiteText        = "\033[37m"
	DefaultColorText = "\033[39m" // Normal text color
	BoldText         = "\033[1m"
	BoldBlackText    = "\033[1;30m"
	BoldRedText      = "\033[1;31m"
	BoldGreenText    = "\033[1;32m"
	BoldYellowText   = "\033[1;33m"
	BoldBlueText     = "\033[1;34m"
	BoldMagentaText  = "\033[1;35m"
	BoldCyanText     = "\033[1;36m"
	FaintText        = "\033[2m"
	FaintBlackText   = "\033[2;30m"
	FaintRedText     = "\033[2;31m"
	FaintGreenText   = "\033[2;32m"
	FaintYellowText  = "\033[2;33m"
	FaintBlueText    = "\033[2;34m"
	FaintMagentaText = "\033[2;35m"
	FaintCyanText    = "\033[2;36m"
	FaintWhiteText   = "\033[2;37m"
	DefaultText      = "\033[22;39m" // Normal text color and intensity
)
const (
	GATM = 1  // GUARDED AREA TRANSFER MODE
	KAM  = 2  // KEYBOARD ACTION MODE
	CRM  = 3  // CONTROL REPRESENTATION MODE
	IRM  = 4  // INSERTION REPLACEMENT MODE
	SRTM = 5  // STATUS REPORT TRANSFER MODE
	ERM  = 6  // ERASURE MODE
	VEM  = 7  // LINE EDITING MODE
	BDSM = 8  // BI-DIRECTIONAL SUPPORT MODE
	DCSM = 9  // DEVICE COMPONENT SELECT MODE
	HEM  = 10 // CHARACTER EDITING MODE
	PUM  = 11 // POSITIONING UNIT MODE (see F.4.1 in annex F)
	SRM  = 12 // SEND/RECEIVE MODE
	FEAM = 13 // FORMAT EFFECTOR ACTION MODE
	FETM = 14 // FORMAT EFFECTOR TRANSFER MODE
	MATM = 15 // MULTIPLE AREA TRANSFER MODE
	TTM  = 16 // TRANSFER TERMINATION MODE
	SATM = 17 // SELECTED AREA TRANSFER MODE
	TSM  = 18 // TABULATION STOP MODE
	GRCM = 21 // GRAPHIC RENDITION COMBINATION
	ZDM  = 22 // ZERO DEFAULT MODE (see F.4.2 in annex F)
)
