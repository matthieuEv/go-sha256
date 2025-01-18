package main

import (
    "testing"
	"github.com/matthieuEv/go-sha256/sha256"
)

func TestRightRotate(t *testing.T) {
	var x uint32 = 0x12345678
	var n uint = 4
	var expected uint32 = 0x81234567
	if sha256.RightRotate(x, n) != expected {
		t.Fail()
	}
}

func TestRightShift(t *testing.T) {
	var x uint32 = 0x12345678
	var n uint = 4
	var expected uint32 = 0x1234567
	if sha256.RightShift(x, n) != expected {
		t.Fail()
	}
}

func TestUint32ToHex(t *testing.T) {
	var num uint32 = 0x12345678
	var expected string = "12345678"
	if sha256.Uint32ToHex(num) != expected {
		t.Fail()
	}
}

func TestGetNbrBits(t *testing.T) {
	var num uint8 = 0b1101
	var expected int = 4
	if sha256.GetNbrBits(num) != expected {
		print(sha256.GetNbrBits(num))
		t.Fail()
	}
}

var hash = []map[string]string{
	{"": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	{"a": "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"},
	{"ab": "fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603"},
	{"abc": "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
	{"abcd": "88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589"},
	{"abcde": "36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c"},
	{"Tout corps plongé dans un liquide, reçoit une poussée de bas en haut qui est égale au volume d'eau déplacée. -Archimède": "bb1126c7218ecc53d1542576b9ba389d32a731216919113499193e1487e997b1"},
}

func TestSum256(t *testing.T) {
	for _, value := range hash {
		for k, v := range value {
			if sha256.Sum256([]byte(k)) != v {
				t.Fail()
			}
		}
	}
}