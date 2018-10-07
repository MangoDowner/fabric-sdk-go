package gm

import (
	"testing"
	"fmt"
	"log"
	"reflect"
     "github.com/tjfoc/gmsm/sm4"

)

func TestSM4Encrypt(t *testing.T) {
	key := []byte("1234567890abcdef")
	fmt.Printf("key = %v\n", key)
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	sm4.WriteKeyToPem("key.pem", key, nil)
	key, err := sm4.ReadKeyFromPem("key.pem", nil)
	fmt.Printf("key = %v\n", key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data = %x\n", data)
	c, err := sm4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	d0 := make([]byte, 16)
	c.Encrypt(d0, data)
	fmt.Printf("d0 = %x\n", d0)
	d1 := make([]byte, 16)
	c.Decrypt(d1, d0)
	fmt.Printf("d1 = %x\n", d1)
	if sa := testCompare(data, d1); sa != true {
		fmt.Printf("Error data!")
	}

}

func testCompare(key1, key2 []byte) bool {
	if len(key1) != len(key2) {
		return false
	}
	for i, v := range key1 {
		if i == 1 {
			fmt.Println("type of v", reflect.TypeOf(v))
		}
		a := key2[i]
		if a != v {
			return false
		}
	}
	return true
}

