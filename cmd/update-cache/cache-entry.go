package main

type CacheTime struct {
	Sec  uint32
	Nsec uint32
}

type CacheEntry struct {
	Ctime CacheTime
	Mtime CacheTime
	Dev   uint32
	Ino   uint32
	Mode  uint32
	Uid   uint32
	Gid   uint32
	Size  uint32
	Sha1  [20]byte
	Flags uint16
	Name  string
}
