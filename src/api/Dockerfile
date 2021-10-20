FROM golang:1.16.6-alpine
WORKDIR /api/
ADD . /api
RUN cd /api && go build
EXPOSE 3000
ENTRYPOINT ["./api"]