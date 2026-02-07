package rules

import (
	"os"
	"strings"
)

type LnSOrderRule struct{}

func (r *LnSOrderRule) ID() string { return "ln_s_order" }

func (r *LnSOrderRule) Match(command string, output string) bool {
	return strings.HasPrefix(command, "ln ") &&
		(strings.Contains(command, "-s") || strings.Contains(command, "--symbolic")) &&
		strings.Contains(output, "File exists")
}

func (r *LnSOrderRule) GetNewCommand(command string, output string) string {
	parts := strings.Fields(command)
	var dest string
	var otherArgs []string

	for _, part := range parts {
		if part == "ln" || part == "-s" || part == "--symbolic" {
			otherArgs = append(otherArgs, part)
			continue
		}
		// Check if part exists
		if _, err := os.Stat(part); err == nil {
			if dest == "" {
				dest = part
			} else {
				// Multiple existing files?
				otherArgs = append(otherArgs, part)
			}
		} else {
			otherArgs = append(otherArgs, part)
		}
	}

	if dest != "" {
		// New command: ln -s [others] dest
		// Reconstruct. Order matters for others?
		// "ln -s dest link". dest exists.
		// "ln -s link dest". dest exists. link doesn't.
		// If command was "ln -s dest link" (correct), and "File exists" -> means link matches an existing file?
		// Then dest is the file we want to link TO.
		// If command was "ln -s link dest" (wrong). link doesn't exist. dest exists.
		// We found dest.
		// We want "ln -s dest link".
		// Implementation: remove dest from parts, append to end? No, `ln -s dest link` puts dest first.
		// Wait. `ln -s TARGET LINK_NAME`.
		// If I typed `ln -s LINK_NAME TARGET`. TARGET exists.
		// I want to swap them.
		// Python code: `parts.remove(destination); parts.append(destination)`.
		// `parts` is `['ln', '-s', 'link', 'dest']`.
		// `dest` removed -> `['ln', '-s', 'link']`.
		// `append(dest)` -> `['ln', '-s', 'link', 'dest']`.
		// This results in the SAME command if `dest` was already at end.
		// But in the wrong case: `ln -s link dest`. `dest` is at end.
		// If `dest` is at end, append puts it at end.
		// Maybe `thefuck` assumes `ln -s dest link` order?
		// `ln -s dest link` -> `dest` exists.
		// If I type `ln -s dest link` and it fails with "File exists", `link` must exist.
		// `thefuck` logic seems to be: if `ln -s A B` fails and A exists, try `ln -s A B`? No.
		// If `ln -s A B` fails and B exists.
		// Python: `_get_destination` checks if path exists.
		// If `ln -s A B` and B exists. `_get_destination` returns B.
		// New command: remove B, append B. -> `ln -s A B`. Same?
		// Unless `_get_destination` returns the FIRST one that exists?
		// `for part in script_parts: if exists: return part`.
		// It returns the FIRST existing part.
		// Case 1: `ln -s target link`. target exists. link does not.
		// `_get_destination` returns `target`.
		// New command: remove `target`, append `target`. -> `ln -s link target`.
		// SWAPPED!
		// Case 2: `ln -s link target`. link does not exist. target exists.
		// `_get_destination` returns `target`.
		// New command: remove `target`, append `target`. -> `ln -s link target`.
		// NO SWAP?
		// Wait. `['ln', '-s', 'link', 'target']`. remove `target`. `['ln', '-s', 'link']`. append `target`. `['ln', '-s', 'link', 'target']`.
		// The python code `parts.remove(destination)` removes first occurrence.
		// If `target` is at the end, it stays at the end.
		// So `thefuck` fixes `ln -s target link` -> `ln -s link target`?
		// That seems backwards. `ln -s target link` is correct (link to target).
		// Why would you want to swap it?
		// Maybe `ln -s link target` creates a link inside target?
		// If I intended to create `link` pointing to `target`.
		// I wrote `ln -s target link`. "File exists" means `link` already exists.
		// Maybe I meant `ln -s link target` (create link inside target)?
		// Or maybe I wrote `ln -s target link` and `target` is the Link Name? (No, first arg is target).

		// Let's assume the user made a mistake and `thefuck` logic works for them.
		// The logic is: Find the first existing argument. Move it to the end.

		parts := strings.Fields(command)
		var dest string
		destIndex := -1

		for i, part := range parts {
			if part == "ln" || part == "-s" || part == "--symbolic" {
				continue
			}
			if _, err := os.Stat(part); err == nil {
				dest = part
				destIndex = i
				break // Return first match
			}
		}

		if destIndex != -1 {
			// Remove at index
			parts = append(parts[:destIndex], parts[destIndex+1:]...)
			// Append to end
			parts = append(parts, dest)
			return strings.Join(parts, " ")
		}
	}
	return command
}
