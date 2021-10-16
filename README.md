# Golang Test Case REST API
Basic Rest Api

## Explanation
I've built a basic rest api project with a few best practices and the company's asked details.

## Features  
- (Required) Add Value with Key
- (Required) Remove Value by Key
- (Required) Flush All Keys and Values
- (Required) Get a Value by Key
- (Required) List All Values and Keys
- (Extra) Add a Game
- (Extra) Delete a Game
- (Extra) Flush All Games
- (Extra) List All Games
- (Extra) Get a Game

## Pre-requisites
1. (Required) `Golang 1.17`
2. (Recommended) `Docker` You need to use docker if you are not able to set env and running profiles in your IDE
2. (Recommended) `Docker Compose` You need to use docker compose if you are not able to set env and running profiles in your IDE
3. (Env - Required) `SAVE_INTERVAL` You have to set this env and it is expected integer and it represets seconds.
4. (Env - Optional) `PORT` You can set this env and it is expected integer and it represets running port, DEFAULT `8080`

## Run the application
If you have docker and docker-compose on your system, It is easy peasy lemon squeezy, 
1. `docker-compose up -d` It will build the project and will start the app to serve.

## Use the App
1. I uploaded Postman Collections and Env files, you can import them. So, every single endpoint will be set in your Postman.
To import them, follow this tutorial = https://learning.postman.com/docs/getting-started/importing-and-exporting-data/

## Endpoints (Required Features)

##### Add Value with Key
`curl --location --request GET 'localhost:8080/fake-redis/add?key=mehmet&value=canhoroz'`

##### Remove Value by Key
`curl --location --request GET 'localhost:8080/fake-redis/delete?key=mehmet'`

##### Flush All Keys and Values
`curl --location --request GET 'localhost:8080/fake-redis/flush'`

##### Get a Value by Key
`curl --location --request GET 'localhost:8080/fake-redis?key=mehmet'`

##### List All Values and Keys
`curl --location --request GET 'localhost:8080/fake-redis'`


## Endpoints (Extra Features)

##### Add a Game
`curl --location --request GET 'localhost:8080/games/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "game_title":"World of warcraft",
     "price": 1
}'`

##### Delete a Game
`curl --location --request GET 'localhost:8080/games/delete?id=1'`

##### Flush All Games
`curl --location --request GET 'localhost:8080/games/flush'`

##### List All Games
`curl --location --request GET 'localhost:8080/games?id=1'`

##### Get a Game
`curl --location --request GET 'localhost:8080/games'`

## 3rd Party Dependencies
As asked, there is no single 3rd party lib in the project. Everything was built in with native golang libs

## License
[MIT](https://github.com/mehmetcanhoroz/golang_test_case/blob/main/LICENSE)
