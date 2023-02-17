package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

type CheckoutData struct {
	ClientSecret string `json:"client_secret"`
}

func main() {
	// This is your test secret API key.
	stripe.Key = "sk_test_nfCbjWhpaU9QgNwa4c0Pq0Ua00OqaMJhYB"

	e := echo.New()

	e.POST("/secret", createCheckoutSession)

	log.Fatal(e.Start(":4242"))
}

func createCheckoutSession(c echo.Context) (err error) {
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"paynow",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("sgd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Carpool Service"),
					},
					UnitAmount: stripe.Int64(1600),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
	}

	result, _ := session.New(params)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, result.URL)
}
