/**
 * 355. 设计推特
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户

["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]
[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]

["Twitter","postTweet","unfollow","getNewsFeed"]
[[],[1,5],[1,1],[1]]

["Twitter","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","postTweet","getNewsFeed"]
[[],[1,5],[1,3],[1,101],[1,13],[1,10],[1,2],[1,94],[1,505],[1,333],[1,22],[1,11],[1]]
[null,null,null,null,null,null,null,null,null,null,null,null,[11,22,333,505,94,2,10,13,101,3,5]]
 */
#include<vector>
#include<unordered_set>
#include<unordered_map>
#include<iostream>
#include<queue>
using namespace std;
typedef struct Tweet
{
    int id;
    int time;
    Tweet* next;  
    Tweet(int id,int time):id(id),time(time),next(nullptr){};  
}Tweet;

class Twitter {
public:
    int timestamp = 0;
    class User{
    private:
        int id;
    public:
        unordered_set<int> followed;
        Tweet* head;

        User(int userId){
            id = userId;
            head = nullptr;
            follow(userId);
        }
        void follow(int userId){
            followed.insert(userId);
        }
        void unfollow(int userId){
            if(userId != this->id)
                followed.erase(userId);
        }
        void post(int twitterId,int timestamp){
            Tweet* t =new Tweet(twitterId,timestamp);
            t->next = head;
            head = t;
        }
    };
    unordered_map<int, User*> userMap;
    /** Initialize your data structure here. */
    Twitter() {
    }
    
    /** Compose a new tweet. */
    void postTweet(int userId, int tweetId) {
        if(userMap.find(userId) == userMap.end()){
            User* user = new User(userId);
            userMap[userId] = user;
        }
        User* user = userMap.find(userId)->second;
        user->post(tweetId,timestamp);
        timestamp++;
    }
    
    /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
    vector<int> getNewsFeed(int userId) {
        vector<int> res;
        if(userMap.find(userId) == userMap.end()) return res;
        unordered_set<int> users = userMap[userId]->followed;
        auto cmp =[](Tweet* x,Tweet* y){return x->time > y->time;};
        priority_queue<Tweet*, vector<Tweet*>, decltype(cmp)> pq(cmp);
        for(int id:users){
            User* user = userMap[id];
            if(!user || !user->head)continue;
            pq.push(user->head);
        }
        while(!pq.empty()){
            if(res.size() == 10) break;
            Tweet* twwi = pq.top();
            res.push_back(twwi->id);
            pq.pop();
            if(twwi->next) pq.push(twwi->next);
        }
        return res;
    }
    
    /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
    void follow(int followerId, int followeeId) {
        if(userMap.find(followerId) == userMap.end()){
            User* user = new User(followerId);
            userMap[followerId] = user;
        }
        if(userMap.find(followeeId) == userMap.end()){
            User* user = new User(followeeId);
            userMap[followeeId] = user;
        }
        userMap[followerId]->follow(followeeId);
    }
    
    /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
    void unfollow(int followerId, int followeeId) {
        if(userMap.find(followerId) != userMap.end()){
            userMap[followerId]->unfollow(followeeId);
        }
    }
};

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter* obj = new Twitter();
 * obj->postTweet(userId,tweetId);
 * vector<int> param_2 = obj->getNewsFeed(userId);
 * obj->follow(followerId,followeeId);
 * obj->unfollow(followerId,followeeId);
 */
