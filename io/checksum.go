package io

type ChecksumMethod int

const (
	MD5 ChecksumMethod = iota
	SHA1
	SHA256
)
