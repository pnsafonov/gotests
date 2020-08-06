package parser

import (
	"regexp"
	"testing"
)

var (
 	rgx1 = regexp.MustCompile(`id\s*\:\s*(?P<id>\d+)`)
)


func TestRegex1(t *testing.T) {
	names := rgx1.SubexpNames()
	if len(names) == 0 {
		t.Fatal("len is nil")
	}

	sm := rgx1.FindStringSubmatch("//id:768")
	if len(sm) == 0 {
		t.Fatal("len is nil")
	}
}