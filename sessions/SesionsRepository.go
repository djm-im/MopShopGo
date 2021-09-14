package sessions

type SessionsRepository struct {
    sessions map[string]SessionToken
}

var sessionsRepository = SessionsRepository{
    sessions: map[string]SessionToken{},
}
