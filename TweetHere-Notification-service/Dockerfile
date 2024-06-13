FROM golang:1.22-alpine AS build-stage
WORKDIR /noti_svc
COPY ./ /noti_svc 
RUN mkdir -p /noti_svc/build
RUN go mod download
RUN go build -v -o /noti_svc/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /noti_svc/build/api / 
COPY --from=build-stage /noti_svc/.env / 
EXPOSE 50057
CMD [ "/api" ]