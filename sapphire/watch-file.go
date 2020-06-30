package sapphire

import (
	"io"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// WatchFile watches the provided file path for updates, running the provided function with the updated data if an update occurs.
func WatchFile(path string, fnUpdate func(io.Reader)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					stream, err := os.Open(path)
					if err != nil {
						log.Fatalln(err)
					}
					fnUpdate(stream)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatalln(err)
	}
	<-done
}
