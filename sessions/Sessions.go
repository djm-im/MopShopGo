package sessions

import "time"

type SessionToken struct {
    Email   string
    Token   string
    Created time.Time
    Expire  time.Time
}
