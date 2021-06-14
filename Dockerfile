FROM golang AS build
ENV location /go/src/github.com/code-newbee/geeker
WORKDIR ${location}/server

ADD ./ ${location}/server

RUN go get -d ./...
RUN go install ./...

RUN CGO_ENABLED=0 go build -o /bin/geeker

FROM scratch

COPY --from=build /bin/geeker /bin/geeker

ENTRYPOINT ["/bin/geeker"]
EXPOSE 50001