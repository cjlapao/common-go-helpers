package io

type ChecksumMethod int

const (
	ChecksumMD5 ChecksumMethod = iota
	ChecksumSHA1
	ChecksumSHA256
)
