package main

import (
	"encoding/json"
	"fmt"
	"log"

	apiclient "github.com/bols-blue/golang-swagger-client-sample"
)

func main() {
	// create the API client, with the transport
	client := apiclient.NewDefaultApi()

	// make the request to get all items
	pl, resp, err := client.HelloGet()
	fmt.Printf("%s\n", resp.Payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", pl.Msg)
	fmt.Printf("%s\n", pl.Test)
	b, err := json.Marshal(pl)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	// test post
	pl2, resp, err := client.HelloPost()
	fmt.Printf("%s\n", resp.Payload)
	if err != nil {
		log.Fatal(err)
	}
	c, err := json.Marshal(pl2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	if pl.Int_ != nil {
		fmt.Printf("%+v\n", pl.Int_)
	}
	if pl.Bool_ != nil {
		fmt.Printf("%+v\n", pl.Bool_)
	}
	fmt.Println(string(c))
	// test put
	pl2, resp, err = client.HelloPut()
	fmt.Printf("%s\n", resp.Payload)
	if err != nil {
		log.Fatal(err)
	}
	c, err = json.Marshal(pl2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Printf("%+v\n", *pl2.Int_)
	fmt.Printf("%+v\n", *pl2.Bool_)
	fmt.Println(string(c))

}
