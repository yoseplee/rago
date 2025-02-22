package another

import (
	"fmt"
	"github.com/spf13/viper"
)

func Do() {
	getString := viper.GetString("openapi.api-key")
	fmt.Println(getString)
}
