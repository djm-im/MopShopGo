package purchases

import (
    "MopShopGo/exceptions"
    "MopShopGo/sessions"
    "MopShopGo/users"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/paymentintent"
    "net/http"
    "strconv"
)

func Buy(response http.ResponseWriter, request *http.Request) {
    cookie, err1 := request.Cookie("session_token")
    if err1 != nil {
        if err1 == http.ErrNoCookie {
            exceptions.BuildErrorResponse(
                &response,
                http.StatusUnauthorized,
                fmt.Sprintf("User is not authorized."))
            return
        }

        // other errors
        exceptions.BuildErrorResponse(
            &response,
            http.StatusBadRequest,
            fmt.Sprintf("Not proper requesrt."))
        return
    }

    token := cookie.Value

    sessionToken, err2 := sessions.GetSessionToken(token)
    if err2 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusUnauthorized,
            fmt.Sprintf(err2.Error()))
        return
    }

    if !sessions.IsValidToken(sessionToken) {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusUnauthorized,
            fmt.Sprintf("Session token is not valid"))
        return
    }

    params := mux.Vars(request)
    userId, _ := strconv.Atoi(params["userId"])

    user, err3 := users.GetUser(userId)

    if err3 != nil {
        exceptions.BuildErrorResponse(
            &response,
            http.StatusConflict,
            fmt.Sprintf("Error %s", err3.Error()))
        return
    }

    // todo: use constant
    // var key = "pk_test_qblFNYngBkEdjEZ16jxxoWSM"
    var secreteKey = "sk_test_26PHem9AhJZvU623DfE1x4sd"

    stripe.Key = secreteKey

    userEmail := user.Email

    // todo : calculate amount
    amount := 1000

    stripParameters := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(int64(amount)),
        Currency: stripe.String(string(stripe.CurrencyUSD)),
        PaymentMethodTypes: stripe.StringSlice([]string{
            "card",
        }),
        ReceiptEmail: stripe.String(userEmail),
    }

    pi, _ := paymentintent.New(stripParameters)

    // todo: confirm Stripe transaction was successful

    fmt.Printf("%s", pi)
}
