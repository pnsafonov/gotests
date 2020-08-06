package parser

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
 	rgx1 = regexp.MustCompile(`id\s*\:\s*(?P<id>\d+)`)
 	rgxSerStr = regexp.MustCompile(`ser\s*\:\s*(?P<skip>\S+)`)
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

func TryParseSerStr(str string) (result string, ok bool) {
	sm := rgxSerStr.FindStringSubmatch(str)
	if len(sm) != 2 {
		return "", false
	}
	return sm[1], true
}

func TryParseSerSkip(str string) bool {
	result, ok := TryParseSerStr(str)
	if !ok {
		// not found, so don't skip
		return false
	}
	return strings.EqualFold(result, "skip")
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
func TestTryParseSerStr1(t *testing.T) {
	r1, ok1 := TryParseSerStr("ser:    skip")
	if !ok1 || r1 != "skip" {
		t.Fatal("parse failed")
	}

	r2, ok2 := TryParseSerStr("//ser:skip")
	if !ok2 || r2 != "skip" {
		t.Fatal("parse failed")
	}

	r3, ok3 := TryParseSerStr("  //	ser	:	skip	")
	if !ok3 || r3 != "skip" {
		t.Fatal("parse failed")
	}
}

func TestTryParseSerSkip(t *testing.T) {
	r := TryParseSerSkip("ser:    skip")
	if !r {
		t.Fatal("parse fail")
	}
	r = TryParseSerSkip("ser:skip2")
	if r {
		t.Fatal("parse fail")
	}
	r = TryParseSerSkip("  //	ser	:	sKiP	")
	if !r {
		t.Fatal("parse failed")
	}
	r = TryParseSerSkip("  //	ser	:	skip1	")
	if r {
		t.Fatal("parse failed")
	}
}
