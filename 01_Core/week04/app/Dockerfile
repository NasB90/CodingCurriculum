#use golang docker image
FROM golang:1.13
#Make directory that image will live in (given by golang docker img)
WORKDIR /go/src/app
#copy everthing in current folder to just defined directory
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


#go get pkg (when it is not a local pkg)
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

EXPOSE 8000

CMD ["app"]

#9min