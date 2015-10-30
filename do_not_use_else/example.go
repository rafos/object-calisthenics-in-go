package do_not_use_else

import (
	"github.com/rafos/object-calisthenics-in-go/do_not_use_else/before"
	"github.com/rafos/object-calisthenics-in-go/do_not_use_else/after"
)

func DoNotUseElseExample() {
	beforeLoginService := before.NewLoginService()
	beforeLoginService.Login("rafos", "hapacz")

	afterLoginService := after.NewLoginService()
	afterLoginService.Login("rafos", "hapacz")
}
