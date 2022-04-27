
export class ListNode<T> {
    Next: ListNode<T>|null;
    Prev: ListNode<T>|null;
    Value: T
}
export enum PrintDirection {
    Normal,
    Reverse
}
export interface ILogger {
    log(message: any): void
}
export class LinkedList<T> {
    private count: number = 0;
    Head: ListNode<T>;
    Tail: ListNode<T>;
    constructor() {
        this.Head = new ListNode<T>();
        this.Tail = new ListNode<T>();
        this.Head.Next = this.Tail;
        this.Tail.Prev = this.Head;
    }
    get Count() {
        return this.count;
    }
    pushToBack(value: T): ListNode<T> {
        if(this.Tail.Prev==null){
            throw new Error("invalid list.");
        }
        let newNode = new ListNode<T>();
        newNode.Value = value;
        this.Tail.Prev.Next = newNode;
        newNode.Prev = this.Tail.Prev;
        newNode.Next = this.Tail;
        this.Tail.Prev = newNode;
        this.count++;
        return newNode;
    }
    pushToFront(value: T): ListNode<T> {
        let newNode = new ListNode<T>();
        newNode.Value = value;
        let first = this.Head.Next;
        if(first==null) {
            throw new Error("invalid list.")
        }
        newNode.Next = first;
        first.Prev = newNode;
        newNode.Prev = this.Head;
        this.Head.Next = newNode;
        this.count++;
        return newNode;
    }
    front(): ListNode<T> | null {
        if (this.count == 0) {
            return null
        }
        return this.Head.Next;
    }
    back(): ListNode<T> | null {
        return this.Tail.Prev;
    }
    remove(e: ListNode<T> | null) {
        if (this.count == 0 || e == null) {
            return
        }
        if(e.Prev==null||e.Next==null){
            throw new Error("invalid list node.");
        }
        e.Prev.Next = e.Next;
        e.Next.Prev = e.Prev;
        this.count--;
    }
    moveToBack(e: ListNode<T> | null) {
        if (e == null ||e.Prev==null||e.Next==null) {
            throw new Error("invalid list node");
        }
        if(this.Tail.Prev==null){
            throw new Error("invalid list.")
        }
        if(e.Next == this.Tail){
            return
        }
        e.Prev.Next = e.Next;
        e.Next.Prev = e.Prev;
        this.Tail.Prev.Next = e;
        e.Prev = this.Tail.Prev;
        e.Next = this.Tail;
        this.Tail.Prev = e;
    }
    print(logger: ILogger, direction: PrintDirection) {
        if (this.count == 0) {
            logger.log(null)
            return;
        }
        var temp:ListNode<T>|null = direction == PrintDirection.Normal ? this.Head : this.Tail
        var n = [];
        while (temp != null) {        
            n.push(temp.Value)
            temp =direction == PrintDirection.Normal? temp.Next: temp.Prev ;
        }
        logger.log(JSON.stringify(n))
    }
}


