
FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/testproject .

FROM alpine:edge

COPY --from=build /tmp/testproject /app/testproject

CMD /app/testproject

