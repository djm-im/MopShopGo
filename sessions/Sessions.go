package sessions

import "time"

type SessionsToken struct {
    Email   string
    Token   string
    Created time.Time
    Expire  time.Time
}
