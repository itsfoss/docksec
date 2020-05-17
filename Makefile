GO?=go
LIB?=
build:
	cp -r src/admin $$HOME/go/src/
	for x in "github.com/sirupsen/logrus" \
			 "github.com/go-yaml/yaml" \
			 "github.com/docker/go-plugins-helpers/authorization"; do \
	$(GO) get $$x; \
	done
	$(GO) build -o bin/admin-authz src/*.go

install:
	install -m700 -o0 -g0 bin/admin-authz /usr/local/bin/admin-authz
	install -m700 -o0 -g0 service/admin-authz.service /lib/systemd/system/admin-authz.service
	install -m700 -o0 -g0 scripts/admin-authz.sh /usr/local/bin/hauthz
