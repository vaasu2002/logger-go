package filter

import (
	"logger/pkg/logging"
	"strings"
)

// KeywordFilter filters logs based on keyword presence in the message
type KeywordFilter struct {
	keyword string
}

func NewKeywordFilter(keyword string) *KeywordFilter {
	return &KeywordFilter{
		keyword: keyword,
	}
}

func (k *KeywordFilter) Accept(log *logging.Log) bool {
	return strings.Contains(strings.ToLower(log.Message), strings.ToLower(k.keyword))
}
