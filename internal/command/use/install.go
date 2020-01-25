package use

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/mitchellh/ioprogress"
)

func GetFileSha256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// download accepts a reader, writer, and expected asset size, copies
// the bytes through and displays progress to the console.
func Download(reader io.Reader, writer io.Writer, size int64, title string) error {
	// Attempt 1:
	// bar := pb.New(size).SetUnits(pb.U_BYTES)
	// bar.Prefix(title)
	// bar.Start()
	// reader = bar.NewProxyReader(reader)

	// Attempt 2: mpb
	//progress := mpb.New()
	//bar := progress.AddBar(size, mpb.BarTrim())
	//reader = bar.ProxyReader(reader)

	// Attempt 3: ioprogress
	reader = &ioprogress.Reader{
		Reader: reader,
		Size:   size,
	}

	_, err := io.Copy(writer, reader)
	if err != nil {
		return err
	}

	//bar.Finish()
	//progress.Stop()

	return nil
}
