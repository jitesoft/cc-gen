package gitwrapper

import "sort"

type sortTagsBy func(t1 *Tag, t2 *Tag) bool

func (by sortTagsBy) Sort(tags []*Tag) {
	sorter := &tagSorter{
		tags: tags,
		by:   by,
	}

	sort.Sort(sorter)
}

type tagSorter struct {
	tags []*Tag
	by func(t1, t2 *Tag) bool
}

func (t *tagSorter) Len () int {
	return len(t.tags)
}

func (t *tagSorter) Swap(i, j int) {
	t.tags[i], t.tags[j] = t.tags[j], t.tags[i]
}

func (t *tagSorter) Less(i, j int) bool {
	return t.by(t.tags[i], t.tags[j])
}
