package page

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	wd string
	// Pages stores all pages we have
	Pages = make(map[string]string)
)

func init() {
	wd, _ = os.Getwd()
	wd += "/page"
	files, _ := os.Open(wd + "/pages")
	defer files.Close()
	nn, _ := files.Readdirnames(-1)
	for _, n := range nn {
		name := strings.Join(strings.Split(n, ".txt"), "")
		Pages[name] = name
	}
}

// Page is a page
type Page struct {
	Title string
	Body  []byte
	Path  string
}

// NewPage return a new page
func NewPage(title string) *Page {
	p := &Page{}
	Path := wd + "/pages/" + title + ".txt"

	p.Title = title
	p.Path = Path
	return p
}

// Save a page
func (p *Page) Save() error {
	file, err := os.OpenFile(p.Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, ok := Pages[p.Title]; !ok {
		Pages[p.Title] = p.Title
	}

	// we may delected some words, need to be covered
	err = os.Truncate(p.Path, 0)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(p.Body)
	writer.Flush()
	return err
}

// Del a page
func (p *Page) Del() error {
	if _, ok := Pages[p.Title]; !ok {
		return nil
	}

	err := os.Remove(p.Path)
	if err != nil {
		return err
	}
	delete(Pages, p.Title)
	return nil
}

// Load a local page to page struct
func (p *Page) Load() error {
	file, err := os.Open(p.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		b := make([]byte, 512)
		n, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		p.Body = append(p.Body, b[:n]...)
	}
	return nil
}
