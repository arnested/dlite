package main

// build hyperkit
//go:generate make --silent -C hyperkit

// copy binary assets
//go:generate sh -c "curl --silent --output assets/qcow-tool 'https://raw.githubusercontent.com/TheNewNormal/corectl.app/master/src/bin/qcow-tool' && chmod a+x assets/qcow-tool"
//go:generate cp hyperkit/build/com.docker.hyperkit assets/

// generate bundled assets
//go:generate go-bindata -pkg main -o assets.go -prefix assets assets/
