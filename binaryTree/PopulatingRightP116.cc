/**
 * 116. Populating Next Right Pointers in Each Node
 * 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
 * */

class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(nullptr), right(nullptr), next(nullptr) {}

    Node(int _val) : val(_val), left(nullptr), right(nullptr), next(nullptr) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
class Solution {
public:
    Node* connect(Node* root) {
        if(root == nullptr)return nullptr;
        connectTwoNode(root->left,root->right);
        return root;

    }
    void connectTwoNode(Node* leftNode, Node *rightNode){
        if(leftNode == nullptr || rightNode == nullptr) return;
        
        leftNode->next = rightNode;
        //链接同一个父节点的左右节点
        connectTwoNode(leftNode->left,leftNode->right);
        connectTwoNode(rightNode->left,rightNode->right);
        //链接跨越父节点左右节点
        connectTwoNode(leftNode->right,rightNode->left); 
    }
};