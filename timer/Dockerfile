FROM golang as build
WORKDIR /go/src/timer
COPY . /go/src/timer/
WORKDIR /go/src/timer/cmd/timer
RUN go build

FROM docker:20.10.12-git as runner
COPY --from=build  /go/src/timer/cmd/timer/timer /run/
WORKDIR run
CMD ["./timer"]