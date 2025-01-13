###############################################################################
# Build Static Website                                                        #
###############################################################################
FROM node:22.12-alpine AS node_builder

WORKDIR /app

# download and install production dependencies. if package.json and
# package-lock.json are unchanged, the cache will be used.
COPY ./app/package.json /app/package.json
COPY ./app/package-lock.json /app/package-lock.json
RUN npm ci --omit=dev

# copy application configuration files
COPY ./app/tsconfig.json /app/tsconfig.json
COPY ./app/astro.config.mjs /app/astro.config.mjs
COPY ./app/tailwind.config.mjs /app/tailwind.config.mjs

# copy application source code
COPY ./app/public /app/public
COPY ./app/src /app/src
COPY ./app/remove-inline.cjs /app/remove-inline.cjs

# build the static website
RUN npm run build

###############################################################################
# Static Website Hosting                                                      #
###############################################################################
FROM nginx:1.27-alpine

# add bash for env variable access (alpine does not ship with bash)
RUN apk update && apk add bash

# copy nginx configuration files
COPY ./nginx.conf /etc/nginx/nginx.conf

# copy static website files
COPY --chown=nginx:nginx --from=node_builder /app/dist /usr/share/nginx/html

EXPOSE 3000

# start nginx in background
CMD ["nginx", "-g", "daemon off;"]