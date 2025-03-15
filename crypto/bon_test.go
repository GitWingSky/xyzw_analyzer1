package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	//b, _ := hex.DecodeString("70782cc2202d2d2b494b4329282828282d2c4a474c512f00202b2d2e415c4d45614c29ea2328282d2e465d454a4d5a29292828282d2d41464c4d5029282828282d2b4b454c2d25415c4d457747584d4658494b432d2b5b4d5929202828282d2c5c41454d2a8fb260b3bd292828")
	b, _ := hex.DecodeString("7078366c626f6f69190f1b6b6b6a6a6a6f6e1e03070f685e510ef1ff6b6a6a6f690b09016b6a6a6a6a6f6e08050e136d4262696f6f03040e0f126b6a6a6a6a6f6c031e0f07230e6ba8616a6a6f6c041f07080f186b6b6a6a6a6f6909070e6f67031e0f0735051a0f041a0b0901")
	r, e := DecryptXAndDecode(b)
	if e != nil {
		panic(e)
	}
	fmt.Println(r)
	r2 := r.(map[string]interface{})
	body, ok := r2["body"].([]byte)
	if ok {
		body_r := DecodeFromBytes(body)
		r2["body"] = body_r
		fmt.Println(r2)
	}

}
func TestEncode(t *testing.T) {

	r := map[string]interface{}{
		"index":  0,
		"itemId": 3010,
		"number": 1,
	}
	fmt.Println(r)
	data := EncodeToBytes(r)

	fmt.Println(hex.EncodeToString(data))
	r2 := map[string]interface{}{
		"ack":  0,
		"body": data,
		"cmd":  "item_openpack",
		"seq":  0,
		"time": int64(1741969398878),
	}
	data2, e := EncodeAndEncryptX(r2)
	if e != nil {
		panic(e)
	}
	fmt.Println(hex.EncodeToString(data2))
}
