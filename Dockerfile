FROM golang:latest as build
RUN mkdir /app
WORKDIR /app
COPY . /app/
RUN go build

FROM centos as prod
RUN mkdir /app
WORKDIR /app
COPY --from=build /app/onibus-io-backend ./
COPY entrypoint.sh /app/
CMD ["./onibus-io-backend"]
