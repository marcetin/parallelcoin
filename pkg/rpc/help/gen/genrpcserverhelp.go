package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/p9c/pod/pkg/logi"

	"github.com/p9c/pod/pkg/rpc/btcjson"
	rpchelp "github.com/p9c/pod/pkg/rpc/help"

	// rpchelp "github.com/p9c/pod/pkg/rpc/help"
)

var outputFile = func() *os.File {
	fi, err := os.Create("../rpcserverhelp.go")
	if err != nil {
		log.L.Error(err)
		log.L.Fatal(err)
	}
	return fi
}()

func writefln(format string, args ...interface{}) {
	_, err := fmt.Fprintf(outputFile, format, args...)
	if err != nil {
		log.L.Error(err)
		log.L.Fatal(err)
	}
	_, err = outputFile.Write([]byte{'\n'})
	if err != nil {
		log.L.Error(err)
		log.L.Fatal(err)
	}
}
func writeLocaleHelp(locale, goLocale string, descs map[string]string) {
	funcName := "helpDescs" + goLocale
	writefln("func %s() map[string]string {", funcName)
	writefln("return map[string]string{")
	for i := range rpchelp.Methods {
		m := &rpchelp.Methods[i]
		helpText, err := btcjson.GenerateHelp(m.Method, descs, m.ResultTypes...)
		if err != nil {
			log.L.Error(err)
			log.L.Fatal(err)
		}
		writefln("%q: %q,", m.Method, helpText)
	}
	writefln("}")
	writefln("}")
}
func writeLocales() {
	writefln("var localeHelpDescs = map[string]func() map[string]string{")
	for _, h := range rpchelp.HelpDescs {
		writefln("%q: helpDescs%s,", h.Locale, h.GoLocale)
	}
	writefln("}")
}
func writeUsage() {
	usageStrs := make([]string, len(rpchelp.Methods))
	var err error
	for i := range rpchelp.Methods {
		usageStrs[i], err = btcjson.MethodUsageText(rpchelp.Methods[i].Method)
		if err != nil {
			log.L.Error(err)
			log.L.Fatal(err)
		}
	}
	usages := strings.Join(usageStrs, "\n")
	writefln("var requestUsages = %q", usages)
}
func main() {
	defer outputFile.Close()
	packageName := "main"
	if len(os.Args) > 1 {
		packageName = os.Args[1]
	}
	writefln("// AUTOGENERATED by internal/rpchelp/genrpcserverhelp.go; do not edit.")
	writefln("")
	writefln("package %s", packageName)
	writefln("")
	for _, h := range rpchelp.HelpDescs {
		writeLocaleHelp(h.Locale, h.GoLocale, h.Descs)
		writefln("")
	}
	writeLocales()
	writefln("")
	writeUsage()
}
