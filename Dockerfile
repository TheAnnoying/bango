FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go generate
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bango

FROM scratch

COPY --from=build /app/bango /bango

ENTRYPOINT ["/bango"]