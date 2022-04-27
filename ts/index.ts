import { ILogger, PrintDirection } from "./list";
import { LRUCache } from "./lru";
import logger from '@ptkdev/logger'

let log = new logger();

class CustomeLogger implements ILogger {
    log(message: any): void {
        log.warning("===================================begin=======================================");
        log.info(message);
        log.warning("===================================e n d=======================================")
    }
}
let clogger = new CustomeLogger;

let c = new LRUCache<number,number>(4);
c.get(1)
c.push(1, 1)
c.push(2, 2)
c.print(clogger, PrintDirection.Normal);
c.print(clogger,PrintDirection.Reverse);
console.log(c.get(3))
c.push(4,4)
c.push(6,6)
c.print(clogger,PrintDirection.Normal)
// list.print();
// list.pushToBack(1);
// console.log(JSON.stringify(list.Head.Next.Value),JSON.stringify(list.Tail.Prev.Value))
// list.print();
// list.pushToBack(2);
// console.log(JSON.stringify(list.Head.Next.Value),JSON.stringify(list.Tail.Prev.Value))
// list.pushToBack(3);
// list.pushToBack(4);
// let first = list.front();
// list.print();
// list.remove(first)
// list.print();
// let first1 = list.front();
// list.moveToBack(first1)

// list.print();
