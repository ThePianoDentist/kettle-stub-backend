FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/ThePianoDentist/kettle-stub-backend/
WORKDIR /go/src/github.com/ThePianoDentist/kettle-stub-backend
RUN go mod download
COPY . /go/src/github.com/ThePianoDentist/kettle-stub-backend
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/kettle-stub-backend github.com/ThePianoDentist/kettle-stub-backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/kettle-stub-backend /go/src/github.com/ThePianoDentist/kettle-stub-backend

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/ThePianoDentist/kettle-stub-backend/build/kettle-stub-backend /usr/bin/kettle-stub-backend
COPY --from=builder /go/src/github.com/ThePianoDentist/kettle-stub-backend/firebaseServiceAccountKey.json /opt/firebaseServiceAccountKey.json
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/kettle-stub-backend"]
