// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CSI sequences
//
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	// For CSI, or "Control Sequence Introducer" commands, the ESC [ is followed by any number (including none) of
	// "parameter bytes" in the range 0x30–0x3F (ASCII 0–9:;<=>?)
	ParameterBytes = "0x30–0x3F (ASCII 0–9:;<=>?)"

	// then by any number of
	// "intermediate bytes" in the range 0x20–0x2F (ASCII space and !"#$%&'()*+,-./)
	IntermediateBytes = `0x20–0x2F (ASCII space and !"#$%&'()*+,-./)`

	// then finally by a single "final byte" in the range 0x40–0x7E (ASCII @A–Z[\]^_`a–z{|}~)
	FinalBytes = "0x40–0x7E (ASCII @A–Z[\\]^_`a–z{|}~)"

	Delimiter = ";"

	// LegalRange
	/*
	   All common sequences just use the parameters as a series of semicolon-separated numbers such as 1;2;3. Missing numbers are treated as 0 (1;;3 acts like the middle number is 0, and no parameters at all in ESC[m acts like a 0 reset code). Some sequences (such as CUU) treat 0 as 1 in order to make missing parameters useful.

	   A subset of arrangements was declared "private" so that terminal manufacturers could insert their own sequences without conflicting with the standard. Sequences containing the parameter bytes <=>? or the final bytes 0x70–0x7E (p–z{|}~) are private.
	*/
	LegalRange = "0x20-0x7E"

	// IllegalRange
	/*
	   The behavior of the terminal is undefined in the case where a CSI sequence contains any character outside of the range 0x20–0x7E. These illegal characters are either C0 control characters (the range 0–0x1F), DEL (0x7F), or bytes with the high bit set. Possible responses are to ignore the byte, to process it immediately, and furthermore whether to continue with the CSI sequence, to abort it immediately, or to ignore the rest of it.*/
	IllegalRange = "0-0x1F,0x7F,0x80-0xFF"

	// Terminal output sequences
	/*
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
									   	                                                           A sequence such as CSI ;5H is a  synonym for CSI 1;5H as
		                                                                                              well as CSI 17;H is the same as CSI 17H and CSI 17;1H
	*/

	// CSI n J
	// ED - Erase in Display Clears part of the screen. If n is 0 (or missing), clear from cursor to end of screen. If n is 1, clear from cursor to beginning of the screen. If n is 2, clear entire screen (and moves cursor to upper left on DOS ANSI.SYS). If n is 3, clear entire screen and delete all lines saved in the scrollback buffer (this feature was added for xterm and is supported by other terminal applications).*/
	CSIclear = "CSI 2 J"

	// CSI functions
	/*
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

)

// --------------------------------------------------

const (
	FMT256 string = "\033[38;5;%vm"
	HrChar string = "="
)

// Ansi 7-bit color codes
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	ansi7fmt string = "\033[%vm"
	Reset    string = "\033[0m"

	Red    string = "\033[31m"
	Green  string = "\033[32m"
	Yellow string = "\033[33m"
	Blue   string = "\033[34m"
	Purple string = "\033[35m"
	Cyan   string = "\033[36m"
	White  string = "\033[37m"

	BgRed    string = "\033[41m"
	BgGreen  string = "\033[42m"
	BgYellow string = "\033[43m"
	BgBlue   string = "\033[44m"
	BgPurple string = "\033[45m"
	BgCyan   string = "\033[46m"
	BgWhite  string = "\033[47m"
)

var (
	ansi            ANSI      = NewANSIWriter(33, 44, 1)
	AnsiFmt         string    = ansi.Build(1, 33, 44)
	AnsiReset       string    = ansi.Build(0, 39, 49)
	defaultioWriter io.Writer = os.Stdout
)

// todo - create a pool of stringbuilders that can go when ready?
// type sbSync struct {
// 	strings.Builder
// 	mu sync.Mutex
// }

// NewANSIWriter returns a new ANSI Writer for use in terminal output.
// If w is nil, the default (os.Stdout) is used.
func NewANSIWriter(fg, bg, ef byte, w io.Writer) ANSI {
	if wr, ok := w.(io.Writer); !ok || w == nil {
		w = defaultioWriter
	}

	return &Ansi{
		fg: ansiFormat(fg),
		bg: ansiFormat(bg),
		ef: ansiFormat(ef),
		bufio.NewWriter(w),
		// sb: strings.Builder{}
	}
}

type ANSI interface {
	io.Writer
	io.StringWriter
	Build(b ...byte) string
}

type Ansi struct {
	bufio.Writer
	fg string
	bg string
	ef string
	// sb *strings.Builder
}

// Build encodes a variadic list of bytes into ANSI 7 bit escape codes.
func (a *Ansi) Build(b ...byte) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for _, n := range b {
		sb.WriteString(fmt.Sprintf(ansi7fmt, n))
	}
	return sb.String()
}

// Set accepts, encodes, and prints a variadic argument list of bytes
// that represent ANSI colors.
func (a *Ansi) Set(b ...byte) (int, error) {
	return fmt.Fprint(os.Stdout, a.Build(b...))
}

func hr(n int) {
	fmt.Println(strings.Repeat(hrChar, n))
}

func br() {
	fmt.Println("")
}

// func ansiFormat(n byte) string {
// 	return fmt.Sprintf(ansi7fmt, n)
// }
// func aPrint(a ...byte) {
// 	fmt.Print(ansi.Build(a...))
// }

// Echo is a helper function that wraps printing to stdout
// in Ansi color escape sequences.
//
// If the first argument is is a string that contains a %
// character, it is used as a format string for fmt.Printf,
// otherwise fmt.Println is used for all arguments.
//
// AnsiFmt is the current text color.
//
// AnsiReset is the Ansi reset code.
//
func Echo(a ...interface{}) {
	fmt.Print(AnsiFmt)

	if fs, ok := a[0].(string); ok {
		if strings.Contains(fs, "%") {
			fmt.Printf(fs, a[1:])
		} else {
			fmt.Println(a...)
		}
	}
	fmt.Print(AnsiReset)
}

// --------------------------------------------------
type AnsiOld uint8

// String returns the string representation of an Ansi
// value as a color escape sequence.
func (a AnsiOld) String() string {
	return fmt.Sprintf("/x1b[%d;", a)
}

// Build returns a string containing multiple ANSI
// color escape sequences.
func (a AnsiOld) Build(list ...Ansi) string {
	var sb strings.Builder
	defer sb.Reset()

	for _, v := range list {
		sb.WriteString(Ansi(v).String())
	}

	return sb.String()
}

// itoa converts the integer value n into an ascii byte slice.
// Negative values produce an empty slice.
func itoa(n int) []byte {
	if n < 0 {
		return []byte{}
	}
	return []byte(strconv.Itoa(n))
}

/*
SGR parameters

SGR (Select Graphic Rendition) sets display attributes. Several attributes can be set in the same sequence, separated by semicolons. Each display attribute remains in effect until a following occurrence of SGR resets it. If no codes are given, CSI m is treated as CSI 0 m (reset / normal).

In ECMA-48 SGR is called "Select Graphic Rendition". In Linux manual pages the term "Set Graphics Rendition" is used.
*/
const (
	Normal = iota
	Bold   // bold or increased intensity
	Faint  // faint, decreased intensity or second color
	Italics
	Underline
	Blink
	FastBlink
	Inverse
	Conceal
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
	SetForeground     // Next arguments are 5;n or 2;r;g;b, see below
	DefaultForeground // default display color (implementation-defined)
	BlackBackground
	RedBackground
	GreenBackground
	YellowBackground
	BlueBackground
	MagentaBackground
	CyanBackground
	WhiteBackground
	SetBackground              // Next arguments are 5;n or 2;r;g;b, see below
	DefaultBackground          // default background color (implementation-defined)
	DisableProportionalSpacing // reserved for cancelling the effect of parameter value 26
	Framed
	Encircled
	Overlined
	NotFramed // NotEncircled
	NotOverlined
	Reserved56
	Reserved57
	SetUnderlineColor // Next arguments are 5;n or 2;r;g;b, see below
	DefaultUnderlineColor
	IdeogramUnderline       // ideogram underline or right side line
	IdeogramDoubleUnderline // ideogram double underline or double line on the right side
	IdeogramOverline        // ideogram overline or left side line
	IdeogramDoubleOverline  // ideogram double overline or double line on the left side
	IdeogramStress          // ideogram stress marking
	IdeogramCancel          // reset the effects of all of 60–64
	Superscript             = 73
	Subscript               = 74
)

const (
	DefaultColors    = "\033[39;49m"
	DefaultText      = "\033[22;39m" // Normal text color and intensity
	NormalText       = "\033[0m"     // Turn off all attributes
	Reset            = "\033[0m"     // alias of NormalText
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
)

/*
8-bit
As 256-color lookup tables became common on graphic cards, escape sequences were added to select from a pre-defined set of 256 colors:[citation needed]

ESC[ 38;5;⟨n⟩ m Select foreground color
ESC[ 48;5;⟨n⟩ m Select background color
  0-  7:  standard colors (as in ESC [ 30–37 m)
  8- 15:  high intensity colors (as in ESC [ 90–97 m)
 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
232-255:  grayscale from black to white in 24 steps
*/
const (
	// ESC[ 38:5:⟨n⟩ m Select foreground color
	ansi8bitFGfmt = "\033[38:5:%vm;"
	// ESC[ 48:5:⟨n⟩ m Select background color
	ansi8bitBGfmt = "\033[48:5:%vm;"
)

const (
	// ESC[ 38;2;⟨r⟩;⟨g⟩;⟨b⟩ m Select RGB foreground color
	ansi24bitFGfmt = "\033[38;2;%v;%v;%vm;"
	// ESC[ 48;2;⟨r⟩;⟨g⟩;⟨b⟩ m Select RGB background color
	ansi24bitBGfmt = "\033[48;2;%v;%v;%vm;"
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
