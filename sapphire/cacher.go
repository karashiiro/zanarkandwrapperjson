package sapphire

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// GetFile downloads a file from the specified URL and caches it, or loads the cached version. TODO make it detect changes.
func GetFile(filename string, url string) (io.Reader, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		// Convert the io.Reader to a bytes.Buffer, so we can then read the bytes.Buffer into the []byte we created earlier
		intermediaryBuf := &bytes.Buffer{}
		nRead, err := io.Copy(intermediaryBuf, res.Body)
		if err != nil {
			return nil, err
		}

		// Read the bytes.Buffer into the []byte and write out to a file
		buf = make([]byte, nRead)
		_, err = intermediaryBuf.Read(buf)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(filename, buf, 0644)
		if err != nil {
			return nil, err
		}
	}
	return bytes.NewReader(buf), nil
}
