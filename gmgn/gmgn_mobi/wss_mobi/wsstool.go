package wss_mobi

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/alax-mx/geckosdk/baseutils"
)

type WSSTool struct {
	chain                  string
	baseUrl                string
	baseParam              string
	Conn                   *websocket.Conn
	newPoolInfoHandler     WsHandler
	trenchesUpdateHandler  WsHandler
	tokenSocialInfoHandler WsHandler
	newLaunchedInfoHandler WsHandler
}

func NewWSSTool(baseUrl string, baseParam string) *WSSTool {
	return &WSSTool{
		chain:                  "",
		baseUrl:                baseUrl,
		baseParam:              baseParam,
		Conn:                   nil,
		newPoolInfoHandler:     nil,
		trenchesUpdateHandler:  nil,
		tokenSocialInfoHandler: nil,
		newLaunchedInfoHandler: nil,
	}
}

func (t *WSSTool) Start(chain string) error {
	t.chain = chain
	err := t.Connect()
	if err != nil {
		return err
	}

	go t.Recv()          // 启动接收消息的 goroutine
	go t.SendHeartbeat() // 启动发送心跳的 goroutine
	return nil
}

func (t *WSSTool) Connect() error {
	dialer := &websocket.Dialer{
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

	uuid := baseutils.RandomSmallHexString(16)
	url := t.baseUrl + "?uuid=" + uuid + "&" + t.baseParam // 必须是 wss://
	// 如果你需要带 token
	header := http.Header{}
	header.Add("User-Agent", "ReactNative")

	conn, resp, err := dialer.Dial(url, header)
	if err != nil {
		if resp != nil {
			return errors.New("wss 连接失败 HTTP 状态码: " + string(rune(resp.StatusCode)))
		} else {
			return errors.New("wss 连接失败 且没有 HTTP 响应")
		}
	}
	t.Conn = conn
	return nil
}

func (t *WSSTool) Close() {
	if t.Conn != nil {
		t.Conn.Close()
	}
}

func (t *WSSTool) Recv() {
	for {
		var msg map[string]any
		err := t.Conn.ReadJSON(&msg)
		if err != nil {
			return
		}

		data, _ := json.MarshalIndent(msg, "", "  ")
		WSSRecvInfo := WSSRecvInfo{}
		json.Unmarshal(data, &WSSRecvInfo)
		// 处理 new_pool_info 频道的消息
		if WSSRecvInfo.Channel == CHANNEL_NEW_POOL_INFO {
			if t.newPoolInfoHandler != nil {
				poolInfo, _ := json.MarshalIndent(WSSRecvInfo.Data, "", "  ")
				t.newPoolInfoHandler(poolInfo)
			}
		}

		// 处理 trenches_update 频道的消息
		if WSSRecvInfo.Channel == CHANNEL_TRENCHES_UPDATE {
			if t.trenchesUpdateHandler != nil {
				trenchesInfo, _ := json.MarshalIndent(WSSRecvInfo.Data, "", "  ")
				t.trenchesUpdateHandler(trenchesInfo)
			}
		}

		// 处理 token_social_info 频道的消息
		if WSSRecvInfo.Channel == CHANNEL_TOKEN_SOCIAL_INFO {
			if t.tokenSocialInfoHandler != nil {
				socialInfo, _ := json.MarshalIndent(WSSRecvInfo.Data, "", "  ")
				t.tokenSocialInfoHandler(socialInfo)
			}
		}

		// 处理 new_launched_info 频道的消息
		if WSSRecvInfo.Channel == CHANNEL_NEW_LAUNCHED_INFO {
			if t.newLaunchedInfoHandler != nil {
				launchedInfo, _ := json.MarshalIndent(WSSRecvInfo.Data, "", "  ")
				t.newLaunchedInfoHandler(launchedInfo)
			}
		}
	}
}

// Subscribe 订阅 new_pool_info 频道
func (t *WSSTool) Subscribe(channel string, handler WsHandler) error {
	if t.Conn == nil {
		return errors.New("WebSocket 连接未建立")
	}

	switch channel {
	case CHANNEL_NEW_POOL_INFO:
		t.newPoolInfoHandler = handler
	case CHANNEL_TRENCHES_UPDATE:
		t.trenchesUpdateHandler = handler
	case CHANNEL_TOKEN_SOCIAL_INFO:
		t.tokenSocialInfoHandler = handler
	case CHANNEL_NEW_LAUNCHED_INFO:
		t.newLaunchedInfoHandler = handler
	default:
		return errors.New("不支持的频道")
	}

	msg := &WSSMessage{
		ID:      baseutils.RandomSmallHexString(16),
		Action:  ACTION_SUBSCRIBE,
		Channel: channel,
		Data: []STChain{
			{
				Chain: t.chain,
			},
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return t.Conn.WriteMessage(websocket.TextMessage, data)
}

// SendHeartbeat 定期发送心跳消息，保持连接活跃
func (t *WSSTool) SendHeartbeat() error {
	if t.Conn == nil {
		return errors.New("WebSocket 连接未建立")
	}

	for {
		time.Sleep(10 * time.Second)
		heartbeat := STHeartBeatMessage{
			ID:      baseutils.RandomSmallHexString(16),
			Action:  ACTION_HEARTBEAT,
			Channel: CHANNEL_PING,
		}

		data, err := json.Marshal(heartbeat)
		if err != nil {
			return err
		}
		err = t.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return err
		}
	}
}
