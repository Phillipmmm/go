// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helper functions exported for the compiler.
// Do not use internally.

package types2

// If t is a pointer, AsPointer returns that type, otherwise it returns nil.
func AsPointer(t Type) *Pointer {
	u, _ := t.Underlying().(*Pointer)
	return u
}

// If t is a signature, AsSignature returns that type, otherwise it returns nil.
func AsSignature(t Type) *Signature {
	u, _ := t.Underlying().(*Signature)
	return u
}

// If t is a type parameter, AsTypeParam returns that type, otherwise it returns nil.
func AsTypeParam(t Type) *TypeParam {
	u, _ := t.Underlying().(*TypeParam)
	return u
}

// If t is a type parameter, StructuralType returns the single underlying
// type of all types in the type parameter's type constraint if it exists,
// or nil otherwise. If t is not a type parameter, StructuralType returns
// the underlying type of t.
func StructuralType(t Type) Type {
	return structuralType(t)
}