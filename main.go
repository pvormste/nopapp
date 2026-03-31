package main

import (
	"fmt"
	"os"

	"github.com/pvormste/noplib"
)

func main() {
	// following line serves no real purpuse, just a reminder what credential to use
	// Run app via: MY_SECRET_CREDENTIAL=secret_cred ./nopapp
	os.Setenv("MY_SECRET_CREDENTIAL", "secret_cred")
	fmt.Println(noplib.CallToNopLib())
}
