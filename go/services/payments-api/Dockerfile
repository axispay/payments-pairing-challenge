FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod tidy || true
COPY . .
RUN go mod tidy && go build -o payments-api .

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/payments-api .
EXPOSE 3001
CMD ["./payments-api"]
