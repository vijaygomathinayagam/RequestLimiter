FROM golang

WORKDIR /go/src/github.com/vijaygomathinayagam/RequestLimiter/

COPY . .

RUN go get ./
RUN go install

CMD [ "/go/bin/RequestLimiter" ]
