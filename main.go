// Licensed under the Creative Commons License.

package main

import (
	"log"

	"github.com/rajasoun/tdd-go/configurator"
)

func main() {
	err := configurator.LoadDotEnv()
	if err != nil {
		log.Fatalln("Error Loading .env File")
	}
}
