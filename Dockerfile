FROM golang:1.17 AS build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /bin/main ./cmd

# for Hot Reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin

CMD ["/bin/main"]

# FROM gcr.io/distroless/base
# COPY --from=build /bin/main /
# COPY --from=build /bin/air /

# CMD ["/main"]
