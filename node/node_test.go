/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T09:45:36+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-02T11:10:00+00:00
 */
package node

import (
	"testing"
)

func TestNode(t *testing.T) {
	nA := Node{Name: "A", Item: "someitem"}
	nB := Node{Name: "B", Item: "SomethingElse"}
	if nA == nB {
		t.Fatalf("Different nodes are somehow equal %v == %v", nA, nB)
	}
	nA.Name = "B"
	nA.Item = "SomethingElse"
	if nA != nB {
		t.Fatalf("Modified and declared nodes not equal %v != %v", nA, nB)
	}
}
