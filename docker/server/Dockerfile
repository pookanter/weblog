FROM golang:1.24-alpine AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN apk update

COPY cmd ./
COPY api ./
COPY pkg ./
COPY go.mod ./
COPY build/api ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]

FROM debian:bookworm AS production

ARG timezone=Asia/Bangkok
ARG hotreload=false
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ $timezone

WORKDIR /app

RUN apt-get update && apt-get install

COPY --from=development /app/dist/authen .
COPY --from=development  /app/config config
COPY --from=development  /app/certs certs

RUN chmod +x authen

CMD ["./authen"]
