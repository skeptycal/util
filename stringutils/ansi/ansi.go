// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
/*
The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"

CSI sequences

For CSI, or "Control Sequence Introducer" commands, the ESC [ is followed by any number (including none) of "parameter bytes" in the range 0x30–0x3F (ASCII 0–9:;<=>?), then by any number of "intermediate bytes" in the range 0x20–0x2F (ASCII space and !"#$%&'()*+,-./), then finally by a single "final byte" in the range 0x40–0x7E (ASCII @A–Z[\]^_`a–z{|}~).[5]:5.4

All common sequences just use the parameters as a series of semicolon-separated numbers such as 1;2;3. Missing numbers are treated as 0 (1;;3 acts like the middle number is 0, and no parameters at all in ESC[m acts like a 0 reset code). Some sequences (such as CUU) treat 0 as 1 in order to make missing parameters useful.[5]:F.4.2

A subset of arrangements was declared "private" so that terminal manufacturers could insert their own sequences without conflicting with the standard. Sequences containing the parameter bytes <=>? or the final bytes 0x70–0x7E (p–z{|}~) are private.

The behavior of the terminal is undefined in the case where a CSI sequence contains any character outside of the range 0x20–0x7E. These illegal characters are either C0 control characters (the range 0–0x1F), DEL (0x7F), or bytes with the high bit set. Possible responses are to ignore the byte, to process it immediately, and furthermore whether to continue with the CSI sequence, to abort it immediately, or to ignore the rest of it.

Reference: https://en.wikipedia.org/wiki/ANSI_escape_code

Terminal output sequences

Code	        Short	    Name	                    Effect
CSI n A	        CUU	        Cursor Up                   Moves the cursor n (default 1) cells in the given
                                                        direction. If the cursor is already at the edge of
                                                        the screen, this has no effect.

CSI n B	        CUD	        Cursor Down
CSI n C	        CUF	        Cursor Forward
CSI n D	        CUB	        Cursor Back
CSI n E	        CNL	        Cursor Next Line	        Moves cursor to beginning of the line n (default 1)
                                                        lines down. (not ANSI.SYS)

CSI n F	        CPL	        Cursor Previous Line	    Moves cursor to beginning of the line n (default 1)
                                                        lines up. (not ANSI.SYS)
CSI n G	        CHA	        Cursor Horizontal Absolute	Moves the cursor to column n (default 1). (not ANSI.SYS)
CSI n ; m H	    CUP	        Cursor Position	            Moves the cursor to row n, column m. The values are
                                                        1-based, and default to 1 (top left corner) if omitted.
                                                        A sequence such as CSI ;5H is a synonym for CSI 1;5H as
                                                        well as CSI 17;H is the same as CSI 17H and CSI 17;1H
CSI n J	        ED	        Erase in Display	        Clears part of the screen. If n is 0 (or missing),
                                                        clear from cursor to end of screen. If n is 1, clear
                                                        from cursor to beginning of the screen. If n is 2,
                                                        clear entire screen (and moves cursor to upper left
                                                        on DOS ANSI.SYS). If n is 3, clear entire screen and
                                                        delete all lines saved in the scrollback buffer (this
                                                        feature was added for xterm and is supported by other
                                                        terminal applications).
CSI n K	        EL	        Erase in Line	            Erases part of the line. If n is 0 (or missing), clear
                                                        from cursor to the end of the line. If n is 1, clear
                                                        from cursor to beginning of the line. If n is 2, clear
                                                        entire line. Cursor position does not change.
CSI n S	        SU	        Scroll Up	                Scroll whole page up by n (default 1) lines. New lines
                                                        are added at the bottom. (not ANSI.SYS)
CSI n T	        SD	        Scroll Down	                Scroll whole page down by n (default 1) lines. New lines
                                                        are added at the top. (not ANSI.SYS)
CSI n ; m f	    HVP	        Horizontal Vertical Pos	    Same as CUP, but counts as a format effector function
                                                        (like CR or LF) rather than an editor function (like CUD
                                                        or CNL). This can lead to different handling in certain
                                                        terminal modes.[5]:Annex A
CSI n m	        SGR	        Select Graphic Rendition	Sets the appearance of the following characters, see
                                                        SGR parameters below.
CSI 5i		                    AUX Port On	            Enable aux serial port usually for local serial printer
CSI 4i		                    AUX Port Off	        Disable aux serial port usually for local serial printer

CSI 6n	        DSR	        Device Status Report	    Reports the cursor position (CPR) to the application as
                                                        (as though typed at the keyboard) ESC[n;mR, where n is
                                                        the row and m is the column.)
*/
package ansi

import (
	"fmt"
	"strconv"
	"strings"
)

type Ansi uint8

// String returns the string representation of an Ansi
// value as a color escape sequence.
func (a Ansi) String() string {
	return fmt.Sprintf("/x1b[%d;", a)
}

// Build returns a string containing multiple ANSI
// color escape sequences.
func (a Ansi) Build(list ...Ansi) string {
	var sb strings.Builder
	defer sb.Reset()

	for _, v := range list {
		sb.WriteString(Ansi(v).String())
	}

	return sb.String()
}

func MarshalAnsi(s string) ([]byte, error) {
	return []byte{}, nil
}

// itoa converts the integer value n into an ascii byte slice.
// Negative values produce an empty slice.
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
