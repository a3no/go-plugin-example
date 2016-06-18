package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/a3no/go-plugin-example/plugin"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient("./plugin1")
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(plugin.Greeter)
	fmt.Println(greeter.Greet())
}
