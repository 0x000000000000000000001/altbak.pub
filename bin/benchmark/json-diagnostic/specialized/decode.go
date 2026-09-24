// Audit-only: hand-written, schema-specialised decoder for Test.JsonDecoding.
//
// Purpose: measure what a schema-specialised code path can do while consuming
// the same parsed Json as the generated Argonaut decoder and producing the
// same final PureScript values (records, Maybe, the Event ADT, Either).
//
// It reads the parser's DOM directly (map[string]any / []any / string /
// float64 / bool / nil) and builds the final gopurs values with the runtime
// constructors. It does not define a new contract: error values, field order
// and wrappers must match the generated decoder exactly, which the audit
// runner checks against the frozen oracle for all seventeen corpus cases.
//
// This file is copied into the audit workspace's output/purescript/ and is
// never part of production code.
package purescript

import (
	"math"
	"unsafe"

	"gopurs/output/gopurs_runtime"
)

// Constructor tags, read from the generated modules in the frozen workspace.
const (
	zzEitherLeft  = 3711209382
	zzEitherRight = 2465973597
	zzMaybeJust   = 930809136

	zzErrTypeMismatch = 2887704423
	zzErrAtIndex      = 1044667600
	zzErrAtKey        = 1896025177
	zzErrNamed        = 2718035288
	zzErrMissingValue = 3199441748

	zzViewTag     = 2782777920
	zzPurchaseTag = 3937542790
)

// ---------------------------------------------------------------------------
// Value constructors, matching the generated encodings
// ---------------------------------------------------------------------------

func zzLeft(err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzEitherLeft, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{Rc: 1, V0: err})}
}

func zzRight(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzEitherRight, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{Rc: 1, V0: value})}
}

func zzNothing() gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzMaybeJust}
}

func zzJust(value gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzMaybeJust, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: value})}
}

func zzMissingValue() gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzErrMissingValue}
}

func zzTypeMismatch(message string) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzErrTypeMismatch, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Argonaut_Decode_Error_TypeMismatch{Rc: 1, V0: message})}
}

func zzAtKey(key string, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzErrAtKey, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Argonaut_Decode_Error_AtKey{Rc: 1, V0: key, V1: err})}
}

func zzAtIndex(index int, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzErrAtIndex, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Argonaut_Decode_Error_AtIndex{Rc: 1, V0: int64(index), V1: err})}
}

func zzNamed(name string, err gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{Type: 9, IntVal: zzErrNamed, UnsafePtr: unsafe.Pointer(
		&Constructor_Data_Argonaut_Decode_Error_Named{Rc: 1, V0: name, V1: err})}
}

// ---------------------------------------------------------------------------
// Scalar decoders (same rules as Data.Argonaut.Decode.Decoders)
// ---------------------------------------------------------------------------

// zzIntFromNumber mirrors Data.Int.FromNumberImpl on this runtime.
func zzIntFromNumber(n float64) (int64, bool) {
	if math.IsNaN(n) || math.IsInf(n, 0) || math.Trunc(n) != n {
		return 0, false
	}
	return int64(int(n)), true
}

func zzDecodeNumber(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	n, ok := raw.(float64)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Number"), false
	}
	return gopurs_runtime.Float(n), gopurs_runtime.Value{}, true
}

func zzDecodeInt(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	n, ok := raw.(float64)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Number"), false
	}
	value, ok := zzIntFromNumber(n)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Integer"), false
	}
	return gopurs_runtime.Int(value), gopurs_runtime.Value{}, true
}

func zzDecodeString(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	text, ok := raw.(string)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("String"), false
	}
	return gopurs_runtime.Str(text), gopurs_runtime.Value{}, true
}

func zzDecodeBoolean(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	value, ok := raw.(bool)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Boolean"), false
	}
	return gopurs_runtime.Bool(value), gopurs_runtime.Value{}, true
}

type zzDecoder func(any) (gopurs_runtime.Value, gopurs_runtime.Value, bool)

// ---------------------------------------------------------------------------
// Field accessors (same wrappers as getField / getFieldOptional' / decodeFieldMaybe)
// ---------------------------------------------------------------------------

func zzField(obj map[string]any, key string, decode zzDecoder) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	raw, present := obj[key]
	if !present {
		return gopurs_runtime.Value{}, zzAtKey(key, zzMissingValue()), false
	}
	value, err, ok := decode(raw)
	if !ok {
		return gopurs_runtime.Value{}, zzAtKey(key, err), false
	}
	return value, gopurs_runtime.Value{}, true
}

func zzMaybeField(obj map[string]any, key string, decode zzDecoder) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	raw, present := obj[key]
	if !present || raw == nil {
		return zzNothing(), gopurs_runtime.Value{}, true
	}
	value, err, ok := decode(raw)
	if !ok {
		return gopurs_runtime.Value{}, zzAtKey(key, err), false
	}
	return zzJust(value), gopurs_runtime.Value{}, true
}

// zzArrayField mirrors decodeArray: a non-array container fails with a plain
// TypeMismatch "Array" (the Kleisli composition short-circuits before the
// Named wrapper), while an element error is wrapped as
// AtKey (field) over Named "Array" over AtIndex.
func zzArrayField(obj map[string]any, key string, decode zzDecoder) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	raw, present := obj[key]
	if !present {
		return gopurs_runtime.Value{}, zzAtKey(key, zzMissingValue()), false
	}
	items, ok := raw.([]any)
	if !ok {
		return gopurs_runtime.Value{}, zzAtKey(key, zzTypeMismatch("Array")), false
	}
	values := make([]gopurs_runtime.Value, len(items))
	for index, item := range items {
		value, err, ok := decode(item)
		if !ok {
			return gopurs_runtime.Value{}, zzAtKey(key, zzNamed("Array", zzAtIndex(index, err))), false
		}
		values[index] = value
	}
	return gopurs_runtime.Array(values), gopurs_runtime.Value{}, true
}

// ---------------------------------------------------------------------------
// Record and ADT decoders (fields in the same order as the generated RowList)
// ---------------------------------------------------------------------------

func zzDecodeProfile(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	obj, ok := raw.(map[string]any)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Object"), false
	}
	city, err, ok := zzField(obj, "city", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	note, err, ok := zzMaybeField(obj, "note", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	scores, err, ok := zzArrayField(obj, "scores", zzDecodeNumber)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	return gopurs_runtime.RecordDict3("city", "note", "scores", city, note, scores), gopurs_runtime.Value{}, true
}

func zzDecodeItem(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	obj, ok := raw.(map[string]any)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Object"), false
	}
	price, err, ok := zzField(obj, "price", zzDecodeNumber)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	quantity, err, ok := zzField(obj, "quantity", zzDecodeInt)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	sku, err, ok := zzField(obj, "sku", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	return gopurs_runtime.RecordDict3("price", "quantity", "sku", price, quantity, sku), gopurs_runtime.Value{}, true
}

func zzDecodeUser(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	obj, ok := raw.(map[string]any)
	if !ok {
		return gopurs_runtime.Value{}, zzTypeMismatch("Object"), false
	}
	active, err, ok := zzField(obj, "active", zzDecodeBoolean)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	id, err, ok := zzField(obj, "id", zzDecodeInt)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	name, err, ok := zzField(obj, "name", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	profile, err, ok := zzMaybeField(obj, "profile", zzDecodeProfile)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	tags, err, ok := zzArrayField(obj, "tags", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	return gopurs_runtime.RecordDict5("active", "id", "name", "profile", "tags", active, id, name, profile, tags), gopurs_runtime.Value{}, true
}

func zzDecodeEvent(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	obj, ok := raw.(map[string]any)
	if !ok {
		// decodeForeignObject short-circuits before Named "ForeignObject", like
		// decodeArray does before Named "Array".
		return gopurs_runtime.Value{}, zzTypeMismatch("Object"), false
	}
	tagRaw, present := obj["tag"]
	if !present {
		return gopurs_runtime.Value{}, zzAtKey("tag", zzMissingValue()), false
	}
	tag, ok := tagRaw.(string)
	if !ok {
		return gopurs_runtime.Value{}, zzAtKey("tag", zzTypeMismatch("String")), false
	}
	switch tag {
	case "view":
		path, err, ok := zzField(obj, "path", zzDecodeString)
		if !ok {
			return gopurs_runtime.Value{}, err, false
		}
		var duration *Constructor_Data_Maybe_Just[int64]
		if durationRaw, present := obj["duration"]; present && durationRaw != nil {
			value, err, ok := zzDecodeInt(durationRaw)
			if !ok {
				return gopurs_runtime.Value{}, zzAtKey("duration", err), false
			}
			duration = &Constructor_Data_Maybe_Just[int64]{Rc: 1, V0: value.IntVal}
		}
		view := gopurs_runtime.Value{Type: 9, IntVal: zzViewTag, UnsafePtr: unsafe.Pointer(
			&Constructor_Test_JsonDecoding_View{Rc: 1, V0: path.StrVal(), V1: duration})}
		return view, gopurs_runtime.Value{}, true
	case "purchase":
		orderID, err, ok := zzField(obj, "orderId", zzDecodeInt)
		if !ok {
			return gopurs_runtime.Value{}, err, false
		}
		items, err, ok := zzArrayField(obj, "items", zzDecodeItem)
		if !ok {
			return gopurs_runtime.Value{}, err, false
		}
		purchase := gopurs_runtime.Value{Type: 9, IntVal: zzPurchaseTag, UnsafePtr: unsafe.Pointer(
			&Constructor_Test_JsonDecoding_Purchase{Rc: 1, V0: orderID.IntVal, V1: *(*[]gopurs_runtime.Value)(items.UnsafePtr)})}
		return purchase, gopurs_runtime.Value{}, true
	default:
		return gopurs_runtime.Value{}, zzTypeMismatch("Event tag"), false
	}
}

// ZzSpecializedDecode decodes the parser's Json into the same Either value the
// generated decoder produces. Fields are evaluated in the generated order:
// events, next, users, version for Payload; active, id, name, profile, tags
// for User; price, quantity, sku for Item; city, note, scores for Profile.
//
// Only an object root arrives as TypeAny wrapping map[string]any. An array
// root is TypeArray and every scalar root keeps its scalar type; all of them
// fail like toObject does, with a plain TypeMismatch "Object".
func ZzSpecializedDecode(json gopurs_runtime.Value) gopurs_runtime.Value {
	var raw any
	if json.Type == gopurs_runtime.TypeAny && json.UnsafePtr != nil {
		raw = *(*any)(json.UnsafePtr)
	}
	value, err, ok := zzDecodePayload(raw)
	if !ok {
		return zzLeft(err)
	}
	return zzRight(value)
}

func zzDecodePayload(raw any) (gopurs_runtime.Value, gopurs_runtime.Value, bool) {
	obj, ok := raw.(map[string]any)
	if !ok {
		// The record instance reports a plain TypeMismatch, without the
		// Named "ForeignObject" wrapper used for FO.Object fields.
		return gopurs_runtime.Value{}, zzTypeMismatch("Object"), false
	}
	events, err, ok := zzArrayField(obj, "events", zzDecodeEvent)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	next, err, ok := zzMaybeField(obj, "next", zzDecodeString)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	users, err, ok := zzArrayField(obj, "users", zzDecodeUser)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	version, err, ok := zzField(obj, "version", zzDecodeInt)
	if !ok {
		return gopurs_runtime.Value{}, err, false
	}
	payload := gopurs_runtime.RecordDict4("events", "next", "users", "version", events, next, users, version)
	return payload, gopurs_runtime.Value{}, true
}
