package do_not_use_else

import (
	"fmt"
	"github.com/rafos/object-calisthenics-in-go/do_not_use_else/after"
	"github.com/rafos/object-calisthenics-in-go/do_not_use_else/before"
)

func DoNotUseElseExample() {
	fmt.Println("*** DoNotUseElseExample ***")

	beforeLoginService := before.NewLoginService()
	beforeLoginService.Login("rafos", "hapacz")

	afterLoginService := after.NewLoginService()
	afterLoginService.Login("rafos", "hapacz")
}
