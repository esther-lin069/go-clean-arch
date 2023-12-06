help: # 顯示所有可用的指令
	@grep -E '^[a-zA-Z0-9 -|/]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

ifeq ($(MAKECMDGOALS), run/websocket/client)
    include ./build/websocket-client/env/local/websocketclient.mk
endif
ifeq ($(MAKECMDGOALS), run/websocket/server)
    include ./build/websocket-server/env/local/websocketserver.mk
endif

.PHONY: run/websocket/client
run/websocket/client: # 本地啟動 websocket client 服務
	@${LOCAL_ENV} go run cmd/main.go websocket-client

.PHONY: run/websocket/server
run/websocket/server: # 本地啟動 websocket server 服務
	@${LOCAL_ENV} go run cmd/main.go websocket-server