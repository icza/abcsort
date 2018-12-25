/*

Package abcsort is a string sorting library that uses a custom, user-defined
alphabet.

Implementation does not convert the input strings into byte or rune slices, so
performance is rather good.

Custom sorting can be easiest achieved by using the Sorter helper type.


abcsort provides the essence of sorting: the implementation of a less() function
required by the standard lib's sort package.

*/
package abcsort
