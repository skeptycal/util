package gofile

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestTruncateFile(t *testing.T) {
	f, err := os.Create("test_truncate_file")
	defer os.Remove(f.Name())
	if err != nil {
		t.Errorf("error creating tmpfile: %v", err)
	}

	n, err := f.WriteString("fake_data")
	if err != nil || n != 9 {
		t.Errorf("error writing to file %s: %v", f.Name(), err)
	}

	truncfile, err := TruncateFile(f.Name())
	if err != nil {
		t.Errorf("error opening file %s: %v", "test_truncate_file", err)
	}

	w := io.WriteCloser(truncfile)

	buf := []byte("fake data")

	n, err = w.Write(buf)
	if err != nil || n != len(buf) {
		t.Errorf("error writing to file %s: %v", truncfile.Name(), err)
	}

	errReader := io.ReadCloser(truncfile)

	_, err = errReader.Read(buf)
	if err == nil {
		t.Errorf("should not be able to read from file %s: %v", truncfile.Name(), err)
	}
	errReader.Close()
	f.Close()

	f, err = os.Open(truncfile.Name())
	if err != nil {
		t.Errorf("error opening file %s to read: %v", f.Name(), err)
	}

	r := io.ReadCloser(f)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("error in reading check from file %s: %v", f.Name(), err)
	}

	if string(b) != "fake data" {
		t.Errorf("error in file data want %q got %q", "fake data", string(b))
	}

}

func TestWriteFile(t *testing.T) {

	errorfile := "/dev/rdisk0"
	_, err := TruncateFile(errorfile)
	if err == nil {
		t.Errorf("file should cause permission error %s: %v", errorfile, err)
	}

	utf8TestFile := "testdata/utf8TestFile"
	data := "€"

	err = WriteFile(utf8TestFile, data)
	if err == nil {
		t.Errorf("file should cause 'incorrect bytes written' error in %s with data '%s': %v", utf8TestFile, data, err)

	}

	filename := "testdata/TestWriteFile"
	data = "fake data"

	err = WriteFile(filename, data)
	if err != nil {
		t.Errorf("error writing data to file %s: %v", filename, err)
	}
	defer os.Remove(filename)

	r, err := os.Open(filename)
	if err != nil {
		t.Errorf("error opening file %s: %v", filename, err)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("error reading file %s: %v", filename, err)
	}

	if !bytes.Equal(b, []byte(data)) {
		t.Errorf("file data did not match want %q got %q", data, b)
	}
}
