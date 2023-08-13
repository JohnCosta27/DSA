package algorithms

import (
	"testing"
)

func getNode(value int) TreeNode {
  return TreeNode{
    value: value,
    left: nil,
    right: nil,
  }
}

func TestBasic(t *testing.T) {

  first := getNode(1);
  second := getNode(2);
  third := getNode(3);
  forth := getNode(4);
  fifth := getNode(5);
  sixth := getNode(6);
  root := getNode(7);

  third.left = &first;
  third.right = &second;

  sixth.left = &forth;
  sixth.right = &fifth;

  root.left = &third;
  root.right = &sixth;

  arr := []int{};
  expectedArr := []int{7, 3, 1, 2, 6, 4, 5};
  PreOrder(root, &arr);

  for i, v := range arr {
    if (expectedArr[i] != v) {
      t.Logf("Expected %d, found %d\n", v, expectedArr[i]);
      t.Fail();
    }
  }

  arr2 := []int{};
  expectedArr2 := []int{1, 3, 2, 7, 4, 6, 5};
  InOrder(root, &arr2);

  for i, v := range arr2 {
    if (expectedArr2[i] != v) {
      t.Logf("Expected %d, found %d\n", v, expectedArr[i]);
      t.Fail();
    }
  }

  counter := 0;
  node := BuildTree(arr, arr2, 0, len(arr2) - 1, &counter);

  rebuildPre := []int{};
  PreOrder(*node, &rebuildPre);

  rebuildIn := []int{};
  InOrder(*node, &rebuildIn);

  for i, v := range rebuildPre {
    if (expectedArr[i] != v) {
      t.Logf("Expected %d, found %d\n", v, expectedArr[i]);
      t.Fail();
    }
  }

  for i, v := range rebuildIn {
    if (expectedArr2[i] != v) {
      t.Logf("Expected %d, found %d\n", v, expectedArr2[i]);
      t.Fail();
    }
  }
}
