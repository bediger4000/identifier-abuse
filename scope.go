package main

type list struct {
	list *list
}

func main() {
	var list list
	xlist(&list)
}

func xlist(list *list) {
list:
	for list.list != nil {
		/*
		* There are 4, almost 5, things named "list" in this scope:
		* 1. type list
		* 2. element list of type list
		* 3. label list
		* 4. shadow variable list.
		* and almost, the formal argument variable named list
		 */
		list := &list // var list has type **list
		*list = (*list).list
		break list
	}
}
