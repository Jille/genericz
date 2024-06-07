package orderedobject

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOrderedObject(t *testing.T) {
	tests := []struct {
		in      string
		want    OrderedObject[int]
		wantErr string
	}{
		{
			in: "null",
		},
		{
			in: "{}",
		},
		{
			in:   `{"x": 5}`,
			want: OrderedObject[int]{{"x", 5}},
		},
		{
			in:   `{"x": 5, "a": 7}`,
			want: OrderedObject[int]{{"x", 5}, {"a", 7}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			var got OrderedObject[int]
			if err := json.Unmarshal([]byte(tc.in), &got); err != nil {
				if tc.wantErr == err.Error() {
					return
				}
				t.Fatalf("Unmarshal failed with wrong error: %v; want %q", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Incorrect result: got %#v, want %#v", got, tc.want)
			}
		})
	}
}
