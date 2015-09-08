package surface

import "testing"

func TestMesh(t *testing.T) {
	m := NewMesh()
	m.AddVertex(0, Point{0, 0, 0})
	m.AddVertex(1, Point{1, 2, 3})
	m.AddVertex(2, Point{2, 3, 4})
	m.AddVertex(3, Point{3, 4, 5})

	if err := m.AddFace(0, m.Vertex(0), m.Vertex(1), m.Vertex(2)); err != nil {
		t.Error(err)
	}

	n := m.Vertex(3)
	if n.Point[1] != 4 {
		t.Error("bad coordinate")
	}
}
