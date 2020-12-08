/**
 * 146. LRU 缓存机制
 * 
 */
#include<unordered_map>
using namespace std;
typedef struct Node{
    int key,val;
    Node *next, *prev;
    Node(int k,int v):key(k), val(v) ,next(nullptr),prev(nullptr) {};
}Node;

class DoubleList{
    Node *head, *tail;
    //链表元素数
    int size;
public:
    void initDoubleList(){
        head = new Node(0,0);
        tail = new Node(0,0);
        head->next = tail;
        tail->prev = head;
        size = 0;
    }
    //在链表尾部添加节点
    void addLast(Node *x){
        x->prev = tail->prev;
        x->next = tail;
        tail->prev->next = x;
        tail->prev = x;
        size++;
    }  
    //删除链表中的x节点
    void remove(Node *x){
        x->prev->next = x->next;
        x->next->prev = x->prev;
        size--;
    }
    //删除链表中到第一个节点
    Node* removeFirst(){
        if(head->next == tail) return nullptr;
        Node *first = head->next;
        remove(first);
        return first;
    }

    int getSize(){
        return size;
    }
};
class LRUCache {
public:
    //key -> Node(key,val)
    unordered_map<int, Node*> map;
    //lined  最后一位是最近使用  第一位是使用最久
    DoubleList cache;
    int cap ;
    LRUCache(int capacity) {
        cache =  DoubleList();
        cache.initDoubleList();
        cap = capacity;
    }
    //将某个key提升为最近使用
    void useRecently(int k){
        auto iterator = map.find(k);
        cache.remove(iterator->second);
        cache.addLast(iterator->second);
    }
    //添加到最近使用
    void addRecently(int key, int val){
        Node *x = new Node(key,val);
        cache.addLast(x);
        map[key] = x;
    }
    //删除某一个key
    void deletekey(int key){
        auto x = map.find(key);
        cache.remove(x->second);
        map.erase(x);
    }
    //删除较长时间未使用
    void removeLeastRecently(){
        Node *firstNode = cache.removeFirst();
        int deleteKey = firstNode->key;
        map.erase(deleteKey);
    }

    
    int get(int key) {
        if(map.find(key) == map.end()) return -1;
        useRecently(key);
        return map[key]->val;
    }
    
    void put(int key, int value) {
        if(map.find(key) != map.end()){
            deletekey(key);
            addRecently(key,value);
            return;
        }
        if(cap == cache.getSize()) removeLeastRecently();
        addRecently(key, value);
    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */