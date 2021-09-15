package sessions

import (
    "github.com/google/uuid"
    "time"
)

func CreateSessionToken(email string) SessionsToken {
    token := uuid.New().String()

    sessionToken := SessionsToken{
        Email:   email,
        Token:   token,
        Created: time.Now(),
        Expire:  time.Now().Add(60 * time.Minute),
    }

    return saveSessionToken(sessionToken)
}

func GetSessionToken(sessionToken string) (SessionsToken, error) {
    return getSessionToken(sessionToken)
}

func IsValidToken(token SessionsToken) bool {
    // todo: does a token valid
    return true
}