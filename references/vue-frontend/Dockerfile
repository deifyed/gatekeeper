### Dependency stage
FROM node:lts-alpine AS dep-stage

WORKDIR /build

COPY package*.json ./

RUN npm ci

### Build stage
FROM node:lts-alpine AS build-stage

WORKDIR /build

COPY --from=dep-stage /build/node_modules ./node_modules

COPY --from=dep-stage /build/package*.json ./

COPY public/ ./public
COPY src/ ./src
COPY index.html .
COPY vite.config.js .

RUN npm run build

### Run
FROM nginx:stable-alpine AS production-stage

EXPOSE 80
VOLUME ["/usr/share/nginx/html", "/etc/nginx/conf.d"]

RUN ln -sf /dev/stdout /var/log/nginx/access.log && \
	ln -sf /dev/stderr /var/log/nginx/error.log

COPY --from=build-stage /build/dist /usr/share/nginx/html

