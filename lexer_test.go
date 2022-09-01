package main

import "testing"

func assertEql(t1, t2 Token) bool {
	return t1.loc == t2.loc && t1.length == t2.length && t1.rawStr == t2.rawStr
}

func TestTokenize(t *testing.T) {
	l1 := "select * from tbl1 where id = 'test';"
	expected := []Token{
		{
			loc:    0,
			length: 6,
			rawStr: "select",
		}, {
			loc:    7,
			length: 1,
			rawStr: "*",
		}, {
			loc:    9,
			length: 4,
			rawStr: "from",
		}, {
			loc:    14,
			length: 4,
			rawStr: "tbl1",
		}, {
			loc:    19,
			length: 5,
			rawStr: "where",
		}, {
			loc:    25,
			length: 2,
			rawStr: "id",
		}, {
			loc:    28,
			length: 1,
			rawStr: "=",
		}, {
			loc:    30,
			length: 4,
			rawStr: "test",
		},
	}

	tokenizer := NewTokenizer(l1)
	for i, tok := range tokenizer.Tokenize() {
		if !assertEql(tok, expected[i]) {
			t.Errorf("expected %+v, but given %+v\n", expected[i], tok)
		}
	}
}
