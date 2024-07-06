package main
import (
	"fmt"
)
// Node repsents a single node in the linked list 
type Node struct {
	data int
	next *Node 
} 

 // LinkedList represents the Linked List
	 type LinkedList struct {
		head *Node
	 }

	 // InsertAtBeginning inserts a new node with given data at the beginning of the list
	  func(ll *LinkedList) InsertAtBeginning(data int) {
		newNode := &Node{
			data : data,
			next: ll.head,
		}
		ll.head = newNode
	  }
	  // InsertAfter insets a new node with give data after a node with specified value
	  func(ll *LinkedList) InsertAfter(value int, newData int) {
			newNode := &Node{
				data: newData,
				next: nil,
			}

			current:= ll.head
			for current != nil {
				if current.data == value {
					newNode.next = current.next
					current.next = newNode 
					return
				}
				current = current.next
			}
			fmt.Println("Node with value", value, "not found in the list.")
	  }
	  // InsertAtEnd inserts a new node with given data data at the of the list
	  func(ll *LinkedList) InsertAtEnd(data int) {
			newNode := &Node{
					data: data,
					next: nil,
			}
			if ll.head == nil {
				ll.head = newNode
				return
			}
			last := ll.head
			for last.next != nil {
					last = last.next
			}
			last.next = newNode
	  }
	  // DeleteAtBeginning deletes the node at beginning of the list
	  func (ll *LinkedList) DeleteAtBeginning() {
		if ll.head == nil {
			fmt.Println("Linked List is empty. Nothing to delete.")
			return
		}
		ll.head = ll.head.next
	  }
	  // Deleteafter deletes the node after a node with specified value
	  func (ll *LinkedList) DeleteAfter(value int) {
		current := ll.head
		for current != nil && current.next != nil {
			if current.data == value {
				current.next = current.next.next
				return
			}
			current = current.next
		}
		fmt.Println("Node with value", value, "Not found in the list")
	  } 
	  // DeleteAtEnd delets the node at the end of the list
	  func (ll *LinkedList) DeleteAtEnd() {
		if ll.head == nil {
			fmt.Println("Linked list is empty. Nothing to delete.")
			return
		}
		if ll.head.next == nil {
			ll.head = nil
			return
		}
		current := ll.head
		for current.next.next !=nil {
			current = current.next
		}
		current.next = nil
	  }
	  // ModifyNodeValue modifies the value of a node with specified value to a new value
	  func(ll *LinkedList) ModifyNodeValue(oldValue int, newValue int) {
		current := ll.head
		for current != nil {
			if current.data == oldValue {
				current.data = newValue
				return
			}
			current = current.next
		}
		fmt.Println("Node with value", oldValue, "not found in the List.")
	  }
	  // PrintList prints all elements in the Linked List
	  func(ll *LinkedList) PrintList() {
		current := ll.head 
		for current != nil {
			fmt.Printf("%d -> ", current.data)
			current = current.next
		}
		fmt.Println("nil")
	  }
	  func main() {
		var ll LinkedList 
		// Insert at beginning
		ll.InsertAtBeginning(3)
		ll.InsertAtBeginning(2)
		ll.InsertAtBeginning(1)

		// Print List
		fmt.Println("Linked List after inser at beginning:")
		ll.PrintList()

		// Insert in the middle
		ll.InsertAfter(5, 3)
			// Print list
			fmt.Println("Linked List after insert at beginning:")
			ll.PrintList()
			// Insert in the midle
			ll.InsertAfter(1, 2)
			//Print list
			fmt.Println("Linked List after insert in the middle:")
			ll.PrintList()
			// Insert at end
			ll.InsertAtEnd(10)
			// Print List
			fmt.Println("Linked List after insert at end:")
			ll.PrintList()
			//Delete at beginning
			ll.DeleteAtBeginning()
			// Print List
			fmt.Println("Linked List after delete at beginning:")
			ll.PrintList()
			// delete in the middle
			ll.DeleteAfter(3)

			//Print list
			fmt.Println("Linked List after delete in the middle:")
			ll.PrintList()
			//Delete at end
			ll.DeleteAtEnd()
			// Print list
			fmt.Println("Linked List after delete at end:")
			ll.PrintList()
			// Modify node value
			ll.ModifyNodeValue(1, 7)
			// Print list
			fmt.Println("Linked List after modify node value:")
			ll.PrintList()
		}
	  