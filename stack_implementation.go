
// 1. stack

package main
 
import (
	"fmt"
)
 
type item struct {
	value interface{} //value as interface type to hold any data type
	next  *item
}
 
type Stack struct {
	top  *item
	size int
}
 
func (stack *Stack) Len() int {
	return stack.size
}
 
func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}
 
func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
 
	return nil
}
 
func main() {
	stack := new(Stack)
	// Push different data type to the stack
	stack.Push(1)
	stack.Push("prasanthi")
	stack.Push(4.0)
 
	// Pop until stack is empty
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}

$go run main.go
4
prasanthi
1


// 2. stack operations

package main

import (
    "container/list"
    "fmt"
)

type customStack struct {
    stack *list.List
}

func (c *customStack) Push(value string) {
    c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
    if c.stack.Len() > 0 {
        ele := c.stack.Front()
        c.stack.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
    if c.stack.Len() > 0 {
        if val, ok := c.stack.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customStack) Size() int {
    return c.stack.Len()
}

func (c *customStack) Empty() bool {
    return c.stack.Len() == 0
}

func main() {
    customQueue := &customStack{
        stack: list.New(),
    }
    fmt.Printf("Push: A\n")
    customQueue.Push("A")
    fmt.Printf("Push: B\n")
    customQueue.Push("B")
    fmt.Printf("Size: %d\n", customQueue.Size())
    for customQueue.Size() > 0 {
        frontVal, _ := customQueue.Front()
        fmt.Printf("Front: %s\n", frontVal)
        fmt.Printf("Pop: %s\n", frontVal)
        customQueue.Pop()
    }
    fmt.Printf("Size: %d\n", customQueue.Size())
}

$go run main.go
Push: A
Push: B
Size: 2
Front: B
Pop: B
Front: A
Pop: A
Size: 0

// 3.Slice Implementation

package main

import (
    "fmt"
    "sync"
)

type customQueue struct {
    stack []string
    lock  sync.RWMutex
}

func (c *customQueue) Push(name string) {
    c.lock.Lock()
    defer c.lock.Unlock()
    c.stack = append(c.stack, name)
}

func (c *customQueue) Pop() error {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.stack = c.stack[:len-1]
        return nil
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        return c.stack[len-1], nil
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
    return len(c.stack)
}

func (c *customQueue) Empty() bool {
    return len(c.stack) == 0
}

func main() {
    customQueue := &customQueue{
        stack: make([]string, 0),
    }
    fmt.Printf("Push: A\n")
    customQueue.Push("A")
    fmt.Printf("Push: B\n")
    customQueue.Push("B")
    fmt.Printf("Size: %d\n", customQueue.Size())
    for customQueue.Size() > 0 {
        frontVal, _ := customQueue.Front()
        fmt.Printf("Front: %s\n", frontVal)
        fmt.Printf("Pop: %s\n", frontVal)
        customQueue.Pop()
    }
    fmt.Printf("Size: %d\n", customQueue.Size())
}


$go run main.go
Push: A
Push: B
Size: 2
Front: B
Pop: B
Front: A
Pop: A
Size: 0