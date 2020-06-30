package sapphire

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// PollForUpdates polls the provided URL for updates, and saves the updated data to the specified path if there's an update
func PollForUpdates(path string, url string) {
	ticker := time.NewTicker(60000 * 5)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		etag := ""
		for {
			_ = <-ticker.C
			res, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			newEtag := res.Header.Get("etag")
			if etag != newEtag {
				buf, err := readerToArrayByte(res.Body)
				if err != nil {
					log.Fatalln(err)
				}

				err = ioutil.WriteFile(path, buf, 0644)
				if err != nil {
					log.Fatalln(err)
				}

				etag = newEtag
			}
		}
	}()
	<-done
}

func readerToArrayByte(reader io.Reader) ([]byte, error) {
	// Convert the io.Reader to a bytes.Buffer, so we can then read the bytes.Buffer into the []byte we created earlier
	intermediaryBuf := &bytes.Buffer{}
	nRead, err3 := io.Copy(intermediaryBuf, reader)
	if err3 != nil {
		return nil, err3
	}

	// Read the bytes.Buffer into the []byte
	retBuf := make([]byte, nRead)
	_, err4 := intermediaryBuf.Read(retBuf)
	if err4 != nil {
		return nil, err4
	}

	return retBuf, nil
}
