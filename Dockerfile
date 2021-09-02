FROM golang AS buildStage

WORKDIR /thisdir
COPY . .
RUN CGO_ENABLED=0  go build -mod vendor -o /go/bin/ ./...
RUN ls -ltr /go/bin

FROM alpine

WORKDIR /app
COPY --from=buildStage /go/bin/* /app/theapp
COPY swagger/swagger.yaml ./swagger/
EXPOSE 8080
ENV BASE_URL=/api/v1
ENV SECTOR_ID=22
RUN adduser -D myuser
USER myuser
ENTRYPOINT ["/app/theapp"]
