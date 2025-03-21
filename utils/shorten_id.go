package utils

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

// EncodeUUIDToBase62 converts UUID to a Base62 string
func EncodeUUIDToBase62(id uuid.UUID) string {
	encoded := base64.RawURLEncoding.EncodeToString(id[:])
	return strings.NewReplacer("-", "", "_", "").Replace(encoded) // Remove non-Base62 chars
}

// DecodeBase62ToUUID converts Base62 back to UUID
func DecodeBase62ToUUID(base62 string) (uuid.UUID, error) {
	data, err := base64.RawURLEncoding.DecodeString(base62)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.FromBytes(data)
}
