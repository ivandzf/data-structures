#![allow(dead_code)]

fn main() {}

struct Stack<T> {
    element: Vec<T>,
}

impl<T> Stack<T> {
    pub fn new() -> Stack<T> {
        Stack { element: Vec::new() }
    }

    pub fn push(&mut self, elem: T) {
        self.element.push(elem);
    }

    pub fn pop(&mut self) -> Option<T> {
        self.element.pop()
    }

    pub fn peek(&self) -> Option<&T> {
        self.element.last()
    }

    pub fn size(&self) -> usize {
        self.element.len()
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn init() {
        let stack: Stack<i32> = Stack::new();
        assert_eq!(stack.size(), 0);
    }

    #[test]
    fn push() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.size(), 1);
    }

    #[test]
    fn pop() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.pop(), Some(1));
        assert_eq!(stack.size(), 0);
    }

    #[test]
    fn peek() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.peek(), Some(&1));
    }

    #[test]
    fn size() {
        let mut stack = Stack::new();
        stack.push(1);
        assert_eq!(stack.size(), 1);
    }
}
