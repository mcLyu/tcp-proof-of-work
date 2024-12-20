package hashing

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
)

type Algorithm string

const (
	algorithmSha256 Algorithm = "SHA-256"
	algorithmSha1   Algorithm = "SHA-1"
)

func GetDefaultAlgorithm() Algorithm {
	return algorithmSha256
}

func (a Algorithm) GetHasher() (hash.Hash, error) {
	switch a {
	case algorithmSha256:
		return sha256.New(), nil
	case algorithmSha1:
		return sha1.New(), nil
	default:
		return nil, fmt.Errorf("no hash function found for hash algorithm %s", a)
	}
}
