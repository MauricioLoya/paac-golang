package utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Mauricio"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello(name)

	if !want.MatchString(msg) || error != nil {
		t.Errorf("Hello(%q) = %q; want match %q", name, msg, want)
	}

}

func TestHelloEmptyName(t *testing.T) {

	msg, error := Hello("")

	fmt.Println(msg)

	if msg != "" && error != nil {
		t.Errorf("Hello(%q) = %q;", "", msg)
	}

}
