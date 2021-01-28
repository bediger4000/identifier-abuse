package main

import "scope/list"

func main() {
	var list list.List
	xlist(&list)
}

func xlist(List *list.List) {
List:
	for List.List != nil {
		/*
		* There are 5, almost 6, things named "List" in this scope:
		* 1. type list.List
		* 2. element List of type List
		* 3. label .ist
		* 4. shadow variable List.
		* and almost, the formal argument variable named List
		 */
		List := &List // var list has type **list
		*List = (*List).List
		break List
	}
}
