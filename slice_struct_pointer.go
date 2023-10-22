package main

type Student struct {
	Name     string
	Email    string
	Password string
}

func addData(slicePtr *[]Student, name string, email string, password string) {
	student := Student{Name: name, Email: email, Password: password}
	*slicePtr = append(*slicePtr, student)
	//fmt.Println(*slicePtr)
}

func updateNameData(slicePtr *[]Student, name string, updateData string) {
	for i := range *slicePtr {
		if (*slicePtr)[i].Name == name {
			(*slicePtr)[i].Name = updateData
			break
		}
	}
}

func getDataByName(slicePtr *[]Student, name string) *Student {
	for i := range *slicePtr {
		if (*slicePtr)[i].Name == name {
			return &(*slicePtr)[i]
			break
		}
	}
	return nil
}
func deletetDataByName(slicePtr *[]Student, name string) {
	for i := range *slicePtr {
		if (*slicePtr)[i].Name == name {
			*slicePtr = append((*slicePtr)[:i], (*slicePtr)[i+1:]...)
			break
		}
	}
}
