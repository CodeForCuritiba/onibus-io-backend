FROM golang:latest as build-stage

RUN mkdir /app
WORKDIR /app
COPY . /app/
RUN go build

FROM centos as prod
RUN mkdir /app
WORKDIR /app
COPY --from=build-stage /app/onibus-io-backend /app
COPY entrypoint.sh /app
CMD /app/entrypoint.sh
