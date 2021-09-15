package sessions

import (
    "MopShopGo/repository"
)

func saveSessionToken(sessionToken SessionsToken) SessionsToken {
    _, err := repository.GetMysql().Exec(""+
        "INSERT INTO sessions "+
        "(email, sessionToken, created, expire) "+
        "VALUES "+
        "(?, ?, ?, ?) "+
        "ON DUPLICATE KEY UPDATE sessionToken = ?",
        sessionToken.Email, sessionToken.Token, sessionToken.Created, sessionToken.Expire, sessionToken.Token,
    )

    if err != nil {
        // todo update error handling
        return SessionsToken{}
    }

    return sessionToken
}
