package purescript

// Validate syntax and build a call-local index together. Every value, including
// ignored members, is validated before schema decoding starts. Strings become
// owned only when they enter the final result; cursors never escape as Json.
import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// Scalars store their source end (plus a string-normalization flag), arrays
// their length, and objects the last key token. Key.next links to the preceding
// key in that object; value.next always skips the complete value subtree.
type directToken struct{ start, end, next uint32 }
type directDocument struct {
	text   string
	tokens []directToken
}
type directCursor struct {
	document *directDocument
	index    uint32
}
type directParser struct {
	document directDocument
	pos      int
}

const directNormalize uint32 = 1 << 31

func directIndex(text string) (directCursor, bool) {
	if uint64(len(text)) >= uint64(directNormalize) {
		return directCursor{}, false
	}
	p := directParser{document: directDocument{text: text, tokens: make([]directToken, 0, len(text)/6+8)}}
	if !p.value(0) {
		return directCursor{}, false
	}
	p.space()
	if p.pos != len(text) {
		return directCursor{}, false
	}
	return directCursor{&p.document, 0}, true
}
func (p *directParser) space() {
	for p.pos < len(p.document.text) {
		switch p.document.text[p.pos] {
		case ' ', '\t', '\n', '\r':
			p.pos++
		default:
			return
		}
	}
}
func (p *directParser) take(c byte) bool {
	if p.pos < len(p.document.text) && p.document.text[p.pos] == c {
		p.pos++
		return true
	}
	return false
}
func (p *directParser) value(depth int) bool {
	p.space()
	text := p.document.text
	if p.pos == len(text) {
		return false
	}
	index := len(p.document.tokens)
	start := p.pos
	switch text[start] {
	case '{', '[':
		if depth >= 10000 {
			return false
		}
		object := text[start] == '{'
		end := byte(']')
		if object {
			end = '}'
		}
		p.pos++
		p.document.tokens = append(p.document.tokens, directToken{start: uint32(start)})
		p.space()
		var count, lastKey uint32
		if !p.take(end) {
			for {
				if object {
					key := uint32(len(p.document.tokens))
					if !p.string() {
						return false
					}
					p.document.tokens[key].next = lastKey
					lastKey = key
					p.space()
					if !p.take(':') {
						return false
					}
				}
				if !p.value(depth + 1) {
					return false
				}
				count++
				p.space()
				if p.take(end) {
					break
				}
				if !p.take(',') {
					return false
				}
				p.space()
			}
		}
		if object {
			count = lastKey
		}
		p.document.tokens[index].end = count
		p.document.tokens[index].next = uint32(len(p.document.tokens))
		return true
	case '"':
		return p.string()
	case 't':
		if !strings.HasPrefix(text[start:], "true") {
			return false
		}
		p.pos += 4
	case 'f':
		if !strings.HasPrefix(text[start:], "false") {
			return false
		}
		p.pos += 5
	case 'n':
		if !strings.HasPrefix(text[start:], "null") {
			return false
		}
		p.pos += 4
	default:
		if !p.number() {
			return false
		}
	}
	p.document.tokens = append(p.document.tokens, directToken{uint32(start), uint32(p.pos), uint32(index + 1)})
	return true
}
func (p *directParser) string() bool {
	text := p.document.text
	start := p.pos
	if !p.take('"') {
		return false
	}
	escaped, nonASCII := false, false
	for p.pos < len(text) {
		c := text[p.pos]
		p.pos++
		if c == '"' {
			end := uint32(p.pos)
			if escaped || (nonASCII && !utf8.ValidString(text[start+1:p.pos-1])) {
				end |= directNormalize
			}
			p.document.tokens = append(p.document.tokens, directToken{uint32(start), end, uint32(len(p.document.tokens) + 1)})
			return true
		}
		if c < 0x20 {
			return false
		}
		if c >= utf8.RuneSelf {
			nonASCII = true
		}
		if c != '\\' {
			continue
		}
		escaped = true
		if p.pos == len(text) {
			return false
		}
		escape := text[p.pos]
		p.pos++
		switch escape {
		case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		case 'u':
			if _, ok := argonautJSONHex4(text, p.pos); !ok {
				return false
			}
			p.pos += 4
		default:
			return false
		}
	}
	return false
}
func (p *directParser) number() bool {
	text := p.document.text
	start := p.pos
	p.take('-')
	if p.pos == len(text) {
		return false
	}
	if !p.take('0') {
		if text[p.pos] < '1' || text[p.pos] > '9' {
			return false
		}
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
	}
	integer := true
	if p.take('.') {
		integer = false
		digits := p.pos
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
		if digits == p.pos {
			return false
		}
	}
	if p.pos < len(text) && (text[p.pos] == 'e' || text[p.pos] == 'E') {
		integer = false
		p.pos++
		if !p.take('+') {
			p.take('-')
		}
		digits := p.pos
		for p.pos < len(text) && text[p.pos] >= '0' && text[p.pos] <= '9' {
			p.pos++
		}
		if digits == p.pos {
			return false
		}
	}
	// Short integers are finite by construction. Overflow in every other
	// number must fail even if the schema never reads that member.
	if !integer || p.pos-start > 15 {
		if _, err := strconv.ParseFloat(text[start:p.pos], 64); err != nil {
			return false
		}
	}
	return true
}
func (cursor directCursor) kind() byte {
	return cursor.document.text[cursor.document.tokens[cursor.index].start]
}
func (cursor directCursor) string(owned bool) string {
	token := cursor.document.tokens[cursor.index]
	text := cursor.document.text[token.start+1 : (token.end&^directNormalize)-1]
	if token.end&directNormalize != 0 {
		value, ok := argonautUnquoteJSON(text)
		if !ok {
			panic("validated string")
		}
		return value
	}
	if owned {
		return strings.Clone(text)
	}
	return text
}
func (cursor directCursor) native() any {
	token := cursor.document.tokens[cursor.index]
	switch cursor.kind() {
	case '{':
		return cursor
	case '[':
		items := make([]any, token.end)
		i := 0
		for at := cursor.index + 1; at < token.next; at = cursor.document.tokens[at].next {
			items[i] = directCursor{cursor.document, at}
			i++
		}
		return items
	case '"':
		return cursor.string(true)
	case 't':
		return true
	case 'f':
		return false
	case 'n':
		return nil
	default:
		value, err := strconv.ParseFloat(cursor.document.text[token.start:token.end], 64)
		if err != nil {
			panic(err)
		}
		return value
	}
}
func (cursor directCursor) JSONLookup(key string) (any, bool) {
	value, ok := tcCursor(cursor).Lookup(key)
	if !ok {
		return nil, false
	}
	return directCursor(value), true
}
func (cursor directCursor) keys() []string {
	keys := []string{}
	for at := cursor.index + 1; at < cursor.document.tokens[cursor.index].next; {
		name := (directCursor{cursor.document, at}).string(true)
		exists := false
		for _, key := range keys {
			if key == name {
				exists = true
				break
			}
		}
		if !exists {
			keys = append(keys, name)
		}
		at = cursor.document.tokens[at+1].next
	}
	return keys
}

// Only the existing cold source-span boundary materializes a public Json.
func directMaterialize(raw any) any {
	cursor, ok := raw.(directCursor)
	if !ok {
		return ntNative(raw)
	}
	if cursor.kind() == '{' {
		object := make(map[string]any)
		for _, key := range cursor.keys() {
			value, _ := cursor.JSONLookup(key)
			object[key] = directMaterialize(value)
		}
		return object
	}
	value := cursor.native()
	if array, ok := value.([]any); ok {
		for i, item := range array {
			array[i] = directMaterialize(item)
		}
	}
	return value
}
