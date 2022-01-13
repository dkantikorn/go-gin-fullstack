## Go gin-gonic backend
mkdir go-gin-fullstack/backend 
cd go-gin-fullstack/backend
code .


## Go downlod for all libraries 
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
go get github.com/gin-contrib/cors 
go get github.com/gin-gonic/contrib
go get github.com/gin-gonic/gin
go get github.com/aws/aws-sdk-go 
go get github.com/sendgrid/sendgrid-go
go get github.com/stretchr/testify
go get github.com/twinj/uuid
go get github.com/matcornic/hermes/v2

## TEST
$ cd tests 
$ go test -v 
$ go test -v --run <test-case-name>
$ go test -v --run TestSaveUser

## TEST from application root directory
$ go test -v ./...

## TEST with Docker
$ docker-compose -f docker-compose.test.yml up --build

## Continuous Integration Tools / circleci tool  / Travis CI
$ mkdir .circleci 
$ cd .circleci 
$ touch config.yml
$ circleci local execute --job build #for run test circleci in local

## Badge in readme file
[![CircleCI](https://circleci.com/gh/<user_name>/<repo_name>.svg?style=svg)](https://circleci.com/gh/<user_name>/<repo_name>)