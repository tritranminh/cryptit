package main

import ("fmt"
"github.com/tritranminh/cryptit/encrypt"
"github.com/tritranminh/cryptit/decrypt"
)

func main() {
	encryptedString := encrypt.Nimbus("Hello World")
	fmt.Println(encryptedString)
	fmt.Println(decrypt.Nimbus(encryptedString))
}
