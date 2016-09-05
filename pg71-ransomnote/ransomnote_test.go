package ransomnote

import (
	"testing"
)

func TestIsRansomNote(t *testing.T) {
	body := `A ransom note can be formed by cutting words out of a magazine to form a new sentence.
How would you figure out if a ransom note (represented as a string) can be formed from a given magazine (string)?`
	note := "how can you form a ransom note from a magazine magazine"
	actual := isRansomNote(note, body)
	if actual != true {
		t.Errorf("want: true got: false")
	}
}
