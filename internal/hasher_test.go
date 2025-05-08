package internal

import "testing"

func TestHashJSON(t *testing.T) {
    input := []byte(`{"hi":"there"}`)
    hash, err := HashJSON(input)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if len(hash) != 64 {
        t.Errorf("unexpected hash length: got %d, want 64", len(hash))
    }
}
