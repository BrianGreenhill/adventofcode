package main

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndexStr(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
