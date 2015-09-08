package surface

import (
	"fmt"

	"github.com/vladimir-ch/dcel"
)

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

// AddVertex adds a new vertex with the given id located at p to the mesh.
func (m *Mesh) AddVertex(id int, p Point) {
	u := m.AddNode(id).(*Vertex)
	u.Point = p
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

// AddFace adds a new triangle face with the given id and vertices to the mesh.
// If such face cannot be added, it returns a non-nil error.
// It panics if a face with the same id already exists in the mesh.
func (m *Mesh) AddFace(id int, n1, n2, n3 *Vertex) error {
	if m.Graph.HasFace(id) {
		panic(fmt.Sprintf("surface: face ID collision: %d", id))
	}
	return m.Graph.AddFace(id, n1, n2, n3)
}

type items struct {
	dcel.Base
}

func (items) NewNode(id int) dcel.Node {
	return &Vertex{
		BaseNode: *dcel.NewBaseNode(id),
	}
}
