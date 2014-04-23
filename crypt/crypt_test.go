package crypt

import "bytes"
import "testing"

func TestEncrypt(t *testing.T) {
    cipher := Encrypt("lykling", "", 1)
    if cipher != "DftjDetmxY" {
        t.Log("`lykling` --> " + cipher)
        t.Fail()
    }

    cipher = Encrypt("", "", 1)
    if cipher != "" {
        t.Log("`` --> " + cipher)
        t.Fail()
    }

    cipher = Encrypt("lykling", "", -1)
    if cipher != "fDjteDmtYx" {
        t.Log("`lykling` --> " + cipher)
        t.Fail()
    }

    cipher = Encrypt("lykling", "echo", 1)
    if cipher != "LXuiLWu21U" {
        t.Log("`lykling` --> " + cipher)
        t.Fail()
    }

    cipher = Encrypt("z9gaimimalo", "ecncswitch", 1)
    println(cipher)

    cipher = Encrypt("20051001", "ecncswitch", 1)
    println(cipher)

    cipher = Encrypt("20051001", "superecnc", 1)
    println(cipher)

    cipher = Encrypt("since2004.", "ecncswitch", 1)
    println(cipher)

    cipher = Encrypt("since2004.", "superecnc", 1)
    println(cipher)

}

func TestDecrypt(t *testing.T) {
    text := Decrypt("DftjDetmxY", "", 1)
    if text != "lykling" {
        t.Log("`DftjDetmxY` --> " + text)
        t.Fail()
    }

    text = Decrypt("fDjteDmtYx", "", -1)
    if text != "lykling" {
        t.Log("`fDjteDmtYx` --> " + text)
        t.Fail()
    }

    text = Decrypt("LXuiLWu21U", "echo", 1)
    if text != "lykling" {
        t.Log("`LXuiLWu21U --> " + text)
        t.Fail()
    }

    text = Decrypt("YhcmQgcrSgi3Dg8", "ecncswitch", 1)
    println(text)
    text = Decrypt("dPeEVLMEELE", "ecncswitch", 1)
    println(text)
    text = Decrypt("mJ3sVLzBELE", "superecnc", 1)
    println(text)
    text = Decrypt("xtmkvg0SELvEbO", "ecncswitch", 1)
    println(text)
    text = Decrypt("1npk3XzTEL3rbO", "superecnc", 1)
    println(text)
}

func TestTransform(t *testing.T) {
    cmap := "echoABCDEFGHIJKLMNOPQRSTUVWXYZabdfgijklmnpqrstuvwxyz0123456789+/"
    src := []byte{'l', 'u'}
    dst := transform([]byte(cmap), src, 1, 1)
    if !bytes.Equal(dst, []byte{'u', '2'}) {
        t.Log("src:" + string(src))
        t.Log("dst:" + string(dst))
    }
}
