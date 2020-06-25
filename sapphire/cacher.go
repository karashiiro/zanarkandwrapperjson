package sapphire

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// GetFile downloads a file from the specified URL and caches it, or loads the cached version.
func GetFile(filename string, url string) (io.Reader, error) {
	fileBuf, err1 := ioutil.ReadFile(filename)

	res, err2 := http.Get(url)
	if err2 != nil {
		log.Println("Could not access internet resource, falling back to cached resource.")
		if err1 != nil {
			return bytes.NewReader(fileBuf), nil
		}
		return nil, errors.New("getfile: no cached resource available")
	}

	defer res.Body.Close()

	// Convert the io.Reader to a bytes.Buffer, so we can then read the bytes.Buffer into the []byte we created earlier
	intermediaryBuf := &bytes.Buffer{}
	nRead, err3 := io.Copy(intermediaryBuf, res.Body)
	if err3 != nil {
		return nil, err3
	}
	// Read the bytes.Buffer into the []byte
	internetBuf := make([]byte, nRead)
	_, err4 := intermediaryBuf.Read(internetBuf)
	if err4 != nil {
		return nil, err4
	}

	if err1 != nil || !reflect.DeepEqual(fileBuf, internetBuf) {
		// If the file doesn't exist or the downloaded stuff is different, just write out the downloaded stuff and return
		err5 := ioutil.WriteFile(filename, internetBuf, 0644)
		if err5 != nil {
			return nil, err5
		}
		return bytes.NewReader(internetBuf), nil
	}

	return bytes.NewReader(fileBuf), nil
}
