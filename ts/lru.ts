
import { LinkedList, ListNode } from "./list";

class LRUCacheItem<T>{
    private _key: string
    private _data: T
    constructor(key: string, value: T) {
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

export class LRUCache<T> {
    private _list: LinkedList<LRUCacheItem<T>>
    private _capability: number = 0;
    private _map: Map<string, ListNode<LRUCacheItem<T>>> = new Map();
    constructor(capability = 10) {
        if (capability < 1) {
            capability = 10
        }
        this._list=new LinkedList();
        this._capability = capability;
    }
    private getCacheNode(key: string): ListNode<LRUCacheItem<T>> | null {
        if (!this._map.has(key)) {
            return null
        }
        let item = this._map.get(key)
        this._list.moveToBack(item ?? null)
        return item ?? null
    }
    push(key: string, value: T) {
        var item = this.getCacheNode(key);
        if (item != null) {
            item.Value.Data = value;
        } else {
            if (this._capability == this._list.Count) {
                var first = this._list.front();
                this._list.remove(first)
                this._map.delete(first?.Value.Key ?? "")
            }
            let newItem = new LRUCacheItem<T>(key, value);
            this._map.set(key, this._list.pushToBack(newItem))
        }
    }
    get(key: string): T | null {
        return this.getCacheNode(key)?.Value?.Data ?? null
    }
    print(){
        this._list.print();
    }
}