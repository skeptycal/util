package ansi

const (
	// Character used for HR function
	HrChar string = "="
)

// Ansi basic color codes
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	RedString      string = "\033[31m"
	GreenString    string = "\033[32m"
	YellowString   string = "\033[33m"
	BlueString     string = "\033[34m"
	PurpleString   string = "\033[35m"
	CyanString     string = "\033[36m"
	WhiteString    string = "\033[37m"
	BgRedString    string = "\033[41m"
	BgGreenString  string = "\033[42m"
	BgYellowString string = "\033[43m"
	BgBlueString   string = "\033[44m"
	BgPurpleString string = "\033[45m"
	BgCyanString   string = "\033[46m"
	BgWhiteString  string = "\033[47m"
)

// Format Strings for Ansi printf commands.
const (
	BasicMask byte = 0xF
	// basic Ansi colors
	FMTansi   string = "\033[%vm"
	FMTansiFG string = "\033[3%vm"
	FMTansiBG string = "\033[4%vm"
	// basic Ansi "bright" colors
	FMTbright string = "\033[1;%vm"
	// basic Ansi "dim" colors
	FMTdim string = "\033[2;%vm"

	FMTansiSet = "\033[%v;%v;%vm"

	// ESC[⟨x⟩8:5:⟨n⟩m Select 8 bit color (x in [ 3, 4 ]) (n in [0..255])
	FMT8bit string = "\033[%v8;5;%vm"
	// ESC[ 38:5:⟨n⟩ m Select foreground color (n in [0..255])
	FMT8bitFG = "\033[38;5;%vm"
	// ESC[ 48:5:⟨n⟩ m Select background color (n in [0..255])
	FMT8bitBG = "\033[48;5;%vm"

	// ESC[⟨x⟩8;2;⟨r⟩;⟨g⟩;⟨b⟩ m     Select RGB color  (x in [ 3, 4 ])
	FMT24bit = "\033[%v8;2;%v;%v;%vm"
	// ESC[ 38;2;⟨r⟩;⟨g⟩;⟨b⟩ m      Select RGB foreground color
	FMT24bitFG = "\033[38;2;%v;%v;%vm"
	// ESC[ 48;2;⟨r⟩;⟨g⟩;⟨b⟩ m      Select RGB background color
	FMT24bitBG = "\033[48;2;%v;%v;%vm"
)

//SGR parameters
/*
SGR (Select Graphic Rendition) sets display attributes. Several attributes can be set in the same sequence, separated by semicolons. Each display attribute remains in effect until a following occurrence of SGR resets it. If no codes are given, CSI m is treated as CSI 0 m (reset / normal).

In ECMA-48 SGR is called "Select Graphic Rendition". In Linux manual pages the term "Set Graphics Rendition" is used.
*/
const (
	Normal byte = iota
	Bold        // bold or increased intensity
	Faint       // faint, decreased intensity or second color
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

// premade common ansi text colors
const (
	BlackText        string = "\033[30m"
	RedText          string = "\033[31m"
	GreenText        string = "\033[32m"
	YellowText       string = "\033[33m"
	BlueText         string = "\033[34m"
	MagentaText      string = "\033[35m"
	CyanText         string = "\033[36m"
	WhiteText        string = "\033[37m"
	DefaultColorText string = "\033[39m" // Normal text color
	BoldText         string = "\033[1m"
	BoldBlackText    string = "\033[1;30m"
	BoldRedText      string = "\033[1;31m"
	BoldGreenText    string = "\033[1;32m"
	BoldYellowText   string = "\033[1;33m"
	BoldBlueText     string = "\033[1;34m"
	BoldMagentaText  string = "\033[1;35m"
	BoldCyanText     string = "\033[1;36m"
	BoldWhiteText    string = "\033[1;37m"
	FaintText        string = "\033[2m"
	FaintBlackText   string = "\033[2;30m"
	FaintRedText     string = "\033[2;31m"
	FaintGreenText   string = "\033[2;32m"
	FaintYellowText  string = "\033[2;33m"
	FaintBlueText    string = "\033[2;34m"
	FaintMagentaText string = "\033[2;35m"
	FaintCyanText    string = "\033[2;36m"
	FaintWhiteText   string = "\033[2;37m"
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

// --------------------------------------------------

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
