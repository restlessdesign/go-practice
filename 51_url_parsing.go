package main

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
)

func main() {
	secure_url := "https://admin:12345@secure-vault.com:80/home?lang=en&logged_in=1"

	parsed_url, err := url.Parse(secure_url)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed_url.Scheme)

	user := parsed_url.User
	username := user.Username()
	password, _ := user.Password()
	fmt.Println(username)
	fmt.Println(password)

	fmt.Println(parsed_url.Host)
	// The host will contain both the hostname and the port by default
	// We can split it up like so:
	host, port, _ := net.SplitHostPort(parsed_url.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(parsed_url.Path)

	params, _ := url.ParseQuery(parsed_url.RawQuery)
	fmt.Println("Language: ", params["lang"][0])
	logged_in, _ := strconv.ParseBool(params["logged_in"][0])
	fmt.Println("Authenticated: ", logged_in)

}
