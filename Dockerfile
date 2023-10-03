FROM golang:alpine AS totobuild
WORKDIR /go/src/bitbucket.org/isbtotogroup/frontend_svelte
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# ---- Svelte Base ----
FROM node:lts-alpine AS totosveltebase
WORKDIR /svelteapp
COPY [ "frontend/package.json" , "frontend/yarn.lock"  , "./"]

# ---- Svelte Dependencies ----
FROM totosveltebase AS totosveltedep
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM totosveltebase AS totosveltebuilder
COPY --from=totosveltedep /svelteapp/prod_node_modules ./node_modules
COPY ./frontend .
RUN yarn build

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
COPY --from=totosveltebuilder /svelteapp/dist ./frontend/dist
COPY --from=totobuild /go/src/bitbucket.org/isbtotogroup/frontend_svelte/app .
COPY --from=totobuild /go/src/bitbucket.org/isbtotogroup/frontend_svelte/env-sample /app/.env
EXPOSE 2116
CMD ["./app"]