# Goly a url shortener like Bitly made with Go!

## Stack:
- Go 1.18
- PostgreSQL with docker
- Fiber web framework (inspired in Express)
- GORM (orm for Go)

<br/>

## First steps:
- Run docker
```sh
docker run --name postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:14
```

- Initialize Go project
```sh
go mod init <project-name>
```
<br/>

## Routes:
| Method | Endpoint       | Description        |
|--------|----------------|--------------------|
|  GET   |  /r/:redirect  | Redirect to url    |
|  GET   |  /goly         | Get all golys      |
|  GET   |  /goly/:id     | Get goly by id     |
|  POST  |  /goly         | Create a goly      |
|  PUT   |  /goly         | Update a goly      |
| DELETE |  /goly/:id     | Delete a goly      |

<br/>

## Body and Responses
### Create goly
- Body
```json
{
	"goly": "",
	"redirect": "https://marciojrdev.com",
	"random": true
}
```
Properties
> Goly: Unique name of goly (example: "goly.com/r/name-of-my-goly")

> Redirect: Url to redirect.

> Random: if true will generate a random goly (example: "goly.com/r/rfBd56ti").

<br/>

- Response
```json
{
  "id": 1,
  "redirect": "https://marciojrdev.com",
  "goly": "rfBd56ti",
  "clicked": 0,
  "random": true
}
```
Properties
> Id: id of goly

> Redirect: Url to redirect.

> Goly: Unique name of goly.
 >> In this case was *"rfBd56ti"*, so the link will be *"goly.com/r/rfBd56ti"*

> Clicked: Number of clicks the link received.

> Random: if true will generate a random goly (example: "goly.com/r/rfBd56ti").
