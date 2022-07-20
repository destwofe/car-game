FROM golang:alpine AS builder
WORKDIR /app
ADD ./ /app
RUN cd /app && go build -o goapp

FROM golang:alpine
WORKDIR /app
COPY --from=builder /app/goapp /app

EXPOSE 4444
ENTRYPOINT ./goapp
