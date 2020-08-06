package parser

import (
	"regexp"
	"strconv"
	"testing"
)

var (
 	rgx1 = regexp.MustCompile(`id\s*\:\s*(?P<id>\d+)`)
)

func TryParseId(str string) (result int, ok bool) {
	sm := rgx1.FindStringSubmatch(str)
	if len(sm) != 2 {
		return 0, false
	}

	v, err := strconv.Atoi(sm[1])
	if err != nil {
		return 0, false
	}

	return v, true
}

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

func TestRegex2(t *testing.T) {
	r1, ok1 := TryParseId("//id:768")
	if !ok1 || r1 != 768 {
		t.Fatal("parse failed")
	}

	r2, ok2 := TryParseId("//id :  	33  ")
	if !ok2 || r2 != 33 {
		t.Fatal("parse failed")
	}

	r3, ok3 := TryParseId("//		   id    : 768dsads")
	if !ok3 || r3 != 768 {
		t.Fatal("parse failed")
	}

	r4, ok4 := TryParseId("/*		   id    : 8  */")
	if !ok4 || r4 != 8 {
		t.Fatal("parse failed")
	}
}