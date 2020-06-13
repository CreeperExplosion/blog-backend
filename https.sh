#!/bin/bash
export DOMAIN="api.ghibran.xyz"
export CERT="/etc/letsencrypt/live/api.ghibran.xyz/fullchain.pem"
export KEY="/etc/letsencrypt/live/api.ghibran.xyz/privkey.pem"
source db.secret

echo $dbIP $dbPW
#go run main.go  --ssl=true