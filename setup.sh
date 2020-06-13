#!/bin/bash

if [ "$blogprojectset" == "true" ]; then
    echo "everything is set. You're ready to go"
    echo "to run with hot-reload run:"
    echo "  air ."

    exit
fi

if [ -z "$GOPATH" ]; then
    echo "\$GOPATH is empty"
    exit
fi

go get -u github.com/cosmtrek/air

echo "#this is config for ghibran.xyz/blog-backend" >>~/.bashrc
echo
echo "export PATH=\"\$PATH:\$GOPATH/bin\"" >>~/.bashrc
echo "export PORT=\"8080\"" >>~/.bashrc
echo "export blogprojectset=\"true\"" >>~/.bashrc
echo "# end config" >> ~/.bashrc


if [-f "db.secret"]; then
    printf "export dbIP=\"\" \nexport dbUSR=\"\" \nexport dbPW=\"\" \nexport dbName=\"\"" > db.secret
fi

echo "add your database environment variables in db.secret"

echo
echo "do 'source db.secret' "
echo "and run with hot-reload do:"
echo "  air ."

go mod tidy
bash
