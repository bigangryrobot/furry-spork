FROM golang:alpine AS builder
COPY . .
RUN go mod download
# RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/filesystem

FROM scratch
COPY --from=builder /bin/filesystem /bin/filesystem
ENTRYPOINT ["/bin/filesystem"]