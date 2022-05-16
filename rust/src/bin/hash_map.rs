// hash map
// under the hood, a hash map uses a dynamic array of linked lists to efficiently store key-value pairs
// when inserting a new key-value pair, a hash function first map the key to an integer value, it will be used for index in the underlying dynamic array
// the map checks if the key already exists. if it does, it overwrites the old value with the new one
// below is an example of a hash map look like under the hood:
// [
//	0: [key1, value1] -> null,
//	1: [key2, value2] -> [key3, value3] -> null,
//	3: [key4, value4] -> null,
//	4: [key5, value5] -> [key6, value6] -> [key7, value7] -> null,
// ]
// in the hash map above, the keys 2, 3 collided by all being hashed to 1, and the keys 4, 5, 6, 7 collided by all being hashed to 4
// - insert -> O(1) on average; O(n) on worst case
// - lookup -> O(1) on average; O(n) on worst case
// - delete -> O(1) on average; O(n) on worst case
// - iterate -> O(n) on average; O(n) on worst case
// - size -> O(1)

fn main() {}

#[cfg(test)]
mod tests {
    // use native map from golang, to complicated if we want to implement our own
    use std::collections::HashMap;

    #[test]
    fn init() {
        let map: HashMap<i32, i32> = HashMap::new();
        assert_eq!(map.len(), 0);
    }

    #[test]
    fn insert() {
        let mut map = HashMap::new();
        map.insert(1, 2);
        assert_eq!(map.len(), 1);
        assert_eq!(map.get(&1), Some(&2));
    }

    #[test]
    fn remove() {
        let mut map = HashMap::new();
        map.insert(1, 2);
        map.remove(&1);
        assert_eq!(map.len(), 0);
    }

    #[test]
    fn contains_key() {
        let mut map = HashMap::new();
        map.insert(1, 2);
        assert!(map.contains_key(&1));
        assert!(!map.contains_key(&2));
    }

    #[test]
    fn get() {
        let mut map = HashMap::new();
        map.insert(1, 2);
        assert_eq!(map.get(&1), Some(&2));
        assert_eq!(map.get(&2), None);
    }

    #[test]
    fn iter() {
        let mut map = HashMap::new();
        map.insert(1, 2);
        map.insert(2, 3);
        map.insert(3, 4);

        for (key, value) in &map {
            println!("{}: {}", key, value);
        }
    }
}
