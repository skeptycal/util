package anansi

func Write(p []byte) (n int, err error) {
	return 0, nil
}

func Read(p []byte) (n int, err error) {
	return 0, nil
}

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
// func WriteString(w Writer, s string) (n int, err error) {
// 	if sw, ok := w.(StringWriter); ok {
// 		return sw.WriteString(s)
// 	}
// 	return w.Write([]byte(s))
// }
