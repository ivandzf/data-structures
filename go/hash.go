package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"

	log "github.com/sirupsen/logrus"
)

type HashKey []byte

func (h HashKey) String() string {
	return hex.EncodeToString(h)
}

func NewHashKey[K any](key K) (HashKey, error) {
	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(key)
	if err != nil {
		return nil, err
	}

	b := buf.Bytes()

	sha := sha1.New()
	sha.Write(b)
	hased := sha.Sum(nil)

	return HashKey(hased), nil
}

func (h *HashKey) GreaterThan(key any) bool {
	b, err := NewHashKey(key)
	if err != nil {
		return false
	}

	result := bytes.Compare(*h, b) == 1

	log.WithFields(log.Fields{
		"a":      h.String(),
		"b":      b.String(),
		"result": result,
	}).Debugln("[HashKey] GreatherThan")

	return result
}

func (h *HashKey) LessThan(key any) bool {
	b, err := NewHashKey(key)
	if err != nil {
		return false
	}

	result := bytes.Compare(*h, b) == -1

	log.WithFields(log.Fields{
		"a":      h.String(),
		"b":      b.String(),
		"result": result,
	}).Debugln("[HashKey] LessThan")

	return result
}
