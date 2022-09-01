package main

type Token struct {
	loc    int
	rawStr string
	length int
}

type Tokenizer struct {
	lines        string
	isInStr      bool
	strStartChar byte
	offset       int
	ptr          int
	tokens       []Token
}

func NewTokenizer(lines string) *Tokenizer {
	return &Tokenizer{
		lines:        lines,
		isInStr:      false,
		strStartChar: 0,
		offset:       0,
		ptr:          0,
		tokens:       make([]Token, 0, 500),
	}
}

func (t *Tokenizer) tok() Token {
	return Token{
		loc:    t.offset,
		length: t.ptr - t.offset - 1,
		rawStr: t.lines[t.offset : t.ptr-1],
	}
}

func (t *Tokenizer) add() {
	t.tokens = append(t.tokens, t.tok())
	t.offset = t.ptr
	t.ptr++
}

func (t *Tokenizer) skip() {
	t.ptr++
  t.offset = t.ptr
}

func (t *Tokenizer) next() (byte, bool) {
	if len(t.lines) <= t.ptr {
		return 0, false
	}

	b := t.lines[t.ptr]
	t.ptr++
	return b, true
}

func (t *Tokenizer) enterStr(startChar byte) {
	t.isInStr = true
	t.strStartChar = startChar
}

func (t *Tokenizer) exitStr() {
	t.isInStr = false
	t.strStartChar = 0
}

func (t *Tokenizer) Tokenize() []Token {
	for {
		b, ok := t.next()
		if !ok {
			break
		}

		if t.isInStr {
			if b == t.strStartChar {
				t.add()
			}
			continue
		}

		switch b {
		case '\'', '"':
			t.enterStr(b)
		case ' ':
			if t.lines[t.offset] == ' ' {
				t.skip()
			} else {
				t.add()
			}
		case ';':
			t.add()
		}
	}

	return t.tokens
}
