package ansi

const (
	FMTansiSet = "\033[%v;%v;%vm"
)

// Format Strings 3/4 bit Ansi printf commands.
//
// ESC[⟨x⟩⟨n⟩m
//
// Select 3/4 bit color
//
// n in [0..8]; basic colors
//
// (x in [3, 4]); 3 = foreground; 4 = background
//
// add an extra effect:
//
// ESC[⟨e⟩;⟨x⟩⟨n⟩m
//
// e can be any valid ANSI escape effect; these are common:
// 	0 Normal
//  1 Bold              // bold or increased intensity
//  2 Faint             // faint, decreased intensity or second color
//  3 Italics
//  4 Underline
//  (5) Blink           // not widely supported
//  (6)FastBlink        // not widely supported
//  7 Inverse
//  8 Conceal
//  9 Strikeout
//
const (
	// basic Ansi colors
	FMTansi   string = "\033[%vm"
	FMTansiFG string = "\033[3%vm"
	FMTansiBG string = "\033[4%vm"
	// basic Ansi "bright" colors
	FMTbright string = "\033[1;%vm"
	// basic Ansi "dim" colors
	FMTdim string = "\033[2;%vm"
)

// Format Strings 8 bit Ansi printf commands.
//
// ESC[⟨x⟩8:5:⟨n⟩m
//
// Select 8bit color
//
// n in [0..255]; 0-231 are colors; 232-255 are grayscale
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	FMT8bit   string = "\033[%v8;5;%vm"
	FMT8bitFG string = "\033[38;5;%vm"
	FMT8bitBG string = "\033[48;5;%vm"
)

// Format Strings 24 bit Ansi printf commands.
//
// ESC[⟨x⟩8;2;⟨R⟩;⟨G⟩;⟨B⟩m
//
// Select RGB color
//
// R, G, B in [0..255]
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	FMT24bit   string = "\033[%v8;2;%v;%v;%vm"
	FMT24bitFG string = "\033[38;2;%v;%v;%vm"
	FMT24bitBG string = "\033[48;2;%v;%v;%vm"
)
