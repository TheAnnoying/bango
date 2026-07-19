FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o bango

FROM scratch

COPY --from=build /app/bango /bango

ENTRYPOINT ["/bango"]