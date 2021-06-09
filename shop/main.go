package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()
	router.HandleOPTIONS = true
	router.GlobalOPTIONS = &options{}

	router.GET("/cakes", cakes)
	router.POST("/create-qr", createQR(wp))
	router.GET("/order-status/:merchantPaymentId", orderStatus(wp))

	log.Print("Server start")
	log.Fatal(http.ListenAndServe(":5000", router))
}

type options struct{}

func (o *options) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,DELETE,PUT,POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Max-Age", "3600")
}

func createQR(wp *paypayopa.WebPayment) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	type request struct {
		OrderItems []struct {
			Name      string `json:"name"`
			Category  string `json:"category"`
			Quantity  uint   `json:"quantity"`
			ProductID uint   `json:"productId"`
			UnitPrice struct {
				Amount   uint   `json:"amount"`
				Currency string `json:"currency"`
			} `json:"unitPrice"`
		} `json:"orderItems"`
		Amount struct {
			Amount   uint   `json:"amount"`
			Currency string `json:"currency"`
		} `json:"amount"`
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w)

		body := &request{}
		b, _ := ioutil.ReadAll(r.Body)

		if err := json.Unmarshal(b, body); err != nil {
			fmt.Fprint(w, `{ "errors": [{
				"code": 400,
				"message": "Bad Request"
			}] }`)

			w.WriteHeader(http.StatusBadRequest)

			return
		}

		merchantPaymentID := uuid.NewString()
		orders := make([]*paypayopa.MerchantOrderItem, 0, 0)

		for _, o := range body.OrderItems {
			orders = append(orders, &paypayopa.MerchantOrderItem{
				Name:      o.Name,
				Category:  o.Category,
				Quantity:  int(o.Quantity),
				ProductID: strconv.FormatUint(uint64(o.ProductID), 10),
				UnitPrice: &paypayopa.MoneyAmount{
					Amount:   int(o.UnitPrice.Amount),
					Currency: paypayopa.CurrencyJPY,
				},
			})
		}

		payload := &paypayopa.CreateQRCodePayload{
			MerchantPaymentID: merchantPaymentID,
			Amount: &paypayopa.MoneyAmount{
				Amount:   int(body.Amount.Amount),
				Currency: paypayopa.CurrencyJPY,
			},
			OrderItems:   orders,
			CodeType:     paypayopa.CodeTypeOrderQR,
			RequestedAt:  time.Now().Unix(),
			RedirectType: paypayopa.RedirectTypeWebLink,
			RedirectURL:  "http://localhost:8080/orderpayment/" + merchantPaymentID,
		}

		qrcode, info, err := wp.CreateQRCode(r.Context(), payload)
		if err != nil {
			fmt.Fprint(w, `{ "errors": [{
				"code": 500,
				"message": "Internal Server Error"
			}] }`)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		if !info.Success() {
			b, _ := json.Marshal(info)
			log.Println(string(b))
			fmt.Fprint(w, `{"resultInfo": `+string(b)+`}`)
			w.WriteHeader(info.StatusCode)

			return
		}

		b, _ = json.Marshal(qrcode)
		log.Println(string(b))
		fmt.Fprint(w, `{"data": `+string(b)+`}`)
		w.WriteHeader(http.StatusOK)
	}
}

func orderStatus(wp *paypayopa.WebPayment) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w)

		payment, info, err := wp.GetPaymentDetails(
			r.Context(),
			ps.ByName("merchantPaymentId"),
		)

		if err != nil {
			fmt.Fprint(w, `{ "errors": [{
				"code": 500,
				"message": "Internal Server Error"
			}] }`)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		if !info.Success() {
			b, _ := json.Marshal(info)
			log.Println(string(b))
			fmt.Fprint(w, `{"resultInfo": `+string(b)+`}`)
			w.WriteHeader(info.StatusCode)

			return
		}

		b, _ := json.Marshal(payment)
		log.Println(string(b))
		fmt.Fprint(w, `{"data": `+string(b)+`}`)
		w.WriteHeader(http.StatusOK)
	}
}

func cakes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setHeader(w)

	fmt.Fprint(w, `[{
		"title": "cake_shop.mississippi",
		"id": 1,
		"price": 120,
		"image": "darkforest.png"
	}, {
		"title": "cake_shop.red_velvet",
		"id": 2,
		"price": 190,
		"image": "redvelvet.png"
	}, {
		"title": "cake_shop.dark_forest",
		"id": 3,
		"price": 100,
		"image": "darkforestcake.png"
	}, {
		"title": "cake_shop.rainbow",
		"id": 4,
		"price": 200,
		"image": "rainbow.png"
	}, {
		"title": "cake_shop.lemon",
		"id": 5,
		"price": 80,
		"image": "lemon.png"
	}, {
		"title": "cake_shop.pineapple",
		"id": 6,
		"price": 110,
		"image": "pineapple.png"
	}, {
		"title": "cake_shop.banana",
		"id": 7,
		"price": 90,
		"image": "banana.png"
	}, {
		"title": "cake_shop.carrot",
		"id": 8,
		"price": 165,
		"image": "carrot.png"
	}, {
		"title": "cake_shop.choco",
		"id": 9,
		"price": 77,
		"image": "choco.png"
	}, {
		"title": "cake_shop.chocochip",
		"id": 10,
		"price": 130,
		"image": "chocochip.png"
	}, {
		"title": "cake_shop.orange",
		"id": 11,
		"price": 140,
		"image": "orange.png"
	}, {
		"title": "cake_shop.butterscotch",
		"id": 12,
		"price": 155,
		"image": "butterscotch.png"
	}]`)

	w.WriteHeader(http.StatusOK)
}
