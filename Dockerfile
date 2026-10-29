FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /api .
FROM scratch
COPY --from=build /api /api
USER 65532:65532
ENV LISTEN_ADDR=0.0.0.0:8090
EXPOSE 8090
ENTRYPOINT ["/api"]
