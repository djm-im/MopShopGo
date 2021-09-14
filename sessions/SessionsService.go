package sessions

import (
    "github.com/google/uuid"
    "time"
)

func CreateSessionToken(email string) SessionToken {
    token := uuid.New().String()
    sessionToken := SessionToken{
        Email:   email,
        Token:   token,
        Created: time.Now(),
        Expire:  time.Now().Add(60 * time.Minute),
    }

    sessionsRepository.sessions[email] = sessionToken

    return sessionToken
}
