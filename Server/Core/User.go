package Core

func (p *User) Init(wsConn string) {
	p.WsConn = wsConn
	p.Role = null
	p.RoomID = ""
	p.Username = ""
	p.Meta = nil
}
func (p *User) SetRole(role UserRole) {
	p.Role = role
}
func (p *User) SetRoomID(roomID string) {
	p.RoomID = roomID
}
func (p *User) SetUsername(username string) {
	p.Username = username
}
func (p *User) SetMeta(meta interface{}) {
	p.Meta = meta
}
