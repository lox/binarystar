package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/tjarratt/babble"
)

func main() {
	dir := os.Args[1]
	rand.Seed(time.Now().Unix())

	for {
		switch randomInt(1, 7) {
		case 1, 2, 3:
			if err := createRandomFile(dir); err != nil {
				log.Fatal("createRandomFile", err)
			}
		case 4, 5:
			if err := modifyRandomFile(dir); err != nil {
				log.Fatal("modifyRandomFile", err)
			}
		case 6:
			if err := deleteRandomFile(dir); err != nil {
				log.Fatal("deleteRandomFile", err)
			}
		}
		time.Sleep(time.Second)
	}
}

type randomDataMaker struct {
	src rand.Source
}

func (r *randomDataMaker) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(r.src.Int63() & 0xff)
	}
	return len(p), nil
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func createRandomFile(dir string) error {
	babbler := babble.NewBabbler()
	babbler.Separator = string(os.PathSeparator)
	babbler.Count = randomInt(1, 10)

	filename := filepath.Join(dir, babbler.Babble()+".txt")
	if err := os.MkdirAll(filepath.Dir(filename), 0700); err != nil {
		return err
	}

	randData := &randomDataMaker{
		rand.NewSource(int64(time.Now().Nanosecond())),
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	n, err := io.CopyN(f, randData, int64(randomInt(100, 1000000)))
	if err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	log.Printf("Wrote %d bytes to %s", n, filename)
	return nil
}

func modifyRandomFile(dir string) error {
	f, err := pickRandomFile(dir)
	if err == os.ErrNotExist {
		return nil
	}
	if err != nil {
		return err
	}
	stat, err := os.Stat(f)
	if err != nil {
		return err
	}
	switch randomInt(1, 6) {
	case 1, 2, 3:
		log.Printf("Modifying content of %s", f)
		b := []byte("llamas")

		fo, err := os.OpenFile(f, os.O_WRONLY, stat.Mode())
		if err != nil {
			return err
		}

		offset := rand.Int63n(stat.Size() - int64(len(b)))
		log.Printf("Writing %d bytes at file offset %d", len(b), offset)

		if _, err = fo.WriteAt(b, offset); err != nil {
			return err
		}

		if err = fo.Close(); err != nil {
			return err
		}

	case 4:
		log.Printf("Modifying mode of %s", f)
		modes := []os.FileMode{
			0777,
			0644,
			0600,
			0700,
			0701,
			0705,
			0655,
		}

		if err := os.Chmod(f, modes[randomInt(0, len(modes))]); err != nil {
			return err
		}

	case 5:
		log.Printf("Modifying time modified of %s", f)
		durations := []time.Duration{
			time.Second * 10,
			time.Hour * 100,
			time.Minute * 1000,
			time.Minute * -1000,
			time.Minute * -10000,
		}

		ctime := stat.ModTime()
		mtime := stat.ModTime().Add(durations[randomInt(0, len(durations))])

		if err := os.Chtimes(f, ctime, mtime); err != nil {
			return err
		}
	}
	return nil
}

func deleteRandomFile(dir string) error {
	log.Printf("Deleting a file")
	f, err := pickRandomFile(dir)
	if err == os.ErrNotExist {
		return nil
	}
	if err != nil {
		return err
	}
	log.Printf("Deleting %s", f)
	return os.Remove(f)
}

func pickRandomFile(dir string) (string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	log.Printf("Found %d files", len(files))
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", os.ErrNotExist
	}
	if len(files) == 1 {
		return files[0], nil
	}
	return files[randomInt(0, len(files)-1)], nil
}
