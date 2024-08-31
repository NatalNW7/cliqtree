FROM golang:1.23.0-alpine3.20 AS build
ENV CLIQTREE_ENV="development"
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o run cmd/cliqtree/main.go

FROM scratch
ENV CLIQTREE_ENV="development"
COPY --from=build  /app/run /
COPY --from=build /app/.env.development /
EXPOSE 8080

ENTRYPOINT [ "/run" ]