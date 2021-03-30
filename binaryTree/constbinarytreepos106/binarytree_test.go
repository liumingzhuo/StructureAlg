package constbinarytreepos106

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	got := BuildTree(inorder, postorder)
	want := []int{3, 9, 20, 0, 0, 15, 7}
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("error")
	}

}
