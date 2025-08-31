# 🌍 StellarX – Operation Borderless

**A stablecoin-powered cross-border payment sandbox** built in **7 days**.  
Simulates real-time FX swaps, multi-currency wallets, and instant transfers across African and global currencies.

>

---

## 🛠️ Tech Stack

- **Backend**: Go (Gin, GORM)
- **Frontend**: React (coming soon)
- **Database**: PostgreSQL
- **Web Server**: nginx + Let’s Encrypt (HTTPS)
- **Deployment**: Render, Ubuntu 22.04 LTS(soon)
- **Infrastructure**: Docker Compose

---

## 🚀 How to Run Locally

Follow these steps to run **StellarX** on your machine.

### 1. Clone the Repository

```bash
git clone https://github.com/zekeriyyah/stellar-x.git
cd stellar-x
```

### 2. Set Environment Variables

```bash
cp .env.example .env
```

### 3. Install Dependencies

```bash
# Requires Docker and Docker Compose Plugin
sudo apt install -y docker.io docker-compose-plugin
```

### 4. Build and Run

```bash
# Start all services
docker compose up --build -d

# Verify services are running
docker compose ps
```

### 5. Test the API

```bash
curl http://localhost:port/api/v1/ping
```

✅ Expected: `{"message":"pong"}`

### 6. Access Services

- **App API**: `http://localhost:port/api/v1/ping`
- **pgAdmin**: `http://localhost:5050` (admin@stellar.com / admin123)

### 7. Stop Services

```bash
docker compose down
```

---

### ⚙️ Project Structure

```
stellarx/
├── cmd/
│   └── server/
│       └── main.go        # Entry point
├── internal/
│   ├── models/            # Database models
│   ├── repositories/        # DB access
│   ├── services/          # Business logic
│   ├── handlers/          # HTTP layer
│   └── database/          # DB connection
|   └── routes/            # Routes setup
├── config/
│   └── nginx.conf         # Reverse proxy config
├── docker-compose.yml     # Services orchestration
├── Dockerfile             # Go app container
└── .env.example           # Environment template
```

---

### 🧪 Testing with Postman

## Run live test with [postman](https://documenter.getpostman.com/view/29195129/2sB3HhsMoi)

### 🐳 Docker Compose Services

| Service   | Port | Purpose               |
| --------- | ---- | --------------------- |
| `app`     | 8080 | Go backend            |
| `db`      | 5040 | PostgreSQL            |
| `pgadmin` | 5050 | Database UI           |
| `nginx`   | 80   | Reverse proxy         |
| `certbot` | 443  | HTTPS (Let’s Encrypt) |

## 🔐 Admin Credentials

- **pgAdmin Email**: `admin@stellar.com`
- **pgAdmin Password**: `admin123`

---

## 🧭 Feature Walkthrough

<details>
<summary>📁 0. Health Check</summary>

<details>
<summary>✅ GET /ping</summary>

#### Request

```http
GET /ping
```

#### Response

```json
{
  "message": "pong"
}
```

✅ Confirms API is live and responsive

</details>

</details>

<details>
<summary>📁 1. Users Endpoints</summary>

<details>
<summary>✅ POST /api/signup</summary>

#### Request

```http
POST /api/signup
Content-Type: application/json
```

```json
{
  "email": "awwalcodestar@america.com",
  "phone": "+1870985847",
  "password": "@america"
}
```

#### Response

```json
{
  "message": "user registered successfully",
  "user": {
    "created_at": "2025-08-30T20:19:42.689537611Z",
    "email": "awwalcodestar@america.com",
    "id": 2
  }
}
```

✅ User created for wallet association

</details>
</details>

<details>
<summary>✅ GET /api/login</summary>
<details>
#### Request

```http
GET /api/login
Content-Type: application/json
```

```json
{
  "email": "awwalcodestar@nigeria.com",
  "password": "@nigeria"
}
```

#### Response

```json
{
  "message": "login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJwdXJwb3NlIjoiIiwiZXhwIjoxNzU2NjA0NDU0LCJpYXQiOjE3NTY1OTM2NTR9.VoeoJeZ93vgDb0Ssj7urHCOTWMOZRJcuQp16WlvdCDs",
  "user_id": 1
}
```

✅ User logged in.

</details>
</details>

<details>
<summary>✅ GET /api/v1/users/:userId</summary>
<details>
#### Request

```http
GET /api/v1/users/1
```

#### Response

```json
{
  "user": {
    "id": 1,
    "email": "awwalcodestar@nigeria.com",
    "phone": "+23470985847",
    "created_at": "2025-08-30T20:15:27.376762Z"
  }
}
```

✅ User details retrieved

</details>
</details>

<details>
<summary>📁 2. Wallet Creation</summary>

<details>
<summary>✅ POST /api/v1/wallet</summary>

#### Request

```http
POST /api/v1/wallet
Content-Type: application/json
```

```json
{
  "email": "awwalEUR@gmail.com",
  "label": "Nigeria Wallet"
}
```

#### Response

```json
{
  "message": "Wallet created successfully",
  "userId": 13,
  "email": "awwalEUR@gmail.com"
}
```

✅ Wallet initialized with zero balances for `cNGN`, `cXAF`, `USDx`, `EURx`

</details>
</details>

<details>
<summary>✅ GET /api/v1/wallet/:userId</summary>
<details>
#### Request

```http
GET /api/v1/wallet/1
```

#### Response

```json
 "body": [
        {
            "id": 1,
            "user_id": 1,
            "label": "Typical Nigerian Wallet",
            "created_at": "2025-08-30T20:29:17.274441Z",
            "updated_at": "2025-08-30T20:29:17.274441Z",
            "balances": [
                {
                    "id": 1,
                    "wallet_id": 1,
                    "currency": "cNGN",
                    "amount": 0,
                    "created_at": "2025-08-30T20:29:17.329304Z",
                    "updated_at": "2025-08-30T20:29:17.329304Z"
                },
                {
                    "id": 2,
                    "wallet_id": 1,
                    "currency": "cXAF",
                    "amount": 0,
                    "created_at": "2025-08-30T20:29:17.329304Z",
                    "updated_at": "2025-08-30T20:29:17.329304Z"
                },
                {
                    "id": 3,
                    "wallet_id": 1,
                    "currency": "USDx",
                    "amount": 0,
                    "created_at": "2025-08-30T20:29:17.329304Z",
                    "updated_at": "2025-08-30T20:29:17.329304Z"
                },
                {
                    "id": 4,
                    "wallet_id": 1,
                    "currency":
                    ...
```

✅ Confirms wallets and balances for a user

</details>

</details>

<details>
<summary>📁 3. Deposit</summary>

<details>
<summary>✅ POST /api/v1/deposit</summary>

#### Request

```http
POST /api/v1/deposit
Content-Type: application/json
```

```json
{
  "user_id": 14,
  "currency": "cNGN",
  "amount": 10000
}
```

#### Response

```json
{
  "message": "Deposit successful",
  "currency": "cNGN",
  "amount": 10000
}
```

✅ Balance updated instantly

</details>

</details>

<details>
<summary>📁 4. FX Swap</summary>

<details>
<summary>✅ POST /api/v1/swap</summary>
<details>
#### Request

```http
POST /api/v1/swap
Content-Type: application/json
```

```json
{
  "walletId": 7,
  "fromCurrency": "cNGN",
  "toCurrency": "USDx",
  "amount": 5000
}
```

#### Response

```json
{
  "message": "Swap successful",
  "transaction": {
    "tx_type": "swap",
    "from_currency": "cNGN",
    "to_currency": "USDx",
    "amount": 5000,
    "converted_amount": 3.33,
    "fx_rate": 0.000666,
    "status": "completed"
  }
}
```

✅ Used live FX rate from `api.frankfurter.dev`

</details>

</details>

<details>
<summary>📁 5. Transfer</summary>

<details>
<summary>✅ POST /api/v1/transfer</summary>

#### Request

```http
POST /api/v1/transfer
Content-Type: application/json
```

```json
{
  "sender_wallet_id": 7,
  "receiver_wallet_id": 8,
  "from_currency": "USDx",
  "to_currency": "cNGN",
  "amount": 100000
}
```

#### Response

```json
{
  "message": "Transfer successful",
  "transaction": {
    "tx_type": "transfer",
    "sender_wallet_id": 7,
    "receiver_wallet_id": 8,
    "from_currency": "USDx",
    "to_currency": "cNGN",
    "amount": 100000,
    "converted_amount": 150000000,
    "fx_rate": 1500,
    "status": "completed"
  }
}
```

✅ Auto-converted using FX rate; atomic transaction

</details>
</details>

<details>
<summary>📁 6. Transaction History</summary>

<details>
<summary>✅ GET /api/v1/transaction/:userId</summary>

#### Request

```http
GET /api/v1/transaction/14
```

#### Response

```json
{
  "userId": 14,
  "transactions": [
    {
      "tx_type": "deposit",
      "from_currency": "cNGN",
      "amount": 10000,
      "created_at": "2025-08-28T10:00:00Z"
    },
    {
      "tx_type": "swap",
      "from_currency": "cNGN",
      "to_currency": "USDx",
      "amount": 5000,
      "fx_rate": 0.000666,
      "created_at": "2025-08-28T10:05:00Z"
    }
  ]
}
```

✅ Chronological order; includes FX rates

</details>
</details>

<details>
<summary>📁 7. Compliance Mode</summary>

<details>
#### Request

```http
GET /api/v1/transaction/14
```

<details>
<summary>✅ GET /api/v1/audit/:userId </summary>

#### Response

```json
{
  "transactions": [
    {
      "id": 4,
      "tx_type": "transfer",
      "sender_wallet_id": 1,
      "receiver_wallet_id": 2,
      "from_currency": "USDx",
      "to_currency": "EURx",
      "amount": 500,
      "converted_amount": 428.89,
      "fx_rate": 0.85778,
      "status": "success",
      "created_at": "2025-08-31T09:53:28.713863Z"
    },
    {
      "id": 3,
      "tx_type": "swap",
      "sender_wallet_id": 1,
      "receiver_wallet_id": null,
      "from_currency": "USDx",
      "to_currency": "cNGN",
      "amount": 3000,
      "converted_amount": 4595923.984740929,
      "fx_rate": 1531.9746615803094,
      "status": "success",
      "created_at": "2025-08-31T09:21:20.928556Z"
    }, ...
```

> ✅ Audit logging middleware is implemented and ready to capture:
>
> - IP Address
> - Device
> - Browser
> - Country

</details>

</details>

<details>
<summary>📁 8. AI Assistant</summary>

<details>
<summary>✅ GET /api/v1/ask?q=what is the latest most stable coin</summary>

#### Request

```http
GET /api/v1/ask?q=what+is+the+latest+most+stable+coin
```

#### Response

```json
{
  "query": "what is the latest most stable coin",
  "answer": "Among the stablecoins in this system (cNGN, cXAF, USDx, EURx), USDx is typically the most stable as it's pegged 1:1 to the US Dollar."
}
```

✅ Powered by OpenAI, grounded in real FX data

</details>
</details>
---
---

## 🌐 Deployed Link

[https://stellar-x.onrender.com](https://stellar-x.onrender.com/)

---

## 📚 API Documentation

## View interactive API docs: [Postman Doc](https://documenter.getpostman.com/view/29195129/2sB3HhsMoi)

## 🗺️ Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    USER ||--o{ WALLET : "1:N"
    WALLET ||--o{ BALANCE : "1:N"
    WALLET ||--o{ TRANSACTION : "1:N sender"
    WALLET ||--o{ TRANSACTION : "1:N receiver"
    USER ||--o{ AUDIT_LOG : "1:N"
    WALLET ||--o{ AUDIT_LOG : "1:N (optional)"

    USER {
        uint id PK
        string email
        string phone
        string password
        datetime created_at
    }

    WALLET {
        uint id PK
        uint user_id FK
        string label
        datetime created_at
        datetime updated_at
    }

    BALANCE {
        uint id PK
        uint wallet_id FK
        string currency
        float amount
        datetime created_at
        datetime updated_at
    }

    TRANSACTION {
        uint id PK
        string tx_type
        uint sender_wallet_id FK
        uint receiver_wallet_id FK
        string from_currency
        string to_currency
        float amount
        float converted_amount
        float fx_rate
        string status
        datetime created_at
    }

    AUDIT_LOG {
        uint id PK
        uint user_id FK
        uint wallet_id FK
        string ip_address
        string device
        string browser
        string country
        string path
        string method
        datetime created_at
    }
```
