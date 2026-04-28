package gmgn_mobi_ws

import (
	"crypto/tls"
	"time"

	"github.com/gorilla/websocket"
)

var (
	// Endpoints
	BaseWsMainURL = "wss://ws.gmgn.mobi/ws"

	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 600
	// WebsocketPongTimeout is an interval for sending a PONG frame in response to PING frame from server
	WebsocketPongTimeout = time.Second * 10
	// WebsocketPingTimeout is an interval for waiting for a PONG response after sending a PING framer
	WebsocketPingTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	// WebsocketTimeoutReadWriteConnection is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	// using for websocket API (read/write)
	WebsocketTimeoutReadWriteConnection = time.Second * 10
	ProxyUrl                            = ""
)

func GetDialer() websocket.Dialer {
	return websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "ws.gmgn.mobi",
			SessionTicketsDisabled: true,
		},
	}
}

func GetWsBaseURL() string {
	return BaseWsMainURL
}
