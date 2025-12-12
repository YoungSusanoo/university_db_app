package models

import "fmt"

func GroupsToStrings(groups []Group) (strs []string) {
	strs = make([]string, len(groups))
	for i, group := range groups {
		strs[i] = group.Name
	}
	return
}

func TeachersToStrings(teachers []Teacher) (strs []string) {
	strs = make([]string, len(teachers))
	for i, teach := range teachers {
		strs[i] = fmt.Sprintf("%s %s %s", teach.FirstName, teach.LastName, teach.FatherName)
	}
	return
}

func StudentsToStrings(students []Student) (strs []string) {
	strs = make([]string, len(students))
	for i, stud := range students {
		strs[i] = fmt.Sprintf("%s %s %s %s", stud.FirstName, stud.LastName, stud.FatherName, stud.Group)
	}
	return
}

func SubjectsToStrings(subjects []Subject) (strs []string) {
	strs = make([]string, len(subjects))
	for i, subj := range subjects {
		strs[i] = subj.Name
	}
	return
}
