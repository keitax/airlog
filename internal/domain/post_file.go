package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"regexp"
)

var (
	frontMatterRegexp = regexp.MustCompile(`(?ms)^---\s*$\n(.*?)^---\s*$\n(.*)`)
	h1Regexp          = regexp.MustCompile(`^#\s+(.+)\s*$`)
)

func ExtractFrontMatter(content string) (map[string]interface{}, string) {
	ms := frontMatterRegexp.FindStringSubmatch(content)
	if len(ms) < 3 {
		return map[string]interface{}{}, content
	}
	if len(ms) > 3 {
		panic("BUG: must not happen")
	}
	metadataSection, bodySection := ms[1], ms[2]
	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(metadataSection), &metadata); err != nil {
		return map[string]interface{}{}, bodySection
	}
	return metadata, bodySection
}

func ExtractH1(content string) (string, string) {
	r := bufio.NewReader(bytes.NewBufferString(content))
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
		if ms := h1Regexp.FindStringSubmatch(string(line)); len(ms) > 1 {
			h1 = ms[1]
		} else {
			fmt.Fprintln(buf, string(line))
		}
	}
	return h1, buf.String()
}
