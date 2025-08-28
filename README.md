Based on your **Insomnia collection** and the **Operation Borderless** mission, here's a **professional, two-level collapsible `README.md`** that mirrors your folder structure and includes real test data.

This README will:

- ✅ Use your exact endpoint names and requests
- ✅ Show real payloads from your tests
- ✅ Be GitHub-compatible with `<details>` and `<summary>`
- ✅ Prove end-to-end functionality

---

````markdown
# 🌍 StellarX – Operation Borderless

**A stablecoin-powered cross-border payment sandbox** built in **7 days**.  
Simulates real-time FX swaps, multi-currency wallets, and instant transfers across African and global currencies.

> "The previous applicant took 14 days. You're doing it in 7."

---

## 🛠️ Tech Stack

- **Backend**: Go (Gin)
- **Frontend**: React (coming soon)
- **Database**: PostgreSQL
- **Web Server**: nginx + Let’s Encrypt (HTTPS)
- **Deployment**: Ubuntu 22.04 LTS
- **Infrastructure**: Docker Compose

---

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
````

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
<summary>✅ POST /api/v1/users</summary>

#### Request

```http
POST /api/v1/users
Content-Type: application/json
```

```json
{
  "email": "awwalEUR@gmail.com",
  "phone": "+237670045009"
}
```

#### Response

```json
{
  "message": "User created successfully",
  "userId": 12,
  "email": "awwalEUR@gmail.com",
  "phone": "+237670045009"
}
```

✅ User created for wallet association

</details>

<details>
<summary>✅ GET /api/v1/users/:userId</summary>

#### Request

```http
GET /api/v1/users/12
```

#### Response

```json
{
  "user": {
    "id": 12,
    "email": "awwalEUR@gmail.com",
    "phone": "+237670045009",
    "created_at": "2025-08-28T10:00:00Z"
  }
}
```

✅ User details retrieved

</details>

<details>
<summary>✅ GET /api/v1/users/email/:email</summary>

#### Request

```http
GET /api/v1/users/email/awwalEUR@gmail.com
```

#### Response

```json
{
  "user": {
    "id": 12,
    "email": "awwalEUR@gmail.com",
    "phone": "+237670045009"
  }
}
```

✅ User retrieved by email

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

<details>
<summary>✅ GET /api/v1/wallet/:userId</summary>

#### Request

```http
GET /api/v1/wallet/13
```

#### Response

```json
{
  "wallet": {
    "id": 13,
    "user_id": 13,
    "label": "Nigeria Wallet",
    "balances": [
      { "currency": "cNGN", "amount": 10000 },
      { "currency": "USDx", "amount": 6.67 }
    ]
  }
}
```

✅ Confirms wallet and balances

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
<summary>✅ GET /api/v1/audit/:userId (Coming Soon)</summary>

> ✅ Audit logging middleware is implemented and ready to capture:
>
> - IP Address
> - Device
> - Browser
> - Country
>
> Will be activated in production deployment.

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

## 🌐 Deployed Link

[https://stellarx.example.com](https://stellarx.example.com)

---

## 📚 API Documentation

View interactive API docs:  
👉 [Insomnia Workspace](https://insomnia.rest/docs/your-link)

---

**Prototype Name**: **StellarX**  
**Mission**: **Operation Borderless**  
**Built in**: **7 days**

```


```
