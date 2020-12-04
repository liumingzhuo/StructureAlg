/**
 * 297. 二叉树的序列化与反序列化
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
 */
#include<iostream>
#include<string>
#include<queue>
#include<sstream>
using namespace std;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string str;
        serialize(root,str);
        return str;
    }

    void serialize(TreeNode* root,string& str){
        if(!root){
            str.append("#").append(",");
            return;
        }
        str.append(to_string(root->val)).append(",");
        serialize(root->left,str);
        serialize(root->right,str);
    } 

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        string item;
        queue<string> q;
        stringstream ss(data);
        while(getline(ss,item,','))
            q.push(item);
        return deserialize(q);
    }
    TreeNode* deserialize(queue<string>& q){
        string first = q.front();
        q.pop();
        if(first == "#") return nullptr;
        TreeNode* root = new TreeNode(stoi(first));
        root->left = deserialize(q);
        root->right= deserialize(q);
        return root;
    }
};
