// Package orderedobject allows you to decode JSON dicts while preserving order.
// It uses the standard MarshalJSON and UnmarshalJSON interface and is thus compatible with most json encoders/decoders.
package orderedobject

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Member is a single key-value pair in an OrderedObject.
type Member[V any] struct {
	Key   string
	Value V
}

// OrderedObject is a JSON dict that preserves order. Keys are strings, and values are generically typed V.
type OrderedObject[V any] []Member[V]

func (o *OrderedObject[V]) UnmarshalJSON(b []byte) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if t == nil {
		// By convention, to approximate the behavior of Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
		return nil
	}
	d, ok := t.(json.Delim)
	if !ok {
		return fmt.Errorf("unexpected %T at start of OrderedObject", t)
	}
	if d != '{' {
		return fmt.Errorf("unexpected %s at start of OrderedObject", d)
	}
	for dec.More() {
		k, err := dec.Token()
		if err != nil {
			return err
		}
		var v V
		if err := dec.Decode(&v); err != nil {
			return err
		}
		*o = append(*o, Member[V]{k.(string), v})
	}
	// We are allowed to assume it's valid JSON, so we don't need to check this is a closing brace.
	if _, err := dec.Token(); err != nil {
		return err
	}
	return nil
}

func (o OrderedObject[V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, m := range o {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err := e.Encode(m.Key); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		if err := e.Encode(m.Value); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
