import { ListNode,LinkedList } from "./list";

import { LRUCache } from "./lru";
let list = new LinkedList<number>();

list.print();
list.pushToBack(1);
console.log(JSON.stringify(list.Head.Next.Value),JSON.stringify(list.Tail.Prev.Value))
list.print();
list.pushToBack(2);
console.log(JSON.stringify(list.Head.Next.Value),JSON.stringify(list.Tail.Prev.Value))
list.pushToBack(3);
list.pushToBack(4);
let first = list.front();
list.print();
list.remove(first)
list.print();
let first1 = list.front();
list.moveToBack(first1)

list.print();

// let c =new LRUCache<number>(2);
// console.log(c.get("1"))
// c.push("1",1)
// c.push("2",2)
// c.print();

// console.log(c.get("1"))
// c.print();

// console.log(c.get("2"))
// c.push("1",1)
// c.push("1",1)
// c.print();