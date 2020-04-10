package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"

	"github.com/tmdvs/Go-Emoji-Utils/utils"
)

// Emoji - Struct representing Emoji
type Emoji struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Descriptor string `json:"descriptor"`
}

// LoadEmojiFile will attempt to load and parse the provided full path to the emoji.json file
// In the event we fail to read the file or fail to unmarshal, we will return an error
func LoadEmojiFile(filepath string) (loadErr error) {
	var jsonFileContents []byte
	if jsonFileContents, loadErr = ioutil.ReadFile(filepath); loadErr != nil { // Attempt to read the file directly
		return // Return with the loadErr if we failed to read the file
	}

	var tempEmojis map[string]Emoji
	loadErr = json.Unmarshal(jsonFileContents, &tempEmojis)

	if len(tempEmojis) != 0 { // If we have emoji map contents
		Emojis = tempEmojis // Override the Emojis map
	}

	return
}

// LookupEmoji - Lookup a single emoji definition
func LookupEmoji(emojiString string) (emoji Emoji, err error) {

	hexKey := utils.StringToHexKey(emojiString)

	// If we have a definition for this string we'll return it,
	// else we'll return an error
	if e, ok := Emojis[hexKey]; ok {
		emoji = e
	} else {
		err = fmt.Errorf("No record for \"%s\" could be found", emojiString)
	}

	return emoji, err
}

// LookupEmojis - Lookup definitions for each emoji in the input
func LookupEmojis(emoji []string) (matches []interface{}) {
	for _, emoji := range emoji {
		if match, err := LookupEmoji(emoji); err == nil {
			matches = append(matches, match)
		} else {
			matches = append(matches, err)
		}
	}

	return
}

// RemoveAll - Remove all emoji
func RemoveAll(input string) string {

	// Find all the emojis in this string
	matches := FindAll(input)

	for _, item := range matches {
		emo := item.Match.(Emoji)
		rs := []rune(emo.Value)
		for _, r := range rs {
			input = strings.ReplaceAll(input, string([]rune{r}), "")
		}
	}

	// Remove and trim and left over whitespace
	return strings.TrimSpace(strings.Join(strings.Fields(input), " "))
	//return input
}
