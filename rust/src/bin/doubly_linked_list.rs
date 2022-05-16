#![allow(dead_code)]

use std::{
    fmt::{Debug, Display, Formatter, Result},
    usize,
};

#[derive(Debug)]
struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>,
    prev: Option<Box<Node<T>>>,
}

impl<T> Node<T> {
    fn new(value: T) -> Self {
        Node {
            value,
            next: None,
            prev: None,
        }
    }
}

impl<T: Debug> Display for Node<T>
where
    T: Display,
{
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", self.value)
    }
}

struct LinkedList<T> {
    head: Option<Node<T>>,
    count: usize,
}

impl<T: Debug> LinkedList<T> {
    fn new() -> Self {
        LinkedList {
            head: None,
            count: 0,
        }
    }

    fn insert_last(&mut self, v: T) {
        //let mut new_node = Node::new(v);
        //match self.head {
        //    None => {
        //        self.head = Some(new_node);
        //    }
        //    Some(tail) => {}
        //}
        self.count += 1;
    }

    // traverse and get the linked list by number of index
    fn get_next_nth_node(&mut self, n: usize) -> Option<&mut Node<T>> {
        let mut nth_node = self.head.as_mut();
        for _ in 0..n {
            nth_node = match nth_node {
                None => return None,
                Some(node) => node.next.as_mut().map(|node| &mut **node),
            }
        }

        nth_node
    }

    fn display(&self) {
        if self.count < 1 {
            println!("Linked List is empty !!");
            return;
        }

        println!(
            "Current length: {}, Linked List: {:?}",
            self.count, self.head
        );
    }
}

fn main() {}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn init() {
        let list: LinkedList<i32> = LinkedList::new();
        assert_eq!(list.count, 0);
    }
}
