package algorithms

type TreeNode struct {
  value int;
  left *TreeNode;
  right *TreeNode;
}

func PreOrder(node TreeNode, arr *[]int) {
  *arr = append(*arr, node.value);

  if (node.left != nil) {
    PreOrder(*node.left, arr);
  }

  if (node.right != nil) {
    PreOrder(*node.right, arr);
  }
}

func InOrder(node TreeNode, arr *[]int) {
  if (node.left != nil) {
    InOrder(*node.left, arr);
  }

  *arr = append(*arr, node.value);

  if (node.right != nil) {
    InOrder(*node.right, arr);
  }
}

func search(traversal []int, value int) int {
  for i := 0; i < len(traversal); i += 1 {
    if (traversal[i] == value) {
      return i;
    }
  }
  panic("not found");
}

func BuildTree(preOrder []int, inOrder []int, inStart int, inEnd int, preOrderIndex *int) *TreeNode {
  if (inStart > inEnd) {
    return nil;
  }

  node := TreeNode{
    value: preOrder[*preOrderIndex],
  }
  *preOrderIndex += 1;

  if (inStart == inEnd) {
    return &node;
  }

  inIndex := search(inOrder, node.value);

  node.left = BuildTree(preOrder, inOrder, inStart, inIndex - 1, preOrderIndex);
  node.right = BuildTree(preOrder, inOrder, inIndex + 1, inEnd, preOrderIndex);
  
  return &node;
}
