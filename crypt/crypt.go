package crypt

import (
    "bytes"
    "encoding/base64"
    "strings"
)

const (
    ORIG_MAP = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    DEFAULT_MAP_RANK = 8
    DEFAULT_MAP_SCALE = 64
    CRYPT_DIRECTION_HORIZONTAL = -1
    CRYPT_DIRECTION_VERTICAL = 1
    CRYPT_ENCRYPT = 1
    CRYPT_DECRYPT = -1
)

func build_map(key string) (result []byte) {
    result = []byte{}
    for _, chr := range key + ORIG_MAP {
        if ! bytes.Contains(result, []byte{byte(chr)}) {
            result = append(result, byte(chr))
        }
    }
    return
}

func transform(keymap, src []byte, flag, dir int) (dst []byte) {
    if len(src) != 2 {
        return
    }
    rank := DEFAULT_MAP_RANK
    scale := DEFAULT_MAP_SCALE
    posl_old := bytes.IndexByte(keymap, src[0])
    posr_old := bytes.IndexByte(keymap, src[1])
    posl_new, posr_new := 0, 0
    deltax := (posl_old % rank) - (posr_old % rank)
    deltay := (posl_old / rank) - (posr_old / rank)
    if deltax == 0 && deltay == 0 {
        posl_new = scale - posl_old - 1
        posr_new = scale - posr_old - 1
    } else if deltax == 0 {
        posl_new = (posl_old + rank * flag) % scale
        posr_new = (posr_old + rank * flag) % scale
    } else if deltay == 0 {
        posl_new = posl_old / rank * rank + (posl_old + flag) % rank
        posr_new = posr_old / rank * rank + (posr_old + flag) % rank
    } else {
        if dir == CRYPT_DIRECTION_VERTICAL {
            posl_new = posl_old - deltay * rank
            posr_new = posr_old + deltay * rank
        } else {
            posl_new = posl_old - deltax
            posr_new = posr_old + deltax
        }
    }
    dst = append([]byte{}, keymap[posl_new], keymap[posr_new])
    return

}

func Encrypt(text, key string, dir int) (cipher string){
    crypt_map := build_map(key)
    b64 := strings.Trim(base64.StdEncoding.EncodeToString([]byte(text)), "=")
    cipher = ""
    for i := 0; i < len(b64) - 1; i += 2 {
        cipher += string(transform(crypt_map, []byte(b64[i:i+2]), CRYPT_ENCRYPT, dir))
    }
    if len(b64) % 2 == 1 {
        cipher += b64[len(b64)-1:]
    }
    return
}

func Decrypt(cipher, key string, dir int) (text string){
    crypt_map := build_map(key)
    b64 := ""
    for i := 0; i < len(cipher) - 1; i += 2 {
        b64 += string(transform(crypt_map, []byte(cipher[i:i+2]), CRYPT_DECRYPT, dir))
    }
    if len(cipher) % 2 == 1 {
        b64 += cipher[len(cipher)-1:]
    }
    b64 += strings.Repeat("=", 3 - ((len(b64) - 1) % 4))
    tmp, _ := base64.StdEncoding.DecodeString(b64)
    text = string(tmp)
    return
}
