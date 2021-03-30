package stringutils

import (
	"regexp"
	"unicode/utf8"
)

// Numbers fundamental to the encoding.
const (
	RuneError     = utf8.RuneError // '\uFFFD'       // the "error" Rune or "Unicode replacement character"
	RuneSelf      = utf8.RuneSelf  // 0x80           // characters below RuneSelf are represented as themselves in a single byte.
	MaxRune       = utf8.MaxRune   // '\U0010FFFF'   // Maximum valid Unicode code point.
	UTFMax        = utf8.UTFMax    // 4              // maximum number of bytes of a UTF-8 encoded Unicode character.
	alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz"
)

var (
	// note: 0x0B is not considered whitespace by regex so an alternative
	// solution must be considered if 0x0B detection is required.
	//
	// regex is the slowest solution (by far!) so this is likely moot.
	reWhitespace = regexp.MustCompile(`[\s\v]+`)

	// leaves out 0x85, 0xA0
	shortASCIIList = []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20}

	// all whitespace code points that are one byte long
	shortByteList = []byte{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0}

	// most common unicode whitespace code points
	longRuneList = []rune{0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x20, 0x85, 0xA0, 0x2000, 0x200A, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFFEF, 0x1680}

	// byte lists transformed to strings
	shortASCIIListString = string(shortASCIIList)
	shortByteListString  = string(shortByteList)
	longRuneListString   = string(longRuneList)

	// UnicodeWhiteSpaceMap provides a mapping from Unicode runes to strings
	// with descriptions of each. It is marginally slower than the bool map.
	//
	// In computer programming, whitespace is any character or series of
	// characters that represent horizontal or vertical space in typography.
	// When rendered, a whitespace character does not correspond to a visible
	// mark, but typically does occupy an area on a page. For example, the
	// common whitespace symbol SPACE (unicode: U+0020 ASCII: 32 decimal 0x20
	// hex) represents a blank space punctuation character in text, used as a
	// word divider in Western scripts.
	//
	// Reference: https://en.wikipedia.org/wiki/Whitespace_character
	UnicodeWhiteSpaceMap = map[rune]string{
		0x0009: `CHARACTER TABULATION <TAB>`,
		0x000A: `ASCII LF`,
		0x000B: `LINE TABULATION <VT>`,
		0x000C: `FORM FEED <FF>`,
		0x000D: `ASCII CR`,
		0x0020: `SPACE <SP>`,
		0x00A0: `NO-BREAK SPACE <NBSP>`,
		0x0085: `NEL; Next Line`,
		0x1680: `Ogham space mark, interword separation in Ogham text`,
		0x2000: `EN QUAD, 0x2002 is preferred`,
		0x2001: `EM QUAD, mutton quad, 0x2003 is preferred`,
		0x2002: `EN SPACE, "nut", &ensp, LaTeX: '\enspace'`,
		0x2003: `EM SPACE, "mutton", &emsp;, LaTeX: '\quad'`,
		0x2004: `THREE-PER-EM SPACE, "thick space", &emsp13;`,
		0x2005: `four-per-em space, "mid space", &emsp14;`,
		0x2006: `SIX-PER-EM SPACE, sometimes equated to U+2009`,
		0x2007: `FIGURE SPACE, width of monospaced char, &numsp;`,
		0x2008: `PUNCTUATION SPACE, width of period or comma, &puncsp;`,
		0x2009: `THIN SPACE, 1/5th em, thousands sep, &thinsp;; LaTeX: '\,'`,
		0x200A: `HAIR SPACE, &hairsp;`,
		0x2028: `LINE SEPARATOR`,
		0x2029: `PARAGRAPH SEPARATOR`,
		0x202F: `NARROW NO-BREAK SPACE`,
		0x205F: `MEDIUM MATHEMATICAL SPACE, MMSP, &MediumSpace, 4/18 em`,
		0x3000: `IDEOGRAPHIC SPACE, full width CJK character cell`,
		0xFFEF: `ZERO WIDTH NO-BREAK SPACE <ZWNBSP> (BOM), deprecated Unicode 3.2 (use U+2060)`,
		// Related Unicode characters property White_Space=no
		// 0x180E: `MONGOLIAN VOWEL SEPARATOR, not whitespace as of Unicode 6.3.0`,
		// 0x200B: `ZERO WIDTH SPACE, ZWSP, "soft hyphen", &ZeroWidthSpace;`,
		// 0x200C: `ZERO WIDTH NON-JOINER, ZWNJ, &zwnj;`,
		// 0x200D: `ZERO WIDTH JOINER, ZWJ, &zwj;`,
		// 0x2060: `WORD JOINER, WJ, not a line break, &NoBreak;`,
	}

	// UnicodeWhiteSpaceMap provides a mapping from Unicode runes to bool{true}.
	// In computer programming, whitespace is any character or series of
	// characters that represent horizontal or vertical space in typography.
	// When rendered, a whitespace character does not correspond to a visible
	// mark, but typically does occupy an area on a page. For example, the
	// common whitespace symbol U+0020 SPACE (also ASCII 32) represents a
	// blank space punctuation character in text, used as a word divider in
	// Western scripts.
	unicodeWhiteSpaceMapBool = map[rune]bool{
		0x0009: true, // CHARACTER TABULATION <TAB>
		0x000A: true, // ASCII LF
		0x000B: true, // LINE TABULATION <VT>
		0x000C: true, // FORM FEED <FF>
		0x000D: true, // ASCII CR
		0x0020: true, // SPACE <SP>
		// > utf8.RuneSelf
		0x00A0: true, // NO-BREAK SPACE <NBSP>
		0x0085: true, // NEL; Next Line
		// > unicode.MaxLatin1
		0x1680: true, // Ogham space mark, interword separation in Ogham text
		0x2000: true, // EN QUAD, 0x2002 is preferred
		0x2001: true, // EM QUAD, mutton quad, 0x2003 is preferred
		0x2002: true, // EN SPACE, "nut", &ensp, LaTeX: '\enspace'
		0x2003: true, // EM SPACE, "mutton", &emsp;, LaTeX: '\quad'
		0x2004: true, // THREE-PER-EM SPACE, "thick space", &emsp13;
		0x2005: true, // four-per-em space, "mid space", &emsp14;
		0x2006: true, // SIX-PER-EM SPACE, sometimes equated to U+2009
		0x2007: true, // FIGURE SPACE, width of monospaced char, &numsp;
		0x2008: true, // PUNCTUATION SPACE, width of period or comma, &puncsp;
		0x2009: true, // THIN SPACE, 1/5th em, thousands sep, &thinsp;; LaTeX: '\,'
		0x200A: true, // HAIR SPACE, &hairsp;
		0x2028: true, // LINE SEPARATOR
		0x2029: true, // PARAGRAPH SEPARATOR
		0x202F: true, // NARROW NO-BREAK SPACE
		0x205F: true, // MEDIUM MATHEMATICAL SPACE, MMSP, &MediumSpace, 4/18 em
		0x3000: true, // IDEOGRAPHIC SPACE, full width CJK character cell
		0xFFEF: true, // ZERO WIDTH NO-BREAK SPACE <ZWNBSP> (BOM), deprecated Unicode 3.2 (use U+2060)
		// Related Unicode characters property White_Space=no
		// 0x180E: true, // MONGOLIAN VOWEL SEPARATOR, not whitespace as of Unicode 6.3.0
		// 0x200B: true, // ZERO WIDTH SPACE, ZWSP, "soft hyphen", &ZeroWidthSpace;
		// 0x200C: true, // ZERO WIDTH NON-JOINER, ZWNJ, &zwnj;
		// 0x200D: true, // ZERO WIDTH JOINER, ZWJ, &zwj;
		// 0x2060: true, // WORD JOINER, WJ, not a line break, &NoBreak;
	}
)
