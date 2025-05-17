FROM node:22-alpine

WORKDIR /app

COPY ./package.json /app/package.json
COPY ./ /app/

RUN yarn

CMD ["yarn", "start"]
