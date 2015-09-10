package surface

import "testing"

func TestMesh(t *testing.T) {
	m := NewMesh()
	m.AddNode(Point{0, 0, 0})
	m.AddNode(Point{1, 2, 3})
	m.AddNode(Point{2, 3, 4})
	m.AddNode(Point{3, 4, 5})

	if err := m.AddFace(m.Node(0), m.Node(1), m.Node(2)); err != nil {
		t.Error(err)
	}

	n := m.Node(3)
	if n.Point[1] != 4 {
		t.Error("bad coordinate")
	}
}
