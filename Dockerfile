FROM golang:1.18.0 AS dev

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /bin/main ./cmd

# for Hot Reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin

CMD ["/bin/air"]

FROM gcr.io/distroless/base AS prd
COPY --from=build /bin/main /

CMD ["/main"]
