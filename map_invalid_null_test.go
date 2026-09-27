package jsoniter

import (
	"testing"
)

func TestUnmarshalInvalidNullDoesNotNilMap(t *testing.T) {
	for _, input := range []string{"none", "nan", "nil", "notjson", "nu", "nul"} {
		t.Run(input, func(t *testing.T) {
			data := map[string]interface{}{"keep": true}
			err := Unmarshal([]byte(input), &data)
			if err == nil {
				t.Fatalf("expected error for %q, got nil (map=%v)", input, data)
			}
			if data == nil {
				t.Fatalf("pre-initialized map became nil for input %q", input)
			}
			if data["keep"] != true {
				t.Fatalf("map contents changed on invalid input %q: %#v", input, data)
			}
		})
	}
}

func TestUnmarshalNullNilsMap(t *testing.T) {
	data := map[string]interface{}{"keep": true}
	err := Unmarshal([]byte("null"), &data)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Fatalf("null should set map to nil, got %#v", data)
	}
}
