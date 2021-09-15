package sessions

import (
    "MopShopGo/repository"
    "errors"
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

func getSessionToken(token string) (SessionsToken, error) {
    result, _ := repository.GetMysql().Query(""+
        "SELECT * "+
        "FROM sessions "+
        "WHERE sessionToken = ?", token,
    )

    var sessionToken SessionsToken
    if result.Next() {
        result.Scan(&sessionToken.Email, &sessionToken.Token, &sessionToken.Created, &sessionToken.Expire)

        return sessionToken, nil
    }

    return SessionsToken{}, errors.New("Canot find session token.")
}
