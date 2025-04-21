package types

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Memo struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
}

func (m Memo) Format() string {
	if len(m.Tags) == 0 {
		return fmt.Sprintf("[%d] %s", m.ID, m.Content)
	}

	// タグに色付け（色は交互に切り替え or 固定）
	coloredTags := []string{}
	colors := []*color.Color{
		color.New(color.FgCyan),
		color.New(color.FgGreen),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
	}

	for i, tag := range m.Tags {
		c := colors[i%len(colors)]
		coloredTags = append(coloredTags, c.Sprint(tag))
	}

	return fmt.Sprintf("[%d] %s [tags: %s]", m.ID, m.Content, strings.Join(coloredTags, ", "))
}
