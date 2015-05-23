package fuzzaldrin

import (
	"math"
	"sort"
)

type candidate struct {
	value string
	score float64
}

type sortableCandidates []*candidate

func (s sortableCandidates) Len() int {
	return len(s)
}

func (s sortableCandidates) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s sortableCandidates) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func filter(candidates []string, query string, queryHasSlases bool, maxResults int) (results []string) {
	sortedCandidates := sortableCandidates{}

	for _, str := range candidates {
		if str == "" {
			continue
		}
		calcScore := score(str, query)
		if !queryHasSlases {
			calcScore = basenameScore(str, query, calcScore)
		}
		if calcScore > 0 {
			c := &candidate{value: str, score: calcScore}
			sortedCandidates = append(sortedCandidates, c)
		}
	}

	sort.Sort(sortedCandidates)

	for _, c := range sortedCandidates {
		results = append(results, c.value)
	}

	if maxResults != -1 {
		maxResults = int(math.Min(float64(maxResults), float64(len(results))))
		results = results[0:maxResults]
	}

	return
}
