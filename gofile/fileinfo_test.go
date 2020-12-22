package gofile

import "testing"

func TestExists(t *testing.T) {
	if Exists("fileinfo_test.go") == false {
		t.Fail()
	}
	if Exists("fake_file_not_present.fake") == true {
		t.Fail()
	}
	if !Exists("fileinfo_test.go") == true {
		t.Fail()
	}
	if !Exists("fake_file_not_present.fake") == false {
		t.Fail()
	}
}
