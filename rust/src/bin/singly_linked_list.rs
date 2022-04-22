#![allow(dead_code)]

use std::{
    fmt::{Debug, Display, Formatter, Result},
    usize,
};

#[derive(Debug, PartialEq)]
struct Node<T> {
    element: T,
    next: Option<Box<Node<T>>>,
}

impl<T: Debug> Display for Node<T>
where
    T: Display,
{
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", self.element)
    }
}

impl<T> Node<T> {
    fn new(v: T) -> Self {
        Node {
            element: v,
            next: None,
        }
    }
}

#[derive(Debug)]
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

    // need traverse operation to get the tail, O(n)
    fn insert_last(&mut self, v: T) {
        let new_node = Node::new(v);
        match self.head {
            None => {
                self.head = Some(new_node);
            }
            Some(_) => {
                self.get_next_nth_node(self.count - 1)
                    .map(|node| node.next = Some(Box::new(new_node)));
            }
        }

        // manually count, to reduce calculate the lenght of linked list
        self.count += 1
    }

    // directly set on the head, O(1)
    fn insert_first(&mut self, v: T) {
        let mut new_node = Node::new(v);
        match self.head {
            None => self.head = Some(new_node),
            Some(_) => {
                let current_head = self.head.take().unwrap();
                new_node.next = Some(Box::new(current_head));
                self.head = Some(new_node);
            }
        }

        // manually count, to reduce calculate the lenght of linked list
        self.count += 1
    }

    // need traverse operation to get the tail, O(n)
    fn remove_last(&mut self) {
        match &self.head {
            None => {}
            Some(_) => {
                if self.count == 1 {
                    self.head = None;
                } else {
                    self.get_next_nth_node(self.count - 2)
                        .map(|node| node.next = None);
                }
                // manually count, to reduce calculate the lenght of linked list
                self.count -= 1
            }
        }
    }

    // directly take on the head, O(1)
    fn remove_first(&mut self) {
        self.head
            .take()
            .map(|head| self.head = head.next.map(|node| *node));
        // manually count, to reduce calculate the lenght of linked list
        self.count -= 1
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
            self.count,
            self.head.as_ref().unwrap()
        );
    }
}

fn main() {}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn init() {
        let ll: LinkedList<i32> = LinkedList::new();
        assert_eq!(ll.count, 0);
    }

    #[test]
    fn insert() {
        let mut ll: LinkedList<i32> = LinkedList::new();

        // insert last
        ll.insert_last(50);
        assert_eq!(ll.count, 1);
        assert_eq!(ll.get_next_nth_node(0).unwrap().element, Node::new(50).element);

        ll.display();

        // insert first
        ll.insert_first(100);
        assert_eq!(ll.count, 2);
        assert_eq!(ll.get_next_nth_node(0).unwrap().element, Node::new(100).element);
        assert_eq!(ll.get_next_nth_node(1).unwrap().element, Node::new(50).element);

        ll.display();
    }

    #[test]
    fn remove() {
        let mut ll: LinkedList<i32> = LinkedList::new();

        // insert last
        ll.insert_last(20);
        assert_eq!(ll.count, 1);
        assert_eq!(ll.get_next_nth_node(0).unwrap().element, Node::new(20).element);

        // remove last
        ll.remove_last();
        assert_eq!(ll.count, 0);

        ll.display();

        // insert first
        ll.insert_first(90);
        assert_eq!(ll.count, 1);
        assert_eq!(ll.get_next_nth_node(0).unwrap().element, Node::new(90).element);

        // remove first
        ll.remove_first();
        assert_eq!(ll.count, 0);

        ll.display();
    }
}
