package api

import (
	"github.com/gorilla/websocket"
	"github.com/kubenext/kubefun/internal/config"
	"net"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			host, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				return false
			}

			return shouldAllowHost(host, acceptedHosts())
		},
	}
)

func websocketService(manager ClientManager, dashConfig config.Dash) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(manager, dashConfig, w, r)
	}
}

func serveWebsocket(manager ClientManager, dashConfig config.Dash, w http.ResponseWriter, r *http.Request) {
	client, err := manager.ClientFromRequest(dashConfig, w, r)
	if err != nil {
		logger := dashConfig.Logger()
		logger.WithErr(err).Errorf("create websocket client")

	}

	go client.readPump()
	go client.writePump()
}
