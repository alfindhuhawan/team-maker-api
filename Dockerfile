################
# BUILD BINARY #
################

FROM golang:1.20.13-alpine3.18 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR $GOPATH/src/operation-dashboard-api
COPY . .

RUN echo $PWD && ls -lah

# Fetch dependencies.
# RUN go get -d -v
RUN go mod tidy
RUN go mod verify

# CMD go build -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o ./operation-dashboard-api .

#####################
# MAKE SMALL BINARY #
#####################
FROM alpine:latest

RUN apk update && apk add --no-cache tzdata
ENV TZ=UTC

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy the executable.
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/operation-dashboard-api/operation-dashboard-api /app
# COPY --from=builder /go/src/operation-dashboard-api/template/report.html /app/template/report.html