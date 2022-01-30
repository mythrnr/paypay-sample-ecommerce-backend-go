# Go SDK Implementation

[日本語](./README.jp.md)

- ⚠️ This sample is **Unofficial**. ⚠️
- Sample repository to use [PayPay Open Payment API SDK for Go](https://github.com/mythrnr/paypayopa-sdk-go).
- Server is ported from [paypay/paypay-sample-ecommerce-backend-php](https://github.com/paypay/paypay-sample-ecommerce-backend-php).
- Use repository [paypay/paypay-sample-ecommerce-frontend](https://github.com/paypay/paypay-sample-ecommerce-frontend) for Frontend.

## Install Requirements

```bash
go mod tidy
```

## Add API Keys to environment

```bash
export API_KEY="REPLACE_WITH_YOUR_API_KEY"
export API_SECRET="REPLACE_WITH_YOUR_SECRET_KEY" 
export MERCHID="REPLACE_WITH_YOUR_MERCHANT_ID"
```

## Run local Go server

```bash
make serve
```

You should now have the dev server running on http://localhost:5000

__⚠️NOTICE: If you run on Mac OS Monterey,
you need to change port or turn off AirPlay receiver.__

## CLI to Refund Operation

- Additional, this repository has command for refund.
- `Order ID` is picked from [Transaction List](https://developer.paypay.ne.jp/dashboard/transactions).

```bash
make refund

Enter Order ID issued by PayPay to refund: 000000000
Enter refund amount: 100
```

## Requirements

- Go 1.13 or above.
