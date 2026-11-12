// Package quilltest holds a deliberately weak key, added only to check that
// Quill's scheduled scan reads a pull request's own code. Not meant to merge.
package quilltest

import (
	"crypto/rand"
	"crypto/rsa"
)

// WeakKey makes an RSA-1024 key, which a scan of this branch should report
// and a scan of main should not.
func WeakKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}
