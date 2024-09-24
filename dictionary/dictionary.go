package dictionary

import (
	"reflect"
)

// Dictionary is a type alias for a generic map.
type Dictionary[K comparable, V any] map[K]V

// CreateDictionary creates a Dictionary with a single key-value pair.
// It takes a key of any comparable type and a value of any type,
// and returns a Dictionary containing the provided key and value.
//
// Parameters:
//   - key: The key of the Dictionary entry.
//   - value: The value associated with the key.
//
// Returns:
//   - A Dictionary containing the provided key and value.
//
// Example:
//
//	dict := CreateDictionary("one", 1)
//	fmt.Println(dict) // Output: map[one:1]
func CreateDictionary[K comparable, V any](key K, value V) Dictionary[K, V] {
	return Dictionary[K, V]{
		key: value,
	}
}

// DefaultDictionary creates an empty Dictionary.
// It returns a Dictionary with no key-value pairs.
//
// Returns:
//   - A new empty Dictionary.
func DefaultDictionary[K comparable, V any]() Dictionary[K, V] {
	return Dictionary[K, V]{}
}

// GetValue retrieves the value associated with the specified key from the Dictionary.
//
// Parameters:
//   - key: The key whose associated value is to be returned.
//
// Returns:
//   - The value associated with the specified key.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	value := dict.GetValue("one") // value will be 1
func (d Dictionary[K, V]) GetValue(key K) V {
	return d[key]
}

// SetValue sets the value for a given key in the Dictionary.
// If the key already exists, its value will be updated.
//
// Parameters:
//   - key: The key for which the value needs to be set.
//   - value: The value to be set for the given key.
//
// Example:
//
//	dict := Dictionary[string, int]{}
//	dict.SetValue("one", 1)
//	dict.SetValue("two", 2)
//	value := dict.GetValue("one") // value will be 1
//	dict.DeleteValue("one")
//	keys := dict.GetKeys() // keys will be ["two"]
//	values := dict.GetValues() // values will be [2]
func (d Dictionary[K, V]) SetValue(key K, value V) {
	d[key] = value
}

// DeleteValue removes the value associated with the specified key from the Dictionary.
// If the key does not exist, the Dictionary remains unchanged.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	dict.DeleteValue("one")
//	// dict is now Dictionary[string, int]{"two": 2}
//
//	dict.DeleteValue("three")
//	// dict remains Dictionary[string, int]{"two": 2} as "three" does not exist
func (d Dictionary[K, V]) DeleteValue(key K) {
	delete(d, key)
}

// GetKeys returns a slice containing all the keys present in the Dictionary.
// It iterates over the Dictionary and collects each key into a slice, which is then returned.
//
// Returns:
//   - []K: A slice of keys of type K.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	keys := dict.GetKeys() // keys will be ["one", "two"]
func (d Dictionary[K, V]) GetKeys() []K {
	keys := make([]K, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}

// GetValues returns a slice containing all the values present in the Dictionary.
// It iterates over the Dictionary and collects each value into a slice, which is then returned.
//
// Returns:
//   - []V: A slice of values of type V.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	values := dict.GetValues() // values will be [1, 2]
func (d Dictionary[K, V]) GetValues() []V {
	values := make([]V, 0, len(d))
	for _, v := range d {
		values = append(values, v)
	}
	return values
}

// GetLength returns the number of key-value pairs present in the Dictionary.
//
// Returns:
//   - int: The number of key-value pairs in the Dictionary.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	length := dict.GetLength() // length will be 2
func (d Dictionary[K, V]) GetLength() int {
	return len(d)
}

// MergeDictionaries merges another Dictionary into the current Dictionary.
// If there are duplicate keys, the values from the other Dictionary will overwrite the current values.
//
// Parameters:
//   - d2: The Dictionary to be merged into the current Dictionary.
//
// Example:
//
//	dict1 := Dictionary[string, int]{"one": 1, "two": 2}
//	dict2 := Dictionary[string, int]{"three": 3, "four": 4}
//	dict1.MergeDictionaries(dict2)
//	// dict1 is now Dictionary[string, int]{"one": 1, "two": 2, "three": 3, "four": 4}
func (d Dictionary[K, V]) MergeDictionaries(d2 Dictionary[K, V]) {
	for k, v := range d2 {
		d[k] = v
	}
}

// CopyDictionary returns a copy of the current Dictionary.
//
// Returns:
//   - Dictionary[K, V]: A copy of the current Dictionary.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	copy := dict.CopyDictionary()
//	// copy is Dictionary[string, int]{"one": 1, "two": 2}
func (d Dictionary[K, V]) CopyDictionary() Dictionary[K, V] {
	copy := make(Dictionary[K, V])
	for k, v := range d {
		copy[k] = v
	}
	return copy
}

// ContainsKey checks if the Dictionary contains the specified key.
//
// Parameters:
//   - key: The key to be checked.
//
// Returns:
//   - bool: True if the key is present, false otherwise.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	contains := dict.ContainsKey("one") // contains will be true
func (d Dictionary[K, V]) ContainsKey(key K) bool {
	_, ok := d[key]
	return ok
}

// ContainsValue checks if the Dictionary contains the specified value.
//
// Parameters:
//   - value: The value to be checked.
//
// Returns:
//   - bool: True if the value is present, false otherwise.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	contains := dict.ContainsValue(1) // contains will be true
func (d Dictionary[K, V]) ContainsValue(value V) bool {
	for _, v := range d {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// ClearDictionary removes all key-value pairs from the Dictionary.
//
// Example:
//
//	dict := Dictionary[string, int]{"one": 1, "two": 2}
//	dict.ClearDictionary()
//	// dict is now empty
func (d Dictionary[K, V]) ClearDictionary() {
	for k := range d {
		delete(d, k)
	}
}

// IsEmpty checks if the Dictionary is empty.
//
// Returns:
//   - bool: True if the Dictionary is empty, false otherwise.
//
// Example:
//
//	dict := Dictionary[string, int]{}
//	empty := dict.IsEmpty() // empty will be true
func (d Dictionary[K, V]) IsEmpty() bool {
	return len(d) == 0
}

// IsEqual checks if the current Dictionary is equal to another Dictionary.
// Two Dictionaries are considered equal if they have the same key-value pairs.
//
// Parameters:
//   - d2: The Dictionary to be compared with.
//
// Returns:
//   - bool: True if the Dictionaries are equal, false otherwise.
//
// Example:
//
//	dict1 := Dictionary[string, int]{"one": 1, "two": 2}
//	dict2 := Dictionary[string, int]{"one": 1, "two": 2}
//	equal := dict1.IsEqual(dict2) // equal will be true
func (d Dictionary[K, V]) IsEqual(d2 Dictionary[K, V]) bool {
	if len(d) != len(d2) {
		return false
	}
	for k, v := range d {
		if !reflect.DeepEqual(v, d2[k]) {
			return false
		}
	}
	return true
}

// IsSubset checks if the current Dictionary is a subset of another Dictionary.
// A Dictionary is considered a subset if all its key-value pairs are present in the other Dictionary.
//
// Parameters:
//   - d2: The Dictionary to be compared with.
//
// Returns:
//   - bool: True if the current Dictionary is a subset, false otherwise.
//
// Example:
//
//	dict1 := Dictionary[string, int]{"one": 1}
//	dict2 := Dictionary[string, int]{"one": 1, "two": 2}
//	subset := dict1.IsSubset(dict2) // subset will be true
func (d Dictionary[K, V]) IsSubset(d2 Dictionary[K, V]) bool {
	if len(d) > len(d2) {
		return false
	}
	for k, v := range d {
		if !reflect.DeepEqual(v, d2[k]) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the current Dictionary is a superset of another Dictionary.
// A Dictionary is considered a superset if it contains all the key-value pairs of the other Dictionary.
//
// Parameters:
//   - d2: The Dictionary to be compared with.
//
// Returns:
//   - bool: True if the current Dictionary is a superset, false otherwise.
//
// Example:
//
//	dict1 := Dictionary[string, int]{"one": 1, "two": 2}
//	dict2 := Dictionary[string, int]{"one": 1}
//	superset := dict1.IsSuperset(dict2) // superset will be true
func (d Dictionary[K, V]) IsSuperset(d2 Dictionary[K, V]) bool {
	if len(d) < len(d2) {
		return false
	}
	for k, v := range d2 {
		if !reflect.DeepEqual(v, d[k]) {
			return false
		}
	}
	return true
}

// IsDisjoint checks if the current Dictionary is disjoint with another Dictionary.
// Two Dictionaries are considered disjoint if they have no keys in common.
//
// Parameters:
//   - d2: The Dictionary to be compared with.
//
// Returns:
//   - bool: True if the Dictionaries are disjoint, false otherwise.
//
// Example:
//
//	dict1 := Dictionary[string, int]{"one": 1}
//	dict2 := Dictionary[string, int]{"two": 2}
//	disjoint := dict1.IsDisjoint(dict2) // disjoint will be true
func (d Dictionary[K, V]) IsDisjoint(d2 Dictionary[K, V]) bool {
	for k := range d {
		if _, ok := d2[k]; ok {
			return false
		}
	}
	return true
}
