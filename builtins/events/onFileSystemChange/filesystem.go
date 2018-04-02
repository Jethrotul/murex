// +build ignore

package onFileSystemChange

import (
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/ansi"
)

const eventType = "onFileSystemChange"

func init() {
	w := newWatch()
	events.AddEventType(eventType, w)
	w.init()
}

type watch struct {
	watcher  *fsnotify.Watcher
	error    error
	mutex    sync.Mutex
	cbBlocks map[string][]rune
}

func newWatch() (w *watch) {
	w = new(watch)
	w.watcher, w.error = fsnotify.NewWatcher()
	w.cbBlocks = make(map[string][]rune)

	return
}

// Callback returns the block to execute upon a triggered event
func (w *watch) callback(path string) (block []rune) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for {
		for len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}

		block = w.cbBlocks[path]

		if len(block) > 0 {
			break
		}

		split := strings.Split(path, "/")
		switch len(split) {
		case 0:
			path = "/"
		case 1:
			path = strings.Join(split, "/")
		default:
			path = strings.Join(split[:len(split)-1], "/")
		}
	}

	return
}

// Init starts a new watch event loop
func (w *watch) init() {
	defer w.watcher.Close()

type Interrupt struct {
	Object: string
	Operation: string
}

	for {
		select {
		case event := <-w.watcher.Events:

			go callback(
				"bob",
				Interrupt{
					Object: event.Name,
					Operation:event.Op.String(),
				},
				w.callback(event.Name),
				proc.ShellProcess.Stdout,
			)

		case err := <-w.watcher.Errors:
			ansi.Stderrln(ansi.FgRed, "error in watcher: "+err.Error())
		}
	}
}

// Add a path to the watch event list
func (w *watch) Add(path string, block []rune) (err error) {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	err = w.watcher.Add(path)
	if err == nil {
		w.cbBlocks[path] = block
	}

	return
}

// Remove a path to the watch event list
func (w *watch) Remove(path string) (err error) {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	w.mutex.Lock()
	defer w.mutex.Unlock()

	err = w.watcher.Remove(path)
	if err == nil {
		delete(w.cbBlocks, path)
	}

	return
}

// Dump returns all the events in w
func (w *watch) Dump() interface{} {
	dump := make(map[string]string)

	w.mutex.Lock()
	for s := range w.cbBlocks {
		dump[s] = string(w.cbBlocks[s])
	}
	defer w.mutex.Unlock()

	return dump
}