package bible

type Specification interface {
	IsSatisfied(v Verse) bool
}

type BookSpecification struct {
	Book string
}

type ChapterSpecification struct {
	Chapter int
}

type BetterFilter struct{}

func (b BookSpecification) IsSatisfied(v Verse) bool {
	return v.Book == b.Book
}

func (c ChapterSpecification) IsSatisfied(v Verse) bool {
	return v.Chapter == c.Chapter
}

func (f *BetterFilter) Filter(verses []Verse, spec Specification) []Verse {
	result := make([]Verse, 0)
	for i, v := range verses {
		if spec.IsSatisfied(v) {
			result = append(result, verses[i])
		}
	}
	return result
}
