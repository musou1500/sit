package sit

import (
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func getObjectDirectory() string {
	return ".sit/objects"
}

func sha1FileName(hash string) string {
	return fmt.Sprintf("%s/%s/%s", getObjectDirectory(), hash[:2], hash[2:])
}

func WriteSha1FilePrepare(tp string, buf []byte) (string, string) {
	hdr := fmt.Sprintf("%s %d", tp, len(buf))
	h := sha1.New()
	io.WriteString(h, hdr)
	h.Write(buf)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return sha1FileName(hash), hdr
}

func WriteSha1File(buf []byte, tp string) error {
	filename, hdr := WriteSha1FilePrepare(tp, buf)
	if _, err := os.Stat(filename); err == nil {
		return nil
	}

	tempfile, err := os.CreateTemp(getObjectDirectory(), "obj_*")
	if err != nil {
		return err
	}

	defer tempfile.Close()

	writer := zlib.NewWriter(tempfile)
	defer writer.Close()
	io.WriteString(writer, hdr)
	writer.Write(buf)
	tempfile.Chmod(0o444)
	err = os.Rename(tempfile.Name(), filename)
	if err != nil {
		return err
	}

	return nil
}

func IndexFd( /*fd, stat, write_object, type*/ ) {}
