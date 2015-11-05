package before

type order struct {
	pid int64
	cid int64
}

func CreateOrder(pid int64, cid int64) order {
	return order{
		pid: pid, cid: cid,
	}
}

func (o order) Submit() (int64, error) {
	// do some logic

	return int64(3252345234), nil
}
