package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"syscall"

	"github.com/musou1500/sit"
)

func addFileToCache() {}

func createCeMode(mode uint32) uint32 {
	if mode&syscall.S_IFLNK != 0 {
		return sit.Htonl(syscall.S_IFLNK)
	}

	var permissions uint32
	if mode&0o100 != 0 {
		permissions = 0o755
	} else {
		permissions = 0o644
	}
	return sit.Htonl(syscall.S_IFREG | permissions)
}

func main() {

	for _, path := range os.Args {
		var stat *syscall.Stat_t
		error := syscall.Lstat(os.Args[1], stat)
		if error != nil {
			log.Fatal(error)
		}

		flags := sit.Htons(uint16(len(path)))
		ce := CacheEntry{
			Name:  path,
			Flags: flags,
			Mode:  createCeMode(stat.Mode),
			Ctime: CacheTime{
				Sec:  sit.Htonl(uint32(stat.Ctim.Sec)),
				Nsec: sit.Htonl(uint32(stat.Ctim.Nsec)),
			},
			Mtime: CacheTime{
				Sec:  sit.Htonl(uint32(stat.Mtim.Sec)),
				Nsec: sit.Htonl(uint32(stat.Mtim.Nsec)),
			},
			Dev:  sit.Htonl(uint32(stat.Dev)),
			Ino:  sit.Htonl(uint32(stat.Ino)),
			Uid:  sit.Htonl(stat.Uid),
			Gid:  sit.Htonl(stat.Gid),
			Size: sit.Htonl(uint32(stat.Size)),
		}

		switch stat.Mode & syscall.S_IFMT {
		case syscall.S_IFREG:
		case syscall.S_IFLNK:
			target, err := os.Readlink(path)
			if err != nil {
				log.Fatal(err)
			}
			err = sit.WriteSha1File(bytes.NewBufferString(target).Bytes(), "blob")
			if err != nil {
				log.Fatal(err)
			}
		default:
			errors.New("invalid mode")
		}
	}

	// TODO: write cache
}
