package utils

import "hash/fnv"

func UserIDGenerator(email string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(email))
	return h.Sum32()
}
