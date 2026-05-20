package project

import (
	"strings"

	"github.com/hanymamdouh82/operatree/internal/subject"
)

const (
	ansiReset  = "\033[0m"
	ansiBold   = "\033[1m"
	ansiDim    = "\033[2m"
	ansiItalic = "\033[3m"
	ansiPurple = "\033[38;5;141m"
	ansiYellow = "\033[38;5;221m"
	ansiGray   = "\033[38;5;244m"
	ansiGreen  = "\033[38;5;114m"
)

func label(s string) string {
	return ansiPurple + ansiBold + s + ansiReset
}

func value(s string) string {
	return ansiReset + s
}

func dim(s string) string {
	return ansiGray + ansiItalic + s + ansiReset
}

func formatPreview(s subject.Subject) string {
	var b strings.Builder

	// Header
	b.WriteString(ansiBold + ansiPurple + s.Name + ansiReset)
	b.WriteString(ansiGray + "  ·  " + string(s.Type) + ansiReset + "\n")
	b.WriteString(ansiGray + strings.Repeat("─", 35) + ansiReset + "\n\n")

	// Core fields
	if s.Date != "" {
		b.WriteString(label("Date      ") + value(s.Date) + "\n")
	}
	if s.Location != "" {
		b.WriteString(label("Location  ") + value(s.Location) + "\n")
	}
	if s.Owner != "" {
		b.WriteString(label("Owner     ") + value(s.Owner) + "\n")
	}
	if s.Status != "" {
		b.WriteString(label("Status    ") + value(s.Status) + "\n")
	}
	if len(s.Paricipants) > 0 {
		b.WriteString(label("People    ") + value(strings.Join(s.Paricipants, ", ")) + "\n")
	}

	// Tags
	if len(s.Tags) > 0 {
		tags := make([]string, len(s.Tags))
		for i, t := range s.Tags {
			tags[i] = ansiYellow + "#" + t + ansiReset
		}
		b.WriteString(label("Tags      ") + strings.Join(tags, "  ") + "\n")
	}

	// Notes
	if s.Notes != "" {
		b.WriteString("\n" + label("Notes") + "\n")
		b.WriteString(dim(s.Notes) + "\n")
	}

	// Relations
	hasRelations := s.RelatedObjective != "" || len(s.RelatedEvents) > 0 || len(s.Outputs) > 0
	if hasRelations {
		b.WriteString("\n" + ansiGray + strings.Repeat("─", 35) + ansiReset + "\n")
		if s.RelatedObjective != "" {
			b.WriteString(label("Objective ") + value(s.RelatedObjective) + "\n")
		}
		if len(s.RelatedEvents) > 0 {
			b.WriteString(label("Events    ") + value(strings.Join(s.RelatedEvents, ", ")) + "\n")
		}
		if len(s.Outputs) > 0 {
			b.WriteString(label("Outputs   ") + value(strings.Join(s.Outputs, ", ")) + "\n")
		}
	}

	return b.String()
}
