package wrap_all_primitives

import (
	"fmt"

	"github.com/rafos/object-calisthenics-in-go/wrap_all_primitives/after"
	"github.com/rafos/object-calisthenics-in-go/wrap_all_primitives/before"
)

func WrapAllPrimitivesExample() {
	fmt.Println("*** WrapAllPrimitivesExample ***")
	pid := int64(1245)
	cid := int64(8877)

	beforeOrder := before.CreateOrder(pid, cid)
	beforeOrderId, err := beforeOrder.Submit()

	if err == nil {
		fmt.Println("BeforeOrderId =", beforeOrderId)
	}

	afterOrder := after.CreateOrder(pid, cid)
	afterOrderId, err := afterOrder.Submit()

	if err == nil {
		fmt.Println("AfterOrderId =", afterOrderId.String())
	}
}
