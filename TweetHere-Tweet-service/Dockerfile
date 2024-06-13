FROM golang:1.22-alpine AS build-stage
WORKDIR /tweet_svc
COPY ./ /tweet_svc 
RUN mkdir -p /tweet_svc/build
RUN go mod download
RUN go build -v -o /tweet_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /tweet_svc/build/api / 
COPY --from=build-stage /tweet_svc/.env / 
EXPOSE 50055
CMD [ "/api" ]