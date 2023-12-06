help: # 顯示所有可用的指令
	@grep -E '^[a-zA-Z0-9 -|/]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

ifeq ($(MAKECMDGOALS), run/event)
    include ./build/websocket-client/env/local/websocketclient.mk
endif

.PHONY: run/event
run/event: # 本地啟動 event 服務
	@ ${LOCAL_ENV} go run cmd/main.go event