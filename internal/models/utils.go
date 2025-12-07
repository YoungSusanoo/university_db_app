package models

func GroupsToNames(groups []Group) (strs []string) {
	strs = make([]string, len(groups))
	for i, group := range groups {
		strs[i] = group.Name
	}
	return
}
