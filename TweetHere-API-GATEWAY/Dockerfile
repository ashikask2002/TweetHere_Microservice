FROM golang:1.22-alpine AS build-stage
WORKDIR /tweethere
COPY ./ /tweethere
RUN mkdir -p /tweethere/build
RUN go mod download
RUN go build -v -o /tweethere/build/api ./cmd/main.go


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /tweethere/build/api /
COPY --from=build-stage /tweethere/template/ /template/
COPY --from=build-stage /tweethere/static /static/
COPY --from=build-stage /tweethere/.env /
EXPOSE 5000
CMD [ "/api" ]