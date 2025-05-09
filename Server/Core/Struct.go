package Core

type UserRole string

const (
	WsConn UserRole = ""
	Client UserRole = "client"
	Guest  UserRole = "guest"
	null   UserRole = ""
)

type Message struct {
	Sender string      `json:"sender"` //for server use only
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
}
type JoinJson struct {
	RoomID string `json:"roomId"`
	User   string `json:"user"`
}
type CloseJson struct {
	Password string `json:"password"`
}
type RoomData struct {
	Host     string   `json:"host"`
	RoomID   string   `json:"roomId"`
	Password string   `json:"password"`
	Clients  []string `json:"clients"`
}
type User struct {
	WsConn   string
	Role     UserRole
	RoomID   string
	Username string
	Meta     interface{}
}

type PageData struct {
	Title   string
	Content string
}
