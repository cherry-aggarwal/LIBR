package models

type UserMsg struct {
	Content   string `json:"content"`
	TimeStamp int64  `json:"ts"`
}

type ModSign struct {
	Content   string `json:"content"`
	TimeStamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

type ModResponse struct {
	Sign      string `json:"sign"`
	PublicKey string `json:"public_key"`
	Status    string `json:"status"`
}

type ModLogEntry struct {
	PublicKey string `json:"public_key"`
	Content   string `json:"content"`
	TimeStamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

type ModConfig struct {
	Forbidden  []string `json:"forbidden"`
	Thresholds string   `json:"thresholds"`
}

type Category struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

type MsgCert struct {
	PublicKey string    `json:"public_key"`
	Msg       Msg       `json:"msg"`
	ModCerts  []ModCert `json:"mod_certs"`
	Sign      string    `json:"sign"`
	Reason    string    `json:"reason,omitempty"`
}

type Msg struct {
	Content string `json:"content"`
	Ts      int64  `json:"ts"`
}

type ModCert struct {
	Sign      string `json:"sign"`
	PublicKey string `json:"public_key"`
	Status    string `json:"status"`
}
