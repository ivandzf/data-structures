#![allow(dead_code)]

fn main() {}

struct Queue<T> {
    element: Vec<T>,
}

impl<T> Queue<T> {
    pub fn new() -> Queue<T> {
        Queue { element: Vec::new() }
    }

    pub fn enqueue(&mut self, elem: T) {
        self.element.push(elem);
    }

    pub fn dequeue(&mut self) -> Option<T> {
        match self.element.get(0) {
            Some(_) => {
                Some(self.element.remove(0))
            }
            None => None,
        }
    }

    pub fn peek(&self) -> Option<&T> {
        self.element.first()
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
        let q: Queue<i32> = Queue::new();
        assert_eq!(q.size(), 0);
    }

    #[test]
    fn enqueue() {
        let mut q: Queue<i32> = Queue::new();
        q.enqueue(1);
        q.enqueue(2);
        assert_eq!(q.size(), 2);
    }

    #[test]
    fn dequeue() {
        let mut q: Queue<i32> = Queue::new();
        q.enqueue(1);
        assert_eq!(q.dequeue(), Some(1));
        assert_eq!(q.size(), 0);
    }

    #[test]
    fn peek() {
        let mut q: Queue<i32> = Queue::new();
        q.enqueue(1);
        assert_eq!(q.peek(), Some(&1));
    }

    #[test]
    fn size() {
        let mut q: Queue<i32> = Queue::new();
        q.enqueue(1);
        assert_eq!(q.size(), 1);
    }
}
