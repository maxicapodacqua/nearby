# nearby
Nearby API, to search nearby phone area codes

## Goal of this project
Trying to implement an API that will return data from a database without using any external go modules 
(or the least amount of modules). 

The idea is to work on a project from a scratch to get familiar with the golang ecosystem, development cycle and tooling 

## References and coding standards
* https://github.com/golang-standards/project-layout

## Tasks:
- [x] Connect to a database
- [x] Add unit tests
- [ ] Implement repository pattern
- [x] Define/implement endpoints
- [ ] Add auto generated swagger doc?

## Documentation
### Start project
Install all needed components with go get:
```shell
go get ./...
```
### Commands
Currently, this project has two different commands, 
one to start the API and another one to seed the API.
#### API
To start the API run:
```shell
go run cmd/nearby/main.go
```
This should start the API in port 8080, if you want to change the
port use set the environment var `PORT` to you desired port

##### Environment variables
* `PORT`: Changes the port where the API will run (default: `8080`)
* `SQLITE_DNS`: DNS to connect to a database (default: `data/db.sqlite`)

#### Seed
To run the database seeder:
```shell
go run cmd/seed/sqlite.go
```
This will insert rows in the only SQL table used by the API `nearby_area_codes`.
Currently, it searches for a `.sql` file called `./data/data.sql`.
The file should contain the `INSERT` statements to seed your database.
For example:
```sql
INSERT INTO `nearby_area_codes` VALUES (1,360,564)
```
For more information on the database schema check `data/schema.sql`

### Test
Run project tests with:
```shell
go test ./...
```

### Generate swagger doc
This project uses https://github.com/swaggo/swag to generate swagger file
```shell
swag init -g cmd/nearby/main.go --ot yaml -o api/
```
### Docker
You can build and run this project as a docker container.

**Make sure you already have your `./data/data.sql` file in place before building the image**

First you need to build the docker image:
```shell
docker build -t nearby-api .
```

And then run it in the exposed port:
```shell
docker run -p 8080:8080 nearby-api:latest
```

This will run the project in localhost:8080