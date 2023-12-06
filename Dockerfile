# 階段一：build
# WORKDIR 統一使用 /app，執行檔案統一叫 main，放置到 /app/ 下
ARG IMAGE_DOMAIN
FROM ${IMAGE_DOMAIN}golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main cmd/main.go

# 階段二：打包
FROM ${IMAGE_DOMAIN}alpine:3.18

# 時區統一用 +08:00
ENV TZ=Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 複製打包好的執行檔案
COPY --from=builder /app/main  /app/main

# 複製 env 下面的 *.yaml 到目標下
COPY --from=builder /app/env/*.yaml  /app/env/

# 統一 WORKDIR 都使用APP
WORKDIR /app
ENTRYPOINT [ "/app/main" ]
