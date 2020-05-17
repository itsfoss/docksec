package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/sirupsen/logrus"
)

func main() {
    logrus.Info("Plugin start")
	plug, err := newPlugin()
	if err != nil {
		logrus.Fatal(err)
	}

	hl := authorization.NewHandler(plug)
	if err := hl.ServeUnix(plug.name, 0); err != nil {
	    logrus.Fatal(err)
	}
}
