package helpers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenUUID() string {
	// data := []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "5", "3", "1", "2", "9", "8", "7", "4", "6", "0", "!", "@", "#", "$", "^", "&", "*", "(", ")", "-", "_", "+", "=", "[", "]", "{", "}", ";", ":", "'", "/", "?", "<", ">", ".", ","}

	// result := ""
	// rand.Seed(time.Now().UnixNano())
	// list := rand.Perm(len(data))
	// no := 1
	// for _, v := range list {
	// 	if no <= 32 {
	// 		result += data[v].(string)
	// 	}
	// 	no++
	// }
	id := uuid.New()
	hasher := md5.New()
	hasher.Write([]byte(id.String()))
	return hex.EncodeToString(hasher.Sum(nil))

}
