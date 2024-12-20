package challenge

import "testing"

var validChallengeSolutionsMap = map[string]string{
	"H:21:1734852810:172.18.0.1:SHA-256:KRm9wg0TFA==": "H:21:1734852810:172.18.0.1:SHA-256:KRm9wg0TFA==:1ad30f",
	"H:22:1734852810:172.18.0.1:SHA-256:miZc8xKv2Q==": "H:22:1734852810:172.18.0.1:SHA-256:miZc8xKv2Q==:bbee6",
}

var invalidChallengeStrings = []string{
	"H:21:1734852810:",
	"Some string",
	"H:23:1734852810:172.18.0.1:SHA-256:RjdVb_sJCA==:697da0",
}

func TestInvalidChallengeFromString(t *testing.T) {
	for _, challengeStr := range invalidChallengeStrings {
		challenge, err := FromString(challengeStr)
		if err == nil {
			t.Errorf("Expected error for invalid challenge: %s", challengeStr)
		}

		if challenge != nil {
			t.Errorf("Challenge is not nil")
		}
	}
}

func TestValidChallengeFromString(t *testing.T) {
	for challengeStr := range validChallengeSolutionsMap {
		challenge, err := FromString(challengeStr)
		if err != nil {
			t.Errorf("Failed to parse challenge from string: %s", challengeStr)
		}

		if challenge == nil {
			t.Errorf("Challenge is nil")
		}
	}
}

func TestValidChallengeFindSolution(t *testing.T) {
	for challengeStr, solution := range validChallengeSolutionsMap {
		challenge, _ := FromString(challengeStr)
		answer, err := challenge.FindSolution()
		if err != nil {
			t.Errorf("Failed to find solution for challenge: %s", challengeStr)
		}

		if answer == nil {
			t.Errorf("Solution is nil")
		}

		if *answer != solution {
			t.Errorf("Invalid solution for challenge: %s, expected: %s, got: %s", challengeStr, solution, *answer)
		}
	}
}
