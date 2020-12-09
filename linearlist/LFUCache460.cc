/**
 * 460. LFU Cache 最近最少使用算法
Input
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Explanation
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);
lfu.put(2, 2);
lfu.get(1);      // return 1
lfu.put(3, 3);   // evicts key 2
lfu.get(2);      // return -1 (not found)
lfu.get(3);      // return 3
lfu.put(4, 4);   // evicts key 1.
lfu.get(1);      // return -1 (not found)
lfu.get(3);      // return 3
lfu.get(4);      // return 4

["LFUCache","put","put","put","put","get"]
[[2],[3,1],[2,1],[2,2],[4,4],[2]]
 */ 
#include <iostream>
#include <list>
#include <map>
using namespace std;
class LFUCache {
public:
    typedef list<pair<int,int>>:: iterator lpit; 
    map<int, lpit> key2val;
    map<int, int> key2freq;
    map<int, list<pair<int,int>> > freq2key;
    int minFreq;
    int cap;
    
    LFUCache(int capacity) {
        this->cap = capacity;
        this->minFreq = 0;
    }
    
    int get(int key) {
        if(key2val.find(key) == key2val.end())return -1;
        //增加key对应的freq
        increaseFreq(key);
        return key2val[key]->second;
    }
    
    void put(int key, int value) {
        if(cap <= 0 ) return;
        //如果key 已经存在  则需要freq+1
        if(key2val.find(key) != key2val.end()){
            //相同的key 更新值
            key2val[key]->second = value;
            increaseFreq(key);
            return;
        }
        //插入key
        //容量已满的话需要淘汰一个 freq 最小的 key 
        if(cap <= key2val.size()){
            removeMinFreqKey();
        }
        //插入kf表
        key2freq[key] = 1;
        //插入fk表   2 1   -》  4 4
        freq2key[1].push_front({key,value});
        //插入kv表
        key2val[key] = freq2key[1].begin();
        minFreq = 1;
    }
    void removeMinFreqKey(){
        lpit keylist = --freq2key[minFreq].end();
        int deleteKey = keylist->first;
        freq2key[minFreq].erase(keylist);
        //更新kv表
        key2val.erase(key2val.find(deleteKey));
        //更新kf表
        key2freq.erase(key2freq.find(deleteKey));

    }
    void increaseFreq(int key){
        int preFreq = key2freq[key];
        int freq = preFreq + 1;
        int res = key2val[key]->second;
        //将key的指针 从freq2key对应的列表中删除
        freq2key[preFreq].erase(key2val[key]);
        //更新kf表
        key2freq[key] = freq;
        //将key加入到preFreq+1中
        freq2key[freq].push_front({key,res});
        key2val[key] = freq2key[freq].begin();
        if(freq2key[preFreq].begin()==freq2key[preFreq].end() && minFreq==preFreq) minFreq+=1;
    }
};

/**
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache* obj = new LFUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */