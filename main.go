package main

import (
	"fmt"
	"os"

	"github.com/pvormste/noplib"
)

func main() {
	os.Setenv("MY_SECRET_CREDENTIAL", "secret_cred")
	fmt.Println(noplib.CallToNopLib())
}
