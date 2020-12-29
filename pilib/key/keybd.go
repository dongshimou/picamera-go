package key

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

var(
	keyboard *keybd_event.KeyBonding
)

func init(){
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	keyboard=&kb
}

func Key()*keybd_event.KeyBonding{
	return keyboard
}
