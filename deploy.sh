#!/bin/bash

go get -v github.com/VaguelySerious/scontrino-api
cd ~/go/src/github.com/VaguelySerious/scontrino-api/
go build
systemctl stop scontrino
cat scontrino-api > /var/www/scontrino/scontrino-api
systemctl start scontrino
