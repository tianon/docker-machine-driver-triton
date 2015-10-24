package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"

	"triton"
)

func main() {
	plugin.RegisterDriver(new(triton.Driver))
}
