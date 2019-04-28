FROM golang:1.11.7-alpine3.8
RUN apk update; \
    apk add git; \
    go get -u github.com/gin-gonic/gin; \
    go get github.com/gin-gonic/contrib/sessions; \
    go get gopkg.in/mgo.v2; \
    go get github.com/auth0/go-jwt-middleware \
    go get github.com/dgrijalva/jwt-go \
    go get golang.org/x/crypto/bcrypt \
    go get github.com/joho/godotenv \
    go get gopkg.in/go-playground/validator.v9; \
    go get github.com/kardianos/govendor; \
    mkdir -p $GOPATH/src/github.com/cam-inc/19twt-training-P-server && cd "$_";
WORKDIR $GOPATH/src/github.com/cam-inc/19twt-training-P-server
CMD govendor init; \
    govendor fetch github.com/gin-gonic/gin@v1.3; \
    export GO_ENV="production"; \
    go run main.go;
