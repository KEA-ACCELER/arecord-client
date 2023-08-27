package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func DirWatchStart() {
	// Create a new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Define a channel to receive events.
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)

				//이벤트가 일어난 파일만 전송
				filepath := event.Name
				info, err := os.Stat(filepath)

				if err != nil {
					log.Panicln(err)
				}

				if event.Op&fsnotify.Create == fsnotify.Create && info.IsDir() {
					watcher.Add(filepath)
				}
				if event.Op&fsnotify.Create == fsnotify.Create && !info.IsDir() {
					RestClientFile(filepath)
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					filepath = LocalAbsToRoot(filepath, getAccelerDir())
					RestClientMessage("removed", []byte(filepath))
				}

				if event.Op&fsnotify.Write == fsnotify.Write {

					RestClientFile(filepath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a file or directory to watch.
	err = watcher.Add(getAccelerDir())
	if err != nil {
		log.Fatal(err)
	}

	// Wait until the channel is closed.
	<-done

}
