package main

import "strings"

// WatchConfig is the configuration used to watch Github repositories
type Diff struct {
	Type DiffType
	File string
}

type DiffType string

const (
	DiffTypeDeleted  DiffType = "D"
	DiffTypeModified DiffType = "M"
	DiffTypeAppended DiffType = "A"
)

func parseDiffOutput(output []byte) []*Diff {
	outputStr := strings.TrimSpace(string(output))

	diffsStr := strings.Split(outputStr, "\n")

	diffs := make([]*Diff, 0)

	for _, diffStr := range diffsStr {
		parts := strings.Split(diffStr, "\t")

		if len(parts) != 2 {
			// TODO: proper error handling
			continue
		}

		diffType := DiffType(parts[0])

		if diffType != DiffTypeDeleted &&
			diffType != DiffTypeModified &&
			diffType != DiffTypeAppended {
			// TODO: proper error handling
			continue
		}

		diffs = append(diffs, &Diff{Type: diffType, File: parts[1]})
	}

	return diffs
}
