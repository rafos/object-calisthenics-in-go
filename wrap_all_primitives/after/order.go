package after

import (
	"strconv"
)

type order struct {
	pid productId
	cid customerId
}

type productId struct {
	id int64
}

// some methods on productId struct

type customerId struct {
	id int64
}

// some methods on customerId struct

type orderId struct {
	id int64
}

func (oid orderId) String() string {
	return strconv.FormatInt(oid.id, 10)
}

// some other methods on orderId struct

func CreateOrder(pid int64, cid int64) order {
	return order{
		pid: productId{pid}, cid: customerId{cid},
	}
}

func (o order) Submit() (orderId, error) {
	// do some logic

	return orderId{int64(3252345234)}, nil
}
