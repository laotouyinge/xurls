/* Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc> */
/* See LICENSE for licensing information */

package xurls

import "regexp"

//go:generate go run tools/tldsgen/main.go
//go:generate go run tools/regexgen/main.go

const (
	letters   = "a-zA-Z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF"
	iriChar   = letters + `0-9`
	pathChar  = iriChar + `.,:;/\-+_()?@&=$~!*%'"`
	ipv4Addr  = `(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[0-9])`
	ipv6Addr  = `([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:[0-9a-fA-F]{0,4}|:[0-9a-fA-F]{1,4})?|(:[0-9a-fA-F]{1,4}){0,2})|(:[0-9a-fA-F]{1,4}){0,3})|(:[0-9a-fA-F]{1,4}){0,4})|:(:[0-9a-fA-F]{1,4}){0,5})((:[0-9a-fA-F]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9a-fA-F]{1,4}:){1,6}|:):[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){7}:`
	ipAddr    = `(` + ipv4Addr + `|` + ipv6Addr + `)`
	iri       = `[` + iriChar + `]([` + iriChar + `\-]{0,61}[` + iriChar + `])?`
	hostName  = `((` + iri + `\.)+` + gtld + `|` + ipAddr + `|localhost)`
	wellParen = `([` + pathChar + `]*(\([` + pathChar + `]*\))+)+`
	path      = `(/(` + wellParen + `|[` + pathChar + `]*[` + iriChar + `/])?)?`
	webURL    = hostName + `(:[0-9]{1,5})?` + path
	email     = `[a-zA-Z0-9._%\-+]{1,256}@` + hostName

	defUrl    = `(` + wellParen + `|[` + pathChar + `]*[` + iriChar + `/])`
	scheme    = `[a-zA-Z.\-+]+://`
	allStrict = scheme + defUrl
	all       = allStrict + `|` + webURL + `|` + email
)

var (
	// All matches all kinds of urls
	All       = regexp.MustCompile(all)
	AllStrict = regexp.MustCompile(allStrict)
)

func init() {
	All.Longest()
	AllStrict.Longest()
}
