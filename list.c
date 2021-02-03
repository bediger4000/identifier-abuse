/*
 * creates a linked list by using the linker to fill in references
 * to the "next" item in the list.
 *
 * Assembler must generate correct relocation information to items
 * in data segment (a.out) for this to work.
 *
 * Compiler has to be able to handle the identifier "list" as several
 * different things: 
 * 1. typedef name
 * 2. struct tag 
 * 3. struct member
 * 4. label
 * 5. actual identifier
 */

#include <stdio.h>

typedef struct list *list;  /* 1. typedef name  */

struct list {               /* 2. struct tag    */
        list list;          /* 3. struct member */
        int number;
};

/* linker abuse */
struct list xlist[] = {
  { &xlist[1], 0},
  { &xlist[2], 1},
  { &xlist[3], 2},
  { &xlist[4], 3},
  { &xlist[5], 4},
  { &xlist[6], 5},
  { NULL,      6}
};

int
main(int ac, char **av)
{
        list list;            /* 4. identifier of a variable */

        list = &xlist[0];

		list:                 /* 5. label */
		/* not "structured", but using the label tests if it's correct */

        printf("item at 0x%x has number %d, next is 0x%x\n",
                list, list->number, list->list);

        list = list->list;

        if (list) goto list;

        return 0;
}
