package main

import "github.com/hktalent/quic-go/http3"

// The contents of this script don't matter.
// We just need to make sure that quic-go is imported.
func main() {
	_ = http3.Server{}
}
