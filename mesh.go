package surface

import "github.com/vladimir-ch/dcel"

// Point is a point in 3D Euclidean space.
type Point [3]float64

// Vertex is a DCEL node with a position in 3D space.
type Vertex struct {
	dcel.BaseNode
	Point Point
}

// Mesh is a triangle mesh in 3D space.
type Mesh struct {
	dcel.Graph
}

// NewMesh returns a new Mesh.
func NewMesh() *Mesh {
	return &Mesh{
		Graph: *dcel.New(items{}),
	}
}

// AddVertex adds a new vertex located at p to the mesh.
func (m *Mesh) AddVertex(p Point) *Vertex {
	u := m.AddNode(m.NewNodeID()).(*Vertex)
	u.Point = p
	return u
}

// Vertex returns a vertex with the given id or nil if such vertex does not
// exist in the mesh.
func (m *Mesh) Vertex(id int) *Vertex {
	u := m.Node(id)
	if u == nil {
		return nil
	}
	return u.(*Vertex)
}

// AddFace adds a new triangle face with the given vertices to the mesh.
// If such face cannot be added, it returns a non-nil error.
func (m *Mesh) AddFace(n1, n2, n3 *Vertex) error {
	return m.Graph.AddFace(m.NewFaceID(), n1, n2, n3)
}

type items struct {
	dcel.Base
}

func (items) NewNode(id int) dcel.Node {
	return &Vertex{
		BaseNode: *dcel.NewBaseNode(id),
	}
}
