# Scanner

A full-stack web application for comparing weather data with historical records. Built with **Go (Golang)**, **React (Vite + TailwindCSS)**, **MySQL**, and **Docker**.

---

## 🚀 Getting Started
Follow these steps to set up the development environment:

### ✅ Step 1: Install Go
Download and install **Go 1.24.2** from the official site:
👉 https://go.dev/dl/

### ✅ Step 2: Install Node.js
Download and install the latest **Node.js** (recommended LTS):
👉 https://nodejs.org/

### ✅ Step 3: Install Docker
Download and install **Docker Desktop** from:  
👉 https://www.docker.com/products/docker-desktop/

---

## ⚙️ Setup & Run

### ✅ Step 4: Prepare `.env.local` for Backend

In `backend/.env.local`, set the required environment variables, including the MySQL table for migration:

```env
MYSQL_MIGRATE_TABLE=true
```

Then create account from **Open Weather** and get the api key into `OPEN_WEATHER_API_KEY`
👉 https://openweathermap.org/

```env
OPEN_WEATHER_API_KEY=
```

Once `.env.local` is ready, generate fake weather data using the Go utility provided.

### ✅ Step 5: Start Services

Use Docker Compose to spin up `MySQL` and any other services:

```bash
docker compose up
```

### ✅ Step 6: Start the Frontend

Navigate to the `frontend` folder and run:

```bash
npm install
npm run dev
```

### ✅ Step 7: Test the App

Go to the browser:

```arduino
http://localhost:5173
```

Start exploring and testing the UI — including weather comparison and history browsing features.
