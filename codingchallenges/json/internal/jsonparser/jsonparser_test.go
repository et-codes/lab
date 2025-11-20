package jsonparser

import "testing"

func TestStep1(t *testing.T) {
	t.Run("returns 0 with valid JSON format", func(t *testing.T) {
		args := []string{"../../test_data/step1/valid.json"}
		jp := NewJSONParser(args)
		code := jp.Run()
		if code != 0 {
			t.Errorf("expected code 0, got %d", code)
		}
	})

	t.Run("returns 1 with invalid JSON format", func(t *testing.T) {
		args := []string{"../../test_data/step1/invalid.json"}
		jp := NewJSONParser(args)
		code := jp.Run()
		if code != 1 {
			t.Errorf("expected code 1, got %d", code)
		}
	})
}
