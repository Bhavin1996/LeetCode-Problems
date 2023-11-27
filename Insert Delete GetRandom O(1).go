package leetcode

type RandomizedSet struct {
	hashMap  map[int]int
	keyValue []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{hashMap: make(map[int]int), keyValue: make([]int, 0, 500)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		return false
	}
	this.hashMap[val] = len(this.keyValue)
	this.keyValue = append(this.keyValue, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		delete(this.hashMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return 1
}
