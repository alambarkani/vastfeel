# syntax=docker/dockerfile:1

################################################################################
# Base stage using node image
FROM node:20.16.0-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

EXPOSE 8080

CMD [ "npm", "run", "preview" ]