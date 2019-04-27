package domain

import (
	"gopkg.in/yaml.v2"
	"regexp"
)

var frontMatterRegexp = regexp.MustCompile(`(?ms)^---\s*$\n(.*?)^---\s*$\n(.*)`)

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

