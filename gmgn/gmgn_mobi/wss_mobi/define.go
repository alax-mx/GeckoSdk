package wss_mobi

const (
	BaseWsMainURL = "wss://ws.gmgn.mobi/main_ws"
	BaseWsBgURL   = "wss://ws.gmgn.mobi/bg_ws"
)

const (
	CHANNEL_PING              string = "ping"
	CHANNEL_NEW_POOL_INFO     string = "new_pool_info"
	CHANNEL_TRENCHES_UPDATE   string = "trenches_update"
	CHANNEL_TOKEN_SOCIAL_INFO string = "token_social_info"
	CHANNEL_NEW_LAUNCHED_INFO string = "new_launched_info"

	ACTION_SUBSCRIBE string = "subscribe"
	ACTION_HEARTBEAT string = "heartbeat"
)

type WsHandler func(message []byte)

type WSSRecvInfo struct {
	Channel string `json:"channel"`
	Data    any    `json:"data"`
}

type STChain struct {
	Chain string `json:"chain"`
}

type STHeartBeatMessage struct {
	ID      string `json:"id"`
	Action  string `json:"action"`
	Channel string `json:"channel"`
}

type WSSMessage struct {
	ID      string    `json:"id"`
	Action  string    `json:"action"`
	Channel string    `json:"channel"`
	Data    []STChain `json:"data"`
}

// {
// 	"c": "sol",
// 	"rg": "3",
// 	"p": [
// 	{
// 		"id": null,
// 		"a": "FbbBiRLf6gtvi4DQiR3L48ArtUQLv53jd5CNWjcUwP97",
// 		"ex": "pump",
// 		"pa": "FbbBiRLf6gtvi4DQiR3L48ArtUQLv53jd5CNWjcUwP97",
// 		"ba": "3SvcfiXgitnuF9dtfApfhktXau5NyZPVNQGHXhy8pump",
// 		"qa": "So11111111111111111111111111111111111111112",
// 		"qr": "0.00109968",
// 		"il": "0.1939175712",
// 		"iqr": "0.00109968",
// 		"l": "Pump.fun",
// 		"lpp": "pump_mayhem",
// 		"p_oc": false,
// 		"ot": 1778127937,
// 		"pts": "pump",
// 		"pt": 1,
// 		"qs": "WSOL",
// 		"bti": {
// 		"s": "BOOBS",
// 		"n": "BOOBS",
// 		"l": "",
// 		"ts": 1000000000,
// 		"v24h": 1.74162950445,
// 		"s24h": 1,
// 		"p": 0.000002466767719497,
// 		"hc": 1,
// 		"br": "0.0000",
// 		"bs": "burn",
// 		"isa": false,
// 		"hl": 0,
// 		"lqdt": "1.3769262024",
// 		"t10hr": 0,
// 		"t70_shr": null,
// 		"rm": 1,
// 		"rfa": 1,
// 		"mc": 2466.767719497,
// 		"ctr": "5pw4KmNLrAbyQf2AcDgM8ymYf1Buf6p9XTE5pQ3ohkae",
// 		"cbr": 0,
// 		"cts": "creator_hold",
// 		"d_ff": "",
// 		"rtar": 0,
// 		"bop": 0,
// 		"sdc": 0,
// 		"rc": 0,
// 		"pg": 0.00089,
// 		"dx_ad": 0,
// 		"dur": 0,
// 		"dx_ul": 0,
// 		"snp": 1,
// 		"f_tt": "",
// 		"cto": false,
// 		"d_cic": 0,
// 		"d_coc": 0,
// 		"d_ccc": 0,
// 		"dx_tb": false,
// 		"dx_bf": 0,
// 		"tg_cc": 0,
// 		"ni": "",
// 		"nt": "",
// 		"ntg": "",
// 		"nws": "",
// 		"dhr": "",
// 		"ihr": "",
// 		"p_oc": false,
// 		"c_t": "",
// 		"ts_n": "",
// 		"ts_s": "",
// 		"s_ihp": "",
// 		"s_os": "",
// 		"s_or": "",
// 		"f_ps": null,
// 		"bc": null,
// 		"s_bt": null,
// 		"s_st": null,
// 		"std": "2022"
// 		},
// 		"sid": "0041810882900995020"
// 	}
// 	]
// }
