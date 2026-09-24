package purescript

// Experimental, call-local structural index. The ordinary TAST decoder consumes
// it synchronously; only decoded, owned strings and final typed values escape.
// No Json containing a cursor is published by the combined driver.
import (
	"encoding/json"
	"gopurs/output/gopurs_runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

type directToken struct{ start, end, next uint32 }
type directDocument struct {
	text   string
	tokens []directToken
}
type directCursor struct {
	document *directDocument
	index    uint32
}

func directIndex(text string) (directCursor, bool) {
	// Full validation precedes schema decoding, including ignored members.
	// The existing parser supplies its exact error on invalid/overflow input.
	if uint64(len(text)) >= 1<<32 || !json.Valid([]byte(text)) {
		return directCursor{}, false
	}
	doc := &directDocument{text: text, tokens: make([]directToken, 0, len(text)/6+8)}
	var local [128]uint32
	stack := local[:0]
	for pos := 0; pos < len(text); {
		start := pos
		switch text[pos] {
		case ' ', '\n', '\r', '\t', ':', ',':
			pos++
			continue
		case '}', ']':
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pos++
			doc.tokens[index].end = uint32(pos)
			doc.tokens[index].next = uint32(len(doc.tokens))
			continue
		case '{', '[':
			index := uint32(len(doc.tokens))
			doc.tokens = append(doc.tokens, directToken{start: uint32(pos)})
			stack = append(stack, index)
			pos++
			continue
		case '"':
			pos++
			for text[pos] != '"' {
				if text[pos] == '\\' {
					pos++
				}
				pos++
			}
			pos++
		case 't':
			pos += 4
		case 'f':
			pos += 5
		case 'n':
			pos += 4
		default:
			for pos < len(text) && ((text[pos] >= '0' && text[pos] <= '9') || text[pos] == '-' || text[pos] == '+' || text[pos] == '.' || text[pos] == 'e' || text[pos] == 'E') {
				pos++
			}
			number := text[start:pos]
			if len(number) > 15 || strings.ContainsAny(number, ".eE") {
				if _, err := strconv.ParseFloat(number, 64); err != nil {
					return directCursor{}, false
				}
			}
		}
		doc.tokens = append(doc.tokens, directToken{uint32(start), uint32(pos), uint32(len(doc.tokens) + 1)})
	}
	return directCursor{doc, 0}, true
}

func (cursor directCursor) kind() byte {
	return cursor.document.text[cursor.document.tokens[cursor.index].start]
}
func (cursor directCursor) string(owned bool) string {
	token := cursor.document.tokens[cursor.index]
	text := cursor.document.text[token.start+1 : token.end-1]
	if strings.IndexByte(text, '\\') >= 0 || !utf8.ValidString(text) {
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
		count := 0
		for at := cursor.index + 1; at < token.next; at = cursor.document.tokens[at].next {
			count++
		}
		items := make([]any, count)
		index := 0
		for at := cursor.index + 1; at < token.next; at = cursor.document.tokens[at].next {
			items[index] = directCursor{cursor.document, at}
			index++
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
	var selected uint32
	found := false
	for at := cursor.index + 1; at < cursor.document.tokens[cursor.index].next; {
		name := directCursor{cursor.document, at}
		value := at + 1
		if name.string(false) == key {
			selected = value
			found = true
		}
		at = cursor.document.tokens[value].next
	}
	if !found {
		return nil, false
	}
	return directCursor{cursor.document, selected}, true
}

// Enumeration is uncommon (reExports and literal records); preserve unique
// keys, last-value precedence and owned keys just like an ordinary object.
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
func (cursor directCursor) JSONLength() int { return len(cursor.keys()) }
func (cursor directCursor) JSONEntry(index int) (string, any) {
	key := cursor.keys()[index]
	value, _ := cursor.JSONLookup(key)
	return key, value
}

// The existing cold source-span decoder delegates to generic Argonaut tuple
// decoding. Materialize just this subtree at that public Json boundary.
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

func directTAST(parse, decode gopurs_runtime.Value, text string) gopurs_runtime.Value {
	cursor, ok := directIndex(text)
	if !ok {
		return gopurs_runtime.Apply(decode, gopurs_runtime.Apply(parse, gopurs_runtime.Str(text)))
	}
	return gopurs_runtime.Apply(decode, gopurs_runtime.Any(cursor))
}
