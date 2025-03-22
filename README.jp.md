# Go SDK Implementation

[English](./README.md)

- ⚠️ このサンプルは **非公式** . ⚠️
- [PayPay Open Payment API の Go 用 SDK](https://github.com/mythrnr/paypayopa-sdk-go) のサンプルリポジトリ.
- 実装は [paypay/paypay-sample-ecommerce-backend-php](https://github.com/paypay/paypay-sample-ecommerce-backend-php) を移植したもの.
- フロントエンドは [paypay/paypay-sample-ecommerce-frontend](https://github.com/paypay/paypay-sample-ecommerce-frontend) を利用する.

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

http://localhost:5000 で動作確認ができる.

**⚠️ NOTICE: Mac OS Monterey 以降は AirPlay receiver が 5000 番を使用する為,
ポート番号を変えるか AirPlay receiver をオフにする必要がある.**

## CLI to Refund Operation

- 返金用の CLI を追加した.
- `Order ID` は [取引一覧](https://developer.paypay.ne.jp/dashboard/transactions) にある `Order ID` を入力する.

```bash
make refund

Enter Order ID issued by PayPay to refund: 000000000
Enter refund amount: 100
```

## Requirements

- Go 1.22 以上.
- Docker (開発時)
