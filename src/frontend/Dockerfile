FROM node:14.8.0-alpine3.11

WORKDIR /app

# adds the bin scripts inside node modules
ENV PATH /app/node_modules/.bin:$PATH

COPY package.json /app/package.json
RUN yarn
RUN yarn global add @vue/cli@3.7.0

CMD ["yarn", "serve"]