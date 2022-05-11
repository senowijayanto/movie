# Movie Backend with Golang

## Software
* [Git](https://git-scm.com/downloads)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Installation
* Clone the project

```bash 
$ git clone https://github.com/senowijayanto/movie.git
```

* Setup the project

```bash
$ cd movie
```

* Run application via docker compose.

```bash
$ docker-compose up -d
```
*Note* : You have to wait about one minute after running the container using docker-compose, because the database service need more time to be running

## List of Container
- **backend**, Movie backend service container
- **db**, Database for movie backend

## List of Services
  - **List of Movies**

    * Method : GET

    * URL : `http://localhost:3000/api/movies`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/movies'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "title": "ALONE TRIP",
            "release_date": "2017-07-25",
            "rating": 10
          },
          {
            "title": "ACADEMY DINOSAUR",
            "release_date": "2017-01-20",
            "rating": 8
          },
          {
            "title": "ALABAMA DEVIL",
            "release_date": "2017-06-11",
            "rating": 7
          },
          {
            "title": "ADAPTATION HOLES",
            "release_date": "2017-04-06",
            "rating": 6
          },
          {
            "title": "CROOKED FROGMEN",
            "release_date": "2021-12-14",
            "rating": 0
          }
        ]
      }
      ```
  - **List of Movies With Pagination**

    * Method : GET

    * URL : `http://localhost:3000/api/movies?limit=5&offset=0`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/movies?limit=5&offset=0'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "title": "ALONE TRIP",
            "release_date": "2017-07-25",
            "rating": 10
          },
          {
            "title": "ACADEMY DINOSAUR",
            "release_date": "2017-01-20",
            "rating": 8
          },
          {
            "title": "ALABAMA DEVIL",
            "release_date": "2017-06-11",
            "rating": 7
          },
          {
            "title": "ADAPTATION HOLES",
            "release_date": "2017-04-06",
            "rating": 6
          },
          {
            "title": "AMADEUS HOLY",
            "release_date": "2017-08-13",
            "rating": 0
          }
        ]
      }
      ```
  - **List of Movies With Filter By Movie Title**

    * Method : GET

    * URL : `http://localhost:3000/api/movies?filter=ban`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/movies?filter=ban'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "title": "BANG KWAI",
            "release_date": "2019-12-22",
            "rating": 0
          },
          {
            "title": "BANGER PINOCCHIO",
            "release_date": "2020-01-25",
            "rating": 0
          }
        ]
      }
      ```
  - **List of Movies With Sort By Rating Average**

    * Method : GET

    * URL : `http://localhost:3000/api/movies?sortBy=average.desc`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/movies?sortBy=average.desc'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "title": "ALONE TRIP",
            "release_date": "2017-07-25",
            "rating": 10
          },
          {
            "title": "ACADEMY DINOSAUR",
            "release_date": "2017-01-20",
            "rating": 8
          },
          {
            "title": "ALABAMA DEVIL",
            "release_date": "2017-06-11",
            "rating": 7
          },
          {
            "title": "ADAPTATION HOLES",
            "release_date": "2017-04-06",
            "rating": 6
          },
          {
            "title": "CROOKED FROGMEN",
            "release_date": "2021-12-14",
            "rating": 0
          },
          {
            "title": "CLONES PINOCCHIO",
            "release_date": "2021-04-14",
            "rating": 0
          },
          {
            "title": "CHOCOLAT HARRY",
            "release_date": "2020-10-10",
            "rating": 0
          }
        ]
      }
      ```
  - **View the detail information of a Movie**

    * Method : GET

    * URL : `http://localhost:3000/api/movies/1`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/movies/1'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": {
          "id": 1,
          "title": "ACADEMY DINOSAUR",
          "release_date": "2017-01-20",
          "description": "A Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies"
        }
      }
      ```
  - **Rate Movie**

    * Method : POST

    * URL : `http://localhost:3000/api/ratings`

    * Request :
      ```bash
      curl --location --request POST 'http://localhost:3000/api/ratings' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "movie_id": 4,
            "rating": 10
        }'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": {
          "rating_id": 4,
          "movie_id": 4,
          "rating": 10
        }
      }
      ```
  - **List Actors**

    * Method : GET

    * URL : `http://localhost:3000/api/actors`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/actors'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "id": 1,
            "name": "Penelope Guiness"
          },
          {
            "id": 2,
            "name": "Nick Wahlberg"
          },
          {
            "id": 3,
            "name": "Ed Chase"
          }
        ]
      }
      ```
  - **List of Casts**

    * Method : GET

    * URL : `http://localhost:3000/api/casts`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/casts'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
            "actor_name": "Penelope Guiness",
            "movie_title": "ACADEMY DINOSAUR"
          },
          {
            "actor_name": "Penelope Guiness",
            "movie_title": "ANACONDA CONFESSIONS"
          },
          {
            "actor_name": "Penelope Guiness",
            "movie_title": "ANGELS LIFE"
          },
          {
            "actor_name": "Nick Wahlberg",
            "movie_title": "ADAPTATION HOLES"
          },
          {
            "actor_name": "Nick Wahlberg",
            "movie_title": "APACHE DIVINE"
          },
          {
            "actor_name": "Nick Wahlberg",
            "movie_title": "BABY HALL"
          },
          {
            "actor_name": "Ed Chase",
            "movie_title": "ALABAMA DEVIL"
          },
          {
            "actor_name": "Ed Chase",
            "movie_title": "ALONE TRIP"
          }
        ]
      }
      ```
  - **List of Casts With Pagination**

    * Method : GET

    * URL : `http://localhost:3000/api/casts?limit=5&offset=0`

    * Request :
      ```bash
      curl --location --request GET 'http://localhost:3000/api/casts?limit=5&offset=0'
      ```

    * Response :
      ```
      {
        "code": 200,
        "status": "OK",
        "data": [
          {
              "actor_name": "Penelope Guiness",
              "movie_title": "ACADEMY DINOSAUR"
          },
          {
              "actor_name": "Penelope Guiness",
              "movie_title": "ANACONDA CONFESSIONS"
          },
          {
              "actor_name": "Penelope Guiness",
              "movie_title": "ANGELS LIFE"
          },
          {
              "actor_name": "Nick Wahlberg",
              "movie_title": "ADAPTATION HOLES"
          },
          {
              "actor_name": "Nick Wahlberg",
              "movie_title": "APACHE DIVINE"
          }
        ]
      }
      ```
