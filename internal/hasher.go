package internal

import (
    "crypto/sha256"
    "encoding/json"
    "fmt"
)

// HashJSON takes JSON input and returns a SHA-256 hash.
func HashJSON(input []byte) (string, error) {
    var temp map[string]interface{}
    if err := json.Unmarshal(input, &temp); err != nil {
        return "", fmt.Errorf("invalid JSON: %w", err)
    }

    normalized, err := json.Marshal(temp)
    if err != nil {
        return "", fmt.Errorf("failed to normalize JSON: %w", err)
    }

    hash := sha256.Sum256(normalized)
    return fmt.Sprintf("%x", hash[:]), nil
}
