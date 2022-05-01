# base
FROM node:16.15.0-alpine3.14 AS base

RUN apk update && \
    apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ENV TZ=Asia/Tokyo

WORKDIR /app

# development
FROM base as dev

ENV PS1='\h /\W\ # '

CMD ["bash"]

# production
FROM base as builder

COPY . .

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock

RUN yarn && yarn build
