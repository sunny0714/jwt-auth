# Go REST API with JWT(JSON Web Token) Authentication in Gin Framework

## Necessary Go package list

- gin framework: 
`go get -u github.com/gin-gonic/gin`
- ORM library: 
`go get -u gorm.io/gorm`
- to authenticate and generate our JWT: 
`go get -u github.com/golang-jwt/jwt/v4`
- to help manage our environment variables: 
`go get -u github.com/joho/godotenv`
- to encrypt our user's password: 
`go get -u golang.org/x/crypto`
- to use mysql: 
`go get -u gorm.io/driver/mysql`