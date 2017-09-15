package main

import (
	"github.com/roscopecoltran/elasticfeed-plugin/pipeline/null"
	"github.com/roscopecoltran/elasticfeed/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPipeline(new(null.Pipeline))
	server.Serve()
}
