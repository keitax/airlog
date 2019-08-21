package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	filenameRegexp    = regexp.MustCompile(`^(\d{8})\-.+\.md$`)
	frontMatterRegexp = regexp.MustCompile(`(?ms)^---\s*$\n(.*?)^---\s*$\n(.*)`)
	h1Regexp          = regexp.MustCompile(`^#\s+(.+)\s*$`)
)

func IsPostFileName(filename string) bool {
	return filenameRegexp.MatchString(filename)
}

type PostFile struct {
	Filename string
	Content  string
}

func (pf *PostFile) GetTimestamp() time.Time {
	ms := filenameRegexp.FindStringSubmatch(pf.Filename)
	if len(ms) < 2 {
		panic(fmt.Errorf("must not happen: %v", ms))
	}
	t, err := time.Parse("20060102", ms[1])
	if err != nil {
		panic(err) // must not happen
	}
	return t
}

func (pf *PostFile) ExtractFrontMatter() map[string]interface{} {
	ms := frontMatterRegexp.FindStringSubmatch(pf.Content)
	if len(ms) < 3 {
		return map[string]interface{}{}
	}
	if len(ms) > 3 {
		panic("BUG: must not happen")
	}
	metadataSection, bodySection := ms[1], ms[2]
	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(metadataSection), &metadata); err != nil {
		return map[string]interface{}{}
	}
	pf.Content = bodySection
	return metadata
}

func (pf *PostFile) ExtractH1() string {
	r := bufio.NewReader(bytes.NewBufferString(pf.Content))
	buf := &bytes.Buffer{}
	var h1 string
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if h1 == "" {
			if ms := h1Regexp.FindStringSubmatch(string(line)); len(ms) > 1 {
				h1 = ms[1]
				continue
			}
		}
		fmt.Fprintln(buf, string(line))
	}
	pf.Content = buf.String()
	return h1
}
