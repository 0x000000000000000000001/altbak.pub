package purescript

// Experimental typed cursor: no cursor is boxed into any or published as Json.
// Final strings are owned; the validated index and input die with this call.
import (
	"gopurs/output/gopurs_runtime"
	"strconv"
)

type tcCursor directCursor
type tcArray struct {
	cursor tcCursor
	count  int
}

func (c tcCursor) kind() byte {
	if c.document == nil {
		return 0
	}
	return directCursor(c).kind()
}
func (c tcCursor) text() (string, bool) {
	if c.kind() != '"' {
		return "", false
	}
	return directCursor(c).string(true), true
}

// Only statically selected discriminator comparisons use this view. All names,
// literals, paths and other strings stored in the final AST still use text().
func (c tcCursor) borrowedText() (string, bool) {
	if c.kind() != '"' {
		return "", false
	}
	return directCursor(c).string(false), true
}

func tcBorrowedTag(c tcCursor) gopurs_runtime.Value {
	value, ok := c.borrowedText()
	if !ok {
		tc_cndFail("String")
	}
	return gopurs_runtime.Str(value)
}
func (c tcCursor) number() (float64, bool) {
	k := c.kind()
	if k != '-' && (k < '0' || k > '9') {
		return 0, false
	}
	t := c.document.tokens[c.index]
	text := c.document.text[t.start:t.end]
	// Short integer tokens are exactly representable. Validation has already
	// checked their grammar. Preserve negative zero and use ParseFloat for every
	// decimal, exponent and larger integer to retain rounding/range behavior.
	if len(text) <= 15 {
		at := 0
		negative := text[0] == '-'
		if negative {
			at++
		}
		var integer int64
		for at < len(text) && text[at] >= '0' && text[at] <= '9' {
			integer = integer*10 + int64(text[at]-'0')
			at++
		}
		if at == len(text) {
			value := float64(integer)
			if negative {
				value = -value
			}
			return value, true
		}
	}
	n, err := strconv.ParseFloat(text, 64)
	return n, err == nil
}
func (c tcCursor) boolean() (bool, bool) { return c.kind() == 't', c.kind() == 't' || c.kind() == 'f' }
func (c tcCursor) array() (tcArray, bool) {
	if c.kind() != '[' {
		return tcArray{}, false
	}
	return tcArray{c, int(c.document.tokens[c.index].end)}, true
}

type tcIterator struct {
	cursor       tcCursor
	index, count int
	token        uint32
}

func (a tcArray) iter() tcIterator { return tcIterator{a.cursor, 0, a.count, a.cursor.index + 1} }
func (it *tcIterator) more() bool  { return it.index < it.count }
func (it *tcIterator) next() (int, tcCursor) {
	i, token := it.index, it.token
	it.index++
	it.token = it.cursor.document.tokens[token].next
	return i, tcCursor{it.cursor.document, token}
}
func (a tcArray) at(index int) tcCursor {
	c := a.cursor
	at := c.index + 1
	for i := 0; i < index; i++ {
		at = c.document.tokens[at].next
	}
	return tcCursor{c.document, at}
}
func (c tcCursor) Lookup(key string) (tcCursor, bool) {
	if c.kind() != '{' {
		return tcCursor{}, false
	}
	for at := c.document.tokens[c.index].end; at != 0; at = c.document.tokens[at].next {
		if (directCursor{c.document, at}).string(false) == key {
			return tcCursor{c.document, at + 1}, true
		}
	}
	return tcCursor{}, false
}
func (c tcCursor) Keys() []string {
	if c.kind() != '{' {
		return nil
	}
	return directCursor(c).keys()
}

func typedTAST(parse, decode gopurs_runtime.Value, text string) gopurs_runtime.Value {
	cursor, ok := directIndex(text)
	if !ok {
		return gopurs_runtime.Apply(decode, gopurs_runtime.Apply(parse, gopurs_runtime.Str(text)))
	}
	result := tcDecodeModule(gopurs_runtime.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), tcCursor(cursor))
	if tc_cndIsLeft(result) {
		// The diagnostic's decode callback unwraps Either, throwing its original
		// Argonaut message on failure. Keep that public boundary exact.
		return gopurs_runtime.Apply(decode, gopurs_runtime.Apply(parse, gopurs_runtime.Str(text)))
	}
	return (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(result.UnsafePtr).V0
}
