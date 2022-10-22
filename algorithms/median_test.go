package algorithms

import "testing"

func TestCrudeRank(t *testing.T) {
  list1 := []int{50, 12, 45, 10, 65}
  rank := CrudeRank(list1, 50)
  if rank != 3 {
    t.Logf("Rank should have been 3, instead: %d\n", rank)
    t.FailNow()
  }
}

func TestCrudeSelect(t *testing.T) {
  list1 := []int{50, 12, 45, 10, 65}
  selected := CrudeSelect(list1, 2)
  if selected != 45 {
    t.Logf("Select should have been 45, instead: %d\n", selected)
    t.FailNow()
  }
}
