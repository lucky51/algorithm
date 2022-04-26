export class ListNode<T> {
    Next:ListNode<T>;
    Prev:ListNode<T>;
    Value:T
}
export class LinkedList<T> {
    private count:number =0;
    Head :ListNode<T>;
    Tail:ListNode<T>;
    constructor(){
        this.Head= new ListNode<T>();
        this.Tail = new ListNode<T>();
        this.Head.Next=this.Tail;
        this.Tail.Prev=this.Head;
    }
    get Count(){
        return this.count;
    }
    pushToBack(value:T):ListNode<T> {
        let newNode=new ListNode<T>();
        newNode.Value=value;
        this.Tail.Prev.Next=newNode;
        newNode.Prev=this.Tail.Prev;
        newNode.Next=this.Tail;
        this.Tail.Prev =newNode;
        this.count++;
        return newNode;
    }
    pushToFront(value:T) :ListNode<T>{
        let newNode=new ListNode<T>();
        newNode.Value=value;
        let first =this.Head.Next;
        newNode.Next=first;
        first.Prev =newNode;
        newNode.Prev = this.Head;
        this.Head.Next=newNode;
        this.count++;
        return newNode;
    }
    front() :ListNode<T> |null {
        if(this.count==0){
            return null
        }
        return  this.Head.Next;
    }
    back(): ListNode<T>|null{
        return this.Tail.Prev;
    }
    remove(e:ListNode<T>|null) {
        if(this.count==0 || e==null){
            return
        }
        e.Prev.Next = e.Next;
        e.Next.Prev = e.Prev;
        this.count--;
    }
    moveToBack(e:ListNode<T>|null){
        if(e==null  ||e.Next==this.Tail ){
            return
        }
        e.Prev.Next = e.Next;
        e.Next.Prev = e.Prev;
        this.Tail.Prev.Next=e;
        e.Prev=this.Tail.Prev;
        e.Next=this.Tail;
    }
    print(){
        if(this.count==0){
            console.log(null);
            return;
        }
        let temp=this.Head
        let n = [];
        while(temp!=null) {
            n.push(temp.Value)
            temp = temp.Next;
        }
        console.log('front to back:',JSON.stringify(n));
        // back to front
        let temp1 = this.Tail;
        let p = [];
        while(temp!=null){
            p.push(temp1.Value);
            temp1 = temp1.Prev
        }
        console.log('back to front:',JSON.stringify(p));
    }
}
