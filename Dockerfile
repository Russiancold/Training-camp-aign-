FROM golang

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/TrainingCamp

ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

ADD . .

CMD ["go", "run", "main.go"]