package store

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Note struct {
	Path     string
	Title    string
	Created  time.Time
	Tags     []string
	Filename string
}

// ParseNote reads a note file and extracts metadata
func ParseNote(path string) (*Note, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	note := &Note{
		Path:     path,
		Filename: filepath.Base(path),
		Tags:     []string{},
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	// Extract title (first # line)
	titleRegex := regexp.MustCompile(`^#\s+(.+)`)
	for _, line := range lines {
		if match := titleRegex.FindStringSubmatch(line); match != nil {
			note.Title = match[1]
			break
		}
	}

	// Extract Created date
	createdRegex := regexp.MustCompile(`Created::\s*(\d{4}-\d{2}-\d{2})`)
	if match := createdRegex.FindStringSubmatch(text); match != nil {
		if t, err := time.Parse("2006-01-02", match[1]); err == nil {
			note.Created = t
		}
	}

	// Extract Tags
	tagsRegex := regexp.MustCompile(`Tags::\s*\[([^\]]*)\]`)
	if match := tagsRegex.FindStringSubmatch(text); match != nil {
		tagStr := match[1]
		tags := strings.Split(tagStr, ",")
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				note.Tags = append(note.Tags, tag)
			}
		}
	}

	return note, nil
}

// GetAllNotes returns all notes with their metadata
func GetAllNotes() ([]*Note, error) {
	paths, err := ListNotes()
	if err != nil {
		return nil, err
	}

	var notes []*Note
	for _, path := range paths {
		note, err := ParseNote(path)
		if err == nil {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetNotesByTag returns notes that have a specific tag
func GetNotesByTag(tag string) ([]*Note, error) {
	allNotes, err := GetAllNotes()
	if err != nil {
		return nil, err
	}

	var filtered []*Note
	for _, note := range allNotes {
		for _, t := range note.Tags {
			if strings.EqualFold(t, tag) {
				filtered = append(filtered, note)
				break
			}
		}
	}

	return filtered, nil
}

// GetAllTags returns all unique tags across all notes
func GetAllTags() ([]string, error) {
	notes, err := GetAllNotes()
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]int)
	for _, note := range notes {
		for _, tag := range note.Tags {
			tagMap[tag]++
		}
	}

	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}

	return tags, nil
}

