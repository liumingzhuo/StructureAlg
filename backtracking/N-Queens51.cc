/**
 * 51. N-Queens
 * The n-queens puzzle is the problem of placing n queens on an n x n chessboard
 * such that no two queens attack each other.
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 * Each solution contains a distinct board configuration of the n-queens' placement, 
 * where 'Q' and '.' both indicate a queen and an empty space, respectively.
 **/
#include<vector>
#include<string>
#include<iostream>
using namespace std;
class Solvetion{
public:
    vector<vector<string> > res;
    vector<vector<string> > solveNQueens(int n){
        vector<string> board(n,string(n,'.'));
        backtrack(board,0);
        return res;
    }   

    void backtrack(vector<string> &board, int row){
        if (board.size() == row){
            res.push_back(board);
            return;
        }
        int cols = board[row].size();
        for(int col = 0; col < cols; col++){
            if (!isValid(board,row,col)) continue;
            board[row][col] = 'Q';
            backtrack(board,row+1);
            board[row][col]='.';
        }
    }

    bool isValid(vector<string> &board, int row, int col){
        int rows = board.size();
        //检查列是否有皇后相互冲突
        for(int i = 0;i< rows; i++){
            if (board[i][col]=='Q') return false;
        }  
        //检查右上方是否有皇后冲突
        for(int i = row-1,j = col+1; i>=0 && j<rows;i--,j++){
            if (board[i][j] == 'Q')return false;
        }
        //检查左上方是否有皇后冲突
        for (int i = row-1,j=col-1; i>=0 && j>=0;i--,j--){
            if (board[i][j] == 'Q') return false;
        }
        return true;
    }
};
