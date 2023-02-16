package random

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func randomStringUUID()string{
    return uuid.NewString()
}
func RandomHexStringFromUUID() string{
    uuidstring := randomStringUUID()
    src := []byte(uuidstring)
	encodedStr := hex.EncodeToString(src)
    return encodedStr
}

func randomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length+2)
    rand.Read(b)
    return fmt.Sprintf("%x", b)[2 : length+2]
}

func HexString(len int)string {
    randStr := randomString(len)
    src := []byte(randStr)
	encodedStr := hex.EncodeToString(src)
    return encodedStr
}

func HexString10()string{
    return HexString(10)
}