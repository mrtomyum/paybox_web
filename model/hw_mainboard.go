package model

type MainBoard struct {
	machineId string `json:"machine_id"`
	Status    string
	Send      chan *Message
}

func (m *MainBoard) Event(c *Client) {
	switch c.Msg.Command {
	case "machine_id":    // ใช้สาหรับการร้องขอหมายเลข Serial Number ของ อุปกรณ์ Main Board
	case "set_ex_output": // สั่งงาน External Output ของ Main board
	case "get_ex_output": // ใช้สาหรับอ่านค่า External Input ของ Main board
	case "get_3g_status": // Event แจ้งสถานะของ Network
	}
}