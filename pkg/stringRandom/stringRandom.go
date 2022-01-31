package stringrandom

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
	// b := make([]byte, n)
	// if _, err := rand.Read(b); err != nil {
	// 	panic(err)
	// }
	// return fmt.Sprintf("%X", b)
	// b := make([]byte, length)
	// for i := range b {
	// 	b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	// }
	// return string(b)
}
