#############
# BUILDER
#############
FROM golang:1.17.6-alpine3.15 AS builder

ENV CGO_ENABLED=0

# Compile Delve
RUN apk add --no-cache git
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN mkdir /app
WORKDIR /app

 # <- COPY go.mod and go.sum files to the workspace
COPY app/go.mod .
COPY app/go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY app .


RUN go build -gcflags="all=-N -l" -o /server /app



#############
# RUNNER
#############
FROM golang:1.17.6-alpine3.15

ENV MIGRATION_PATH=/migration/db

COPY --from=builder /app/migration/db $MIGRATION_PATH
COPY --from=builder /go/bin/dlv /

COPY --from=builder /server /

EXPOSE 3000

# CMD ["/server"]
CMD ["/dlv", "--listen=:40000", "--headless=true", "--log", "--continue", "--accept-multiclient", "--api-version=2", "exec", "/server"]