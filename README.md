# Go URL Shortener (Backend Only)

A lightweight **in-memory URL shortener** built with **Go**.  
This project is **backend-only** — no database and no frontend. 

URLs are stored only **in memory and are not persistent**, so all data will be lost when the server restarts or the Render free instance sleeps.

---

## Endpoints

### 1. Root url https://go-urlshortner.onrender.com/
**GET /**  
Returns a simple message: hello world 

---

### 2. Shorten a URL
**POST /shorten** (POSTMAN) 

**Full URL:**  https://go-urlshortner.onrender.com/shorten


**Request Body (JSON):**
```json
{
  "url": "https://xyz.com"
}
```
**Response (JSON) Like:**
```json
{
  "short_url": "d9736711"
}
```
Copy the generated hex ID and append it to this URL in place of **d9736711** in your browser
https://go-urlshortner.onrender.com/redirect/**d9736711**

