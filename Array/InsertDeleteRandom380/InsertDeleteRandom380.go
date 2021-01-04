/*
380. 常数时间插入、删除和获取随机元素
设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。

insert(val)：当元素 val 不存在时，向集合中插入该项。
remove(val)：元素 val 存在时，从集合中移除该项。
getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
*/
package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums      []int
	val2Index map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.val2Index[val]; ok {
		return false
	}
	this.val2Index[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.val2Index[val]; !ok {
		return false
	}
	index := this.val2Index[val]
	//将最后一个元素对应的下标 改为index
	this.val2Index[this.nums[len(this.nums)-1]] = index
	//交换val 和最后一个元素的位置
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	//删除最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	//删除字典中val以及下标
	delete(this.val2Index, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)
	param_2 := obj.Insert(1)
	fmt.Println(param_1)
	fmt.Println(param_2)
	param_3 := obj.Remove(1)
	fmt.Println(param_3)
}
