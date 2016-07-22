package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encodeMe := "not-so-secret-message'!@#$%^<>{}&*()~1234567890[];"
	stdStr := base64.StdEncoding.EncodeToString([]byte(encodeMe))
	fmt.Println(stdStr)

	urlStr := base64.URLEncoding.EncodeToString([]byte(encodeMe))
	fmt.Println(urlStr)

	urlStrRaw := base64.RawURLEncoding.EncodeToString([]byte(encodeMe))
	fmt.Println(urlStrRaw)

	dec, _ := base64.StdEncoding.DecodeString(stdStr)
	// encoding requires casting string to byte array;
	// decoding requires casting byte array to string
	fmt.Println(string(dec))
}
