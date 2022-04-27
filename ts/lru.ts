
import { LinkedList, ListNode,ILogger ,PrintDirection} from "./list";

class LRUCacheItem<K,T>{
    private _key: K
    private _data: T
    constructor(key: K, value: T) {
        this._key = key;
        this._data = value;
    }
    get Key() {
        return this._key;
    }
    get Data() {
        return this._data;
    }
    set Data(data: T) {
        this._data = data;
    }
}
/**
 * provide a LRU Cache
 */
export class LRUCache<K,T> {
    private _list: LinkedList<LRUCacheItem<K,T>>
    private _capacity: number = 0;
    private _map: Map<K, ListNode<LRUCacheItem<K,T>>> = new Map();
    constructor(capacity = 10) {
        if (capacity < 1) {
            capacity = 10
        }
        this._list=new LinkedList();
        this._capacity = capacity;
    }
    private getCacheNode(key: K): ListNode<LRUCacheItem<K,T>> | null {
        if (!this._map.has(key)) {
            return null
        }
        let item = this._map.get(key)
        this._list.moveToBack(item ?? null)
        return item ?? null
    }
    // push a k-v
    push(key: K, value: T) {
        var item = this.getCacheNode(key);
        if (item != null) {
            item.Value.Data = value;
        } else {
            if (this._capacity == this._list.Count) {
                var first = this._list.front();
                this._list.remove(first)
                if(first?.Value.Key==undefined){
                    throw new Error("invalid list node:key");
                }
                this._map.delete(first?.Value.Key)
            }
            let newItem = new LRUCacheItem<K,T>(key, value);
            this._map.set(key, this._list.pushToBack(newItem))
        }
    }
    // get value
    get(key: K): T | null {
        return this.getCacheNode(key)?.Value?.Data ?? null
    }
    // print list
    print(logger:ILogger,direction:PrintDirection){
        this._list.print(logger,direction);
    }
}