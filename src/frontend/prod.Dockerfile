# build environment
FROM node:14.8.0-alpine3.11
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json /app/package.json
RUN yarn --silent
RUN yarn global add @vue/cli@3.7.0
COPY . /app
RUN yarn build

# production environment
FROM nginx:1.16.0-alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
