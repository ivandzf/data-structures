#![allow(dead_code)]

fn main() {}

struct Stack<T> {
    data: Vec<T>,
}

impl<T> Stack<T> {
    fn new() -> Stack<T> {
        Stack { data: Vec::new() }
    }

    fn push(&mut self, elem: T) {
        self.data.push(elem);
    }

    fn pop(&mut self) -> Option<T> {
        self.data.pop()
    }

    fn peek(&self) -> Option<&T> {
        self.data.last()
    }

    fn size(&self) -> usize {
        self.data.len()
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
