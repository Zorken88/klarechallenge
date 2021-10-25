FROM node:14-alpine as front

WORKDIR /app

COPY ./frontend/ ./

RUN npm install 
RUN npm run build

## Build

FROM golang:1.17-alpine as build

WORKDIR /app

COPY ./pricematic/go.mod ./
COPY ./pricematic/go.sum ./
RUN go mod download

COPY ./pricematic/ .

## Build over /
RUN go build -o /pricematic

## Deploy

FROM alpine:latest

WORKDIR /

## Get from Build and Front stage
COPY --from=build /pricematic /pricematic
COPY --from=build /app/data.json ./
COPY --from=front /app/build ./public/

EXPOSE 8080

CMD [ "/pricematic" ]