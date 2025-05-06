package Core

type Message struct {
	Sender string      `json:"sender"` //for server use only
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
}
type JoinJson struct {
	RoomID string `json:"roomId"`
	User   string `json:"user"`
}
type RoomData struct {
	Host     string   `json:"host"`
	RoomID   string   `json:"roomId"`
	Password string   `json:"password"`
	Clients  []string `json:"clients"`
}
