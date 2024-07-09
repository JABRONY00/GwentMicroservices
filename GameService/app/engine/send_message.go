package engine

func (pf *PlayerField) SendJson(resp interface{}) error {
	pf.Conn.Mut.Lock()
	defer pf.Conn.Mut.Unlock()
	err := pf.Conn.WriteJSON(resp)
	return err
}
