FROM --platform=linux/amd64 golang:1.15 as builder
WORKDIR /backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o backend-api
# RUN CGO_ENABLED=0 GOOS=linux go build -v -o nest8-savety-api-dev #dev

FROM --platform=linux/amd64 alpine
COPY --from=builder /backend /backend
WORKDIR /backend

EXPOSE 5500
# EXPOSE 8902 #dev
CMD ["./backend-api"]
# CMD ["./nest8-savety-api-dev"] #dev