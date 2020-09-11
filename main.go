package main

import (
	_ "github.com/mdempsky/gocode"
	_ "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
	_ "github.com/ramya-rao-a/go-outline"
	_ "github.com/acroca/go-symbols"
	_ "github.com/fatih/gomodifytags"
	_ "github.com/josharian/impl"
	_ "github.com/davidrjenni/reftools/cmd/fillstruct"
	_ "github.com/haya14busa/goplay/cmd/goplay"
    _ "github.com/godoctor/godoctor"
    _ "github.com/go-delve/delve/cmd/dlv"
    _ "github.com/stamblerre/gocode"
    _ "github.com/rogpeppe/godef"
    _ "github.com/sqs/goreturns"
    _ "golang.org/x/tools/cmd/guru"
    _ "golang.org/x/tools/cmd/gorename"
    _ "golang.org/x/lint/golint"
)

func main() {
	fmt.Println("Install plugin for vscode!")
}