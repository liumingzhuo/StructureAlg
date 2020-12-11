/**
 * 130. 被围绕的区域
 * 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
X X X X
X O O X
X X O X
X O X X
->
X X X X
X X X X
X X X X
X O X X
 */ 
#include<vector>
using namespace std;

class Solution {
public:
    void solve(vector<vector<char>>& board) {
        if(board.size() == 0) return;
        int m = board.size(); // 行
        int n = board[0].size(); // 列
        for(int i = 0; i < m; i++){
            if(board[i][0] == 'O') dfs(board,i,0);
            if(board[i][n-1] == 'O') dfs(board,i,n-1);
        }
        for(int j = 0; j < n; j++){
            if(board[0][j] == 'O') dfs(board,0,j);
            if(board[m-1][j] == 'O') dfs(board,m-1,j);
        }
        for(int i = 0; i < m; i++){
            for(int j = 0; j < n; j++){
                if(board[i][j] == 'O') board[i][j] = 'X';
                if(board[i][j] == '#') board[i][j] = 'O';
            }
        }
    }
    void dfs(vector<vector<char>> &board, int i, int j){
        if(i < 0|| j < 0 || board.size() <= i || board[0].size() <= j
        || board[i][j] == 'X' || board[i][j] == '#') return;
        board[i][j] = '#';
        dfs(board, i - 1, j); //上
        dfs(board, i + 1, j);
        dfs(board, i, j - 1);
        dfs(board, i, j + 1);//右
    }
};