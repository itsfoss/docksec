package main

import (
	auth "github.com/docker/go-plugins-helpers/authorization"
	"github.com/itsfoss/docksec/conf"
)

type plugin struct {
	name string
	desc bool
}

func newPlugin() (*plugin, error) {
	return &plugin{name: "docksec", desc: conf.GetDescStat()}, nil
}

func (plug *plugin) AuthZReq(req auth.Request) auth.Response {
	nallowed, dmsg := conf.GetStatus(req.RequestMethod, req.RequestURI)
	if nallowed {
		return auth.Response{Allow: false, Msg: dmsg}
	}
	return auth.Response{Allow: true}
}

func (plug *plugin) AuthZRes(req auth.Request) auth.Response {
	return auth.Response{Allow: true}
}
