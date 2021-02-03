# Abuse of Go Identifiers

I got curious about how many ways a Go program can use
a single identifier in a single scope.
This repo represents my attempt at a maximum number of ways
a single identifer can appear in a scope and still have a
program that compiles.

I got to 6, arguably 7.
I had to use a package to get the 6th use of an identifier.
I chose the identifier "list".

1. package list
2. type list
3. literal struct of type list
4. element list of type list
5. label list
6. shadow variable list
7. shadowed variable list

Inside [list/types.go], `func Xlist` has all 7 uses:

```go
package list // 1. package

type list struct { // 2. type
	list *list // 3. element of type
}

func Xlist(n int) {
	list := &list{}  // 4. Shadowed variable, 5. list struct literal
list: // 6. label
	for list.list != nil {
		list := &list // 7. shadowing variable, type **list
		*list = (*list).list
		break list
	}
}
```

Sure, there's a little cheating going on,
the struct literal and the label appear strictly
out of the scope of the for-loop.
But the label does get used inside the for-loop-scope.

The `list := &list{}` would seem exactly equivalent to

```go
   var list **list
   list = &list{}
```

but this 2-line version causes compiler errors.
I conjecture that there are hidden problems with ":="
shortcut creation and initialization,
but I can't figure out what they may be.

---

## C language version

[C identifier abuse](list.c)

I only got to 5 identifers of different things in C89 C.

1. typedef name
2. struct tag 
3. struct member
4. label
5. variable

I'm not sure this is worth doing in Python or PHP,
and I'm revolted by C++ and Java.
