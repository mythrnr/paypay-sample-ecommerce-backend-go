package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mythrnr/paypayopa-sdk-go"
)

func main() {
	wp := paypayopa.NewWebPayment(
		paypayopa.NewCredentials(
			paypayopa.EnvSandbox,
			os.Getenv("API_KEY"),
			os.Getenv("API_SECRET"),
			os.Getenv("MERCHID"),
		),
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Order ID issued by PayPay to refund: ")
	scanner.Scan()
	paymentID := scanner.Text()

	fmt.Print("Enter refund amount: ")
	scanner.Scan()
	refundAmount, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	res, info, err := wp.RefundPayment(context.Background(), &paypayopa.RefundPaymentPayload{
		MerchantRefundID: uuid.NewString(),
		PaymentID:        paymentID,
		RequestedAt:      time.Now().Unix(),
		Amount: &paypayopa.MoneyAmount{
			Amount:   int(refundAmount),
			Currency: paypayopa.CurrencyJPY,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	{
		b, _ := json.MarshalIndent(info, "", "  ")
		log.Println("Result Info:", string(b))
	}

	{
		b, _ := json.MarshalIndent(res, "", "  ")
		log.Println("Refund Data:", string(b))
	}
}
