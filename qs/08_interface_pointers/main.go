// Boss Level: Pointers and Interfaces
//
// Task: You have an in-memory cache that saves strings. Because the `Save`
// method appends to the struct's internal slice, it MUST use a pointer receiver `(c *MemoryCache)`.
//
// If you run this code right now, it will fail to compile with this exact error:
// "MemoryCache does not implement StateSaver (Save method has pointer receiver)"
//
// 1. Look at `main`.
// 2. Fix the ONE line where `CommitData` is called so that it compiles and successfully
//    modifies the original cache.

package main

import "fmt"

type StateSaver interface {
	Save(data string)
}

type MemoryCache struct {
	Storage []string
}

// Notice the pointer receiver *MemoryCache. We need this to modify the slice!
func (c *MemoryCache) Save(data string) {
	c.Storage = append(c.Storage, data)
	fmt.Println("Saved to cache:", data)
}

// This function accepts the interface
func CommitData(s StateSaver, data string) {
	s.Save(data)
}

func main() {
	cache := MemoryCache{
		Storage: []string{},
	}

	// BUG: This line causes a compiler error!
	// How do you pass 'cache' so that it satisfies the pointer receiver?
	CommitData(&cache, "User login event")

	fmt.Println("Final Cache State:", cache.Storage)
}
