package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_input(t *testing.T) {
	tmpfile := generate_tempfile("172")
	old_stdin := os.Stdin
	defer func() { os.Stdin = old_stdin }()
	os.Stdin = &tmpfile

	if _, err := input(); err != nil {
		t.Errorf("User Input Failed: %v", err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	if awb_prefix, _ := input(); awb_prefix != "172" {
		t.Errorf("User Input Failed, expected 172, got  %v", awb_prefix)
	}
}

func Test_prefixInput(t *testing.T) {
	tmpfile := generate_tempfile("172")
	old_stdin := os.Stdin
	defer func() { os.Stdin = old_stdin }()
	os.Stdin = &tmpfile

	if awb_prefix := prefixInput(); awb_prefix != "172" {
		t.Errorf("User Input Failed, expected 172, got  %v", awb_prefix)
	}
}

func Test_generate_awb(t *testing.T) {
	awb := generate_awb("172", 1234567)
	if awb != "172-12345672" {
		t.Error("Wrong AWB-Format", awb)
	}
}

func Test_generate_new_seed(t *testing.T) {
	awb_number := generate_random_seed(10000000)
	// the output in main() is always +1, so the value here must be one lower

	if awb_number != 6993927 {
		t.Error("Problem with the first AWB-Number generation, expected 6993928, received: ", awb_number)
	}
}

func generate_tempfile(file_content string) os.File {
	awb := []byte(file_content)
	tmpfile, err := ioutil.TempFile("", "awbgen")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := tmpfile.Write(awb); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}
	return *tmpfile
}
