package sessions

import (
    "github.com/google/uuid"
    "time"
)

func CreateSessionToken(email string) SessionsToken {
    token := uuid.New().String()

    sessionToken := SessionsToken{
        Email: email,
        Token: token,
        Created: time.Now(),
        Expire:  time.Now().Add(60 * time.Minute),
    }

    savedSessionsToken := saveSessionToken(sessionToken)

    return savedSessionsToken
}
