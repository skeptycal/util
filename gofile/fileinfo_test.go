package gofile

import "testing"

func TestExists(t *testing.T) {
	if exists("fileinfo_test.go") == false {
		t.Fail()
	}
	if exists("fake_file_not_present.fake") == true {
		t.Fail()
	}
	if !exists("fileinfo_test.go") == true {
		t.Fail()
	}
	if !exists("fake_file_not_present.fake") == false {
		t.Fail()
	}
}
