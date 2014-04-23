package crypt

import "testing"

func TestEncrypt(t *testing.T) {
    cipher := Encrypt("lykling", "", 1)
    if cipher != "DftjDetmxY" {
        t.Log("lykling --> " + cipher)
        t.Fail()
    }
}

func TestDecrypt(t *testing.T) {
    text := Decrypt("DftjDetmxY", "", 1)
    if text != "lykling" {
        t.Log("DftjDetmxY --> " + text)
        t.Fail()
    }
}
