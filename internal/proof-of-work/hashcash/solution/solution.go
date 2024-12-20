package solution

import (
	"fmt"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/challenge"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/hashing"
	"strconv"
	"strings"
	"time"
)

type Solution struct {
	challenge.Challenge
	Answer string
}

func FromString(stringSolution string) (*Solution, error) {
	params := strings.Split(stringSolution, ":")
	if len(params) != 7 {
		return nil, fmt.Errorf("wrong format of hashcash solution string, should be exactly 7 params between semicolons")
	}

	difficulty, _ := strconv.Atoi(params[1])
	expiresAt, _ := strconv.ParseInt(params[2], 10, 64)
	if time.Now().Sub(time.Unix(expiresAt, 0)) > 0 {
		return nil, fmt.Errorf("hashcash challenge expired")
	}

	return &Solution{
		Challenge: challenge.Challenge{
			Tag:        params[0],
			Difficulty: difficulty,
			ExpiresAt:  expiresAt,
			Subject:    params[3],
			Algorithm:  hashing.Algorithm(params[4]),
			Salt:       params[5],
		},
		Answer: params[6],
	}, nil
}
