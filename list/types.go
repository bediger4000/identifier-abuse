package list

type list struct {
	list *list
}

func Xlist(n int) {
	/*
	 * This won't work:
	 * var list *list
	 * list = &list{}
	 */
	list := &list{} // type *list
list:
	for list.list != nil {
		/*
		 * There are 6, almost 7, things named "list" in this scope:
		 * 1. package list
		 * 2. type list
		 * 3. literal struct
		 * 4. element list of type list
		 * 5. label list
		 * 6. shadow variable list.
		 * and almost, the shadowed variable list,
		 * initialized outside the for-loop
		 */
		list := &list // inside for-loop variable list has type **list
		*list = (*list).list
		break list
	}
}
