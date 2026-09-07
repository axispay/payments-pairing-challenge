FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod tidy || true
COPY . .
RUN go mod tidy && go build -o merchant-analytics .

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/merchant-analytics .
EXPOSE 3002
CMD ["./merchant-analytics"]
