FROM node:13.8-alpine as build
WORKDIR /frontend
COPY . .
RUN yarn install

ENV HOST 0.0.0.0
EXPOSE 8080