// Copyright (c) 2012, Martin Angers & Contributors
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation and/or
// other materials provided with the distribution.
// * Neither the name of the author nor the names of its contributors may be used to
// endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS
// OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
// WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
Package goquery implements features similar to jQuery, including the chainable
syntax, to manipulate and query an HTML document.

It depends on Go's experimental html package, which must be installed so that it
can be imported as "exp/html". See this tutorial on how to install it 
accordingly: http://code.google.com/p/go-wiki/wiki/InstallingExp

It uses Cascadia as CSS selector (similar to Sizzle for jQuery). This dependency
is automatically installed when using "go get ..." to install GoQuery.

To provide a chainable interface, error management is strict, and goquery panics
if an invalid Cascadia selector is used (this is consistent with the behavior of
jQuery/Sizzle/document.querySelectorAll, where an error is thrown). This is
necessary since multiple return values cannot be used to allow a chainable 
interface.

It is hosted on GitHub, along with additional documentation in the README.md
file: https://github.com/puerkitobio/goquery

The various methods are split into files based on the category of behavior:

* array.go : array-like positional manipulation of the selection.
    - First()
    - Last()
    - Eq()
    - Get()
    - Index...()
    - Slice()

* filter.go : filtering methods, that reduce the selection's set.
    - Filter...()
    - Not...()
    - Has...()
    - End()
    - Intersection(), which is an alias of FilterSelection()

* expand.go : methods that expand or augment the selection's set.
    - Add...()
    - AndSelf()
    - Union(), which is an alias for AddSelection()

* query.go : methods that query, or reflect, a node's identity.
    - Contains()
    - HasClass()
    - Is...()

* property.go : methods that inspect and get the node's properties values.
    - Attr()
    - Contents()
    - Html()
    - Length()
    - Size(), which is an alias for Length()
    - Text()
    - Val()

* traversal.go : methods to traverse the HTML document tree.
    - Children...()
    - Closest()
    - Find...()
    - Next...()
    - Parent[s]...()
    - Prev...()
    - Siblings...()

* iteration.go : methods to loop over the selection's nodes.
    - Each()
    - Map()

* type.go : definition of the types exposed by GoQuery.
    - Document
    - Selection
*/
package goquery

// DONE array.go : Positional Manipulation: First(), Last(), Eq(), Get(), Index(), Slice()
// DONE filter.go : Filtering: Filter(), Not(), Has(), End()
// DONE expand.go : "Expanding": Add(), AndSelf()
// DONE query.go : Reflect (query) node: Is(), Contains(), HasClass()
// property.go : Inspect node: Contents(), Html(), Text(), Attr(), ** Does it make sense in a static HTML tree? Val()**, Length(), Size()
// traversal.go : Traversal: Find(), Children(), Parents...(), Next...(), Prev...(), Closest(), Siblings()
// DONE iteration.go : Iteration: Each(), Map()
// DONE type.go : Selection and Document

// TODO : Benchmarks

// TODO : Check each method, if it applies to any node or only Element nodes (Cascadia's selectors already make sure of that)

// TODO : Add the following methods:
// x Add()
// x AndSelf()
// x Attr()
// x Children()
// - Closest() - Tree Traversal
// x Contains()
// - Contents() (similar to Children(), but includes text and comment nodes, so Children() should filter them out) - Misc. Traversing
// x Each()
// x End()
// x Eq()
// x Filter()
// - Find() : Complete with Selection object and Node object as selectors - Tree Traversal
// x First()
// x Get()
// x Has()
// x HasClass()
// - Html() ? - Attributes
// x Index()
// - Is() - Filtering
// x Last()
// x Length() / Size()
// x Map()
// - Next() - Tree traversal
// - NextAll() - Tree traversal
// - NextUntil() - Tree traversal
// x Not()
// - Parent() - Tree traversal
// - Parents() - Tree traversal
// - ParentsUntil() - Tree traversal
// - Prev() - Tree traversal
// - PrevAll() - Tree traversal
// - PrevUntil() - Tree traversal
// x PushStack()
// - Siblings() - Tree traversal
// x Slice()
// - Text() - DOM Manipulation
// x ToArray()
// x Unique() internally only
// - Val() ? - Attributes