# Book-API CRUD with PostgreSQL

## Table of contents 👀
* [General info](#general-info)
* [Technologies](#technologies)
* [Blog](#blog)
* [Setup](#setup)


### General info
BAPI or Book-API is a Golang REST API made to show some detail book for person. 

#### The BAPI Object 🍵
| Properties | Description | Type  |
|:----------- |:---------------|:--------|
|judul_buku| the identity title book | String| 
|penulis| author book | String |
|tgl_publikasi| date of publication | String | 


#### Routes ⚡
| Routes | HTTP Methods| Description
|:------- |:---------------|:--------------
| /api/book/     | GET                  | Displays all book
| /api/book/      | POST               | Creates a new book
|/api/book/{id}| GET     | Displays a specific book, given its name
|/api/book/{id}| PUT  | Update identitiy book
|/api/user/{id}}| DELETE | Deletes a specific book, given its id
	
### Technologies
Project is created with:

* Golang 
* gorilla/mux 
* lib/pq  
* joho/godotenv 
* PostgreSQL



### How I built it
👉 [Check out the series here!](https://berkaryasemampunya.medium.com/book-api-using-golang-and-postgresql-4870ff69989)


### Setup
To run this project locally, clone repo and add an `.env` file in the root:
```
POSTGRES_URL="Postgres connection string"
```

Then execute in command prompt:
```
$ cd go-postgres
$ go mod tidy
$ go run main.go
```

## API Reference

These are the endpoints available from the app

### `GET /api/book/`

Returns result identity

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "data": [
    {
     "judul_buku":"Al Wajiz",
     "penulis": "Imam Ibnu Taimiyah",
     "tgl_publikasi":"2020-10-11",
    }
  ]
}
```

</p>
</details>

---


### `POST /api/book/`

Creates a new identity book

#### Request 

This request requires body payload, you can find the example below.

<details><summary>Show example payload</summary>
<p>

```json
{
    "judul_buku":"Al Wajiz",
     "penulis": "Imam Ibnu Taimiyah",
     "tgl_publikasi":"2020-10-11",
}
```
</p>
</details>


### `GET /api/book/:id`

Returns a book by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": 1,
    "judul_buku":"Al Wajiz",
     "penulis": "Imam Ibnu Taimiyah",
     "tgl_publikasi":"2020-10-11",
  }
}
```

</p>
</details>


---

### `UPDATE /api/book/:ID`

Update value of identity book
	
#### Response

<details><summary>Show example response</summary>
<p>

```json
{
    "judul_buku":"Al Wajiz",
    "penulis": "Imam Ibnu Taimiyah",
    "tgl_publikasi":"2020-10-11", 
}
```

</p>
</details>


---
	
### `DELETE /api/book/:ID`

Delete team by id