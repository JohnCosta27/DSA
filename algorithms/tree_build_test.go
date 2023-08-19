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

  first := getNode(97);
  second := getNode(1);
  third := getNode(98);
  forth := getNode(0);
  fifth := getNode(100);
  sixth := getNode(99);
  root := getNode(2);

  root.left = &first;
  root.right = &second;

  second.left = &third;
  second.right = &forth;

  forth.left = &fifth;
  forth.right = &sixth;


  arr := []int{};
  expectedArr := []int{2, 97, 1, 98, 0, 100, 99};
  PreOrder(root, &arr);

  for i, v := range arr {
    if (expectedArr[i] != v) {
      t.Logf("Pre Order | Expected %d, found %d\n", expectedArr[i], v);
      t.Fail();
    }
  }

  arr2 := []int{};
  expectedArr2 := []int{97, 2, 98, 1, 100, 0, 99};
  InOrder(root, &arr2);

  t.Log(arr2);

  for i, v := range arr2 {
    if (expectedArr2[i] != v) {
      t.Logf("In Order | Expected %d, found %d\n", expectedArr[i], v);
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

func TestDuplicate(t *testing.T) {
  first := getNode(97);
  second := getNode(0);
  third := getNode(98);
  forth := getNode(0);
  fifth := getNode(100);
  sixth := getNode(99);
  root := getNode(0);

  root.left = &first;
  root.right = &second;

  second.left = &third;
  second.right = &forth;

  forth.left = &fifth;
  forth.right = &sixth;

  arr := []int{};
  expectedArr := []int{0, 97, 0, 98, 0, 100, 99};
  PreOrder(root, &arr);

  for i, v := range arr {
    if (expectedArr[i] != v) {
      t.Logf("Pre Order | Expected %d, found %d\n", expectedArr[i], v);
      t.Fail();
    }
  }

  counter := 0;
  node := BuildSimpleTree(arr, &counter);

  t.Log(node.value);

  preOrderTree := []int{};
  PreOrder(*node, &preOrderTree);

  for i, v := range preOrderTree {
    if (expectedArr[i] != v) {
      t.Logf("Expected %d, found %d\n", v, expectedArr[i]);
      t.Fail();
    }
  }
}
