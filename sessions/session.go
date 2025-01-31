package sessions

import "github.com/google/uuid"

// TODO: replace this w/ redis for redundancy
// TODO: make tokens expire after some period of time

var sessionMap = make(map[string][]string)

func NewSession(userID string) string {
	token := uuid.NewString()
	sessionMap[userID] = append(sessionMap[userID], token)
	return token
}

func ValidToken(token string) bool {
	for _, list := range sessionMap {
		for _, t := range list {
			if token == t {
				// this is possibly insecure but that's a problem for future me
				return true
			}
		}
	}
	return false
}

func GetTokens(userID string) []string {
	return sessionMap[userID]
}
