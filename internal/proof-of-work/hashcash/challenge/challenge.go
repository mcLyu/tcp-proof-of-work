package challenge

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/hashing"
	"hash"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// defaultTag represents Hashcash V1 version
	defaultTag = "H"

	// saltMaximumLen is the maximum length of salt string
	saltMaximumLen = 7

	// DefaultDifficulty is the default number of leading 0-bits that needs to be in the solution
	DefaultDifficulty = 20
)

// Challenge represents a Hashcash challenge
// Adapted from https://therootcompany.com/blog/http-hashcash/
type Challenge struct {
	// Tag is the version of Hashcash Algorithm Version
	Tag string
	// Difficulty is the number of leading 0-bits that needs to be in the solution
	Difficulty int
	// ExpiresAt is the unix time when this challenge expires
	ExpiresAt int64
	// Subject is the resource requesting the challenge
	Subject string
	// Algorithm is the hashing algorithm
	Algorithm hashing.Algorithm
	// Salt is the random salt string to be added in challenge
	Salt string
}

func NewDefaultChallenge(subject string, difficulty int) *Challenge {
	return &Challenge{
		Tag:        defaultTag,
		Difficulty: difficulty,
		ExpiresAt:  time.Now().Add(48 * time.Hour).Unix(),
		Subject:    subject,
		Algorithm:  hashing.GetDefaultAlgorithm(),
		Salt:       generateSalt(),
	}
}

func NewChallenge(tag string, difficulty int, expiresAt int64, subject string, algorithm hashing.Algorithm) *Challenge {
	return &Challenge{
		Tag:        tag,
		Difficulty: difficulty,
		ExpiresAt:  expiresAt,
		Subject:    subject,
		Algorithm:  algorithm,
		Salt:       generateSalt(),
	}
}

func (c *Challenge) String() string {
	return fmt.Sprintf("%s:%d:%d:%s:%s:%s", c.Tag, c.Difficulty, c.ExpiresAt, c.Subject, c.Algorithm, c.Salt)
}

// FromString converts input hashcash challenge string to HashcashChallenge struct
func FromString(challenge string) (*Challenge, error) {
	params := strings.Split(challenge, ":")
	if len(params) != 6 {
		return nil, fmt.Errorf("wrong format of hashcash challenge string, should be exactly 6 params between semicolons")
	}

	difficulty, _ := strconv.Atoi(params[1])
	expiresAt, _ := strconv.ParseInt(params[2], 10, 64)

	return &Challenge{
		Tag:        params[0],
		Difficulty: difficulty,
		ExpiresAt:  expiresAt,
		Subject:    params[3],
		Algorithm:  hashing.Algorithm(params[4]),
		Salt:       params[5],
	}, nil
}

func (c *Challenge) FindSolution() (*string, error) {
	answer := 0
	challengeString := c.String()

	var solution string
	hasher, err := c.Algorithm.GetHasher()
	if err != nil {
		fmt.Printf("Error getting hash function for algorithm %s", c.Algorithm)
		return nil, err
	}

	for {
		solution = fmt.Sprintf("%s:%x", challengeString, answer)
		if c.IsValidSolution(solution, hasher) {
			return &solution, nil
		}

		answer++
	}
}

func (c *Challenge) IsValidSolution(solution string, hasher hash.Hash) bool {
	hasher.Reset()
	hasher.Write([]byte(solution))
	sum := hasher.Sum(nil)
	sumUint64 := binary.BigEndian.Uint64(sum)
	sumBits := strconv.FormatUint(sumUint64, 2)
	zeroes := 64 - len(sumBits)

	return zeroes >= c.Difficulty
}

// generateSalt generates a random URL-safe salt
func generateSalt() string {
	b := make([]byte, saltMaximumLen)
	_, err := rand.Read(b)
	if err != nil {
		log.Print("Error salt generation")
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}
