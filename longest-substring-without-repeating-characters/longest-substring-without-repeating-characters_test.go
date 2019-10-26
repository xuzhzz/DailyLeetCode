package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLS(t *testing.T)  {
//	s:="abc"
	assert.Equal(t, 3, LengthOfLongestSubstring("abcabcbb"))
}
