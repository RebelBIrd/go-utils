package osutl

import (
	"os"
	"os/signal"
	"fmt"
)

func OsKill(closeAction func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <- c
	close(c)
	if sig == os.Interrupt || sig == os.Kill {
		fmt.Println("Close signal!")
		closeAction()
	}
}
