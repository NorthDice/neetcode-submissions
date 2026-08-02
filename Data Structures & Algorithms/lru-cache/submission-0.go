type Node struct {
    Key   int
    Value int
    Next  *Node
    Prev  *Node
}

type LRUCache struct {
    Cap  int
    Data map[int]*Node
    Head *Node
    Tail *Node
}

func Constructor(capacity int) LRUCache {
    data := make(map[int]*Node, capacity)
    head, tail := &Node{}, &Node{}
    head.Next = tail
    tail.Prev = head

    return LRUCache{
        Cap:  capacity,
        Data: data,
        Head: head,
        Tail: tail,
    }
}

func (this *LRUCache) remove(node *Node) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *Node) {
    node.Next = this.Head.Next
    node.Prev = this.Head
    this.Head.Next.Prev = node
    this.Head.Next = node
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.Data[key]
    if !ok {
        return -1
    }
    
    this.remove(node)
    this.addToHead(node)
    return node.Value
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.Data[key]; ok {
        this.remove(node)
        this.addToHead(node)
        node.Value = value
        return
    }
    
    if len(this.Data) >= this.Cap {
        tailNode := this.Tail.Prev
        this.remove(tailNode)
        delete(this.Data, tailNode.Key)
    }
    
    newNode := &Node{
        Key:   key,
        Value: value,
    }
    this.addToHead(newNode)
    this.Data[key] = newNode
}