package imascg

// SortableStringSlice is a type for sortable string slices
type SortableStringSlice []string

func (s SortableStringSlice) Len() int {
	return len(s)
}

func (s SortableStringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableStringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
