package v2

const (
	LedNose   = 0
	LedLeft   = 1
	LedMiddle = 2
	LedRight  = 3
	LedBottom = 4
)

func (n *NabConn) write(data string) error {
	_, err := n.conn.Write([]byte(data))
	return err
}
