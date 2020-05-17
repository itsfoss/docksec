package main

import (
	auth "github.com/docker/go-plugins-helpers/authorization"
    "conf"
)

type plugin struct {
    name string
    port uint32
    socket bool
    desc bool
}

func newPlugin() (*plugin, error) {
    return &authz{name: "docksec", port: conf.GetPort(), socket: conf.GetSockStat(), desc: conf.GetDescStat()}, nil
}

func (plug *plugin) AuthZReq(req auth.Request) auth.Response {
    nallowed, dmsg, _ := conf.GetStatus(req.RequestMethod, req.RequestURI)
	if nallowed {
		return auth.Response{Allow: false, Msg: dmsg}
	}
    return auth.Response{Allow: true}
}

func (plug *plugin) AuthZRes(req auth.Request) auth.Response {
	return auth.Response{Allow: true}
}
