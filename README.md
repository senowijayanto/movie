# Install via Docker

Clone the project

`git clone https://github.com/senowijayanto/movie.git`

Setup the project

`$ cd movie`

Run application via docker compose.

`$ docker compose up -d`

After all services is running, the Movie Backend App can be access on http://localhost:3000/ and the database can access on port 3307



# API

### List Movies

  - Method : GET

  - URL : http://localhost:3000/api/movies

  - Request :
   `curl --location --request GET 'http://localhost:3000/api/movies'`

  - Response :
  `{
     "code": 200,
     "status": "OK,
     "data": {
       ""
     }
   }`


