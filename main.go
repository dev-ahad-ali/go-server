package main

import "ecommerce/cmd"

func main() {
	cmd.Serve()

	// secret := []byte("my-secret")
	// message := []byte("Hello world")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)

	// text := h.Sum(nil)

	// fmt.Println(text)
}
