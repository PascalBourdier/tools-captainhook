package main

import (
	"fmt"

	cli "github.com/jawher/mow.cli"

	"github.com/mirakl/tools-captain-hook/github"
)

var version = "dev"

func main() {
	app := cli.App("cgtls", "A tool to repeat Github webhook to Mirakl Kubernetes clusters.\nFor issues or feature requests, please go to https://github.com/mirakl/tools-captain-hook. Pull requests are welcome !")
	app.Version("v version", fmt.Sprintf("cgtls %s", version))

	github.WebhookReceiver()
}
