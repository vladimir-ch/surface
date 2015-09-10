package obj

import (
	"os"
	"testing"

	"github.com/vladimir-ch/surface"
)

type readerTest struct {
	name  string
	nodes int
	edges int
	faces int
}

var tests = []readerTest{
	{
		name:  "square",
		nodes: 4,
		edges: 5,
		faces: 2,
	},
	{
		name:  "cube",
		nodes: 8,
		edges: 18,
		faces: 12,
	},
	{
		name:  "cube_relative_indices",
		nodes: 24,
		edges: 30,
		faces: 12,
	},
}

func TestOBJReader(t *testing.T) {
	for _, test := range tests {
		filename := "testdata/" + test.name + ".obj"
		file, err := os.Open(filename)
		if err != nil {
			panic("error opening file" + filename)
		}
		defer file.Close()

		m := surface.NewMesh()
		err = NewReader(file).ReadGeometry(m)
		if err != nil {
			t.Error(err)
		}

		nodes := len(m.Nodes())
		if nodes != test.nodes {
			t.Errorf("%v: want %d nodes, got %d", test.name, test.nodes, nodes)
		}

		edges := len(m.Edges())
		if edges != test.edges {
			t.Errorf("%v: want %d edges, got %d", test.name, test.edges, edges)
		}

		faces := len(m.Faces())
		if faces != test.faces {
			t.Errorf("%v: want %d faces, got %d", test.name, test.faces, faces)
		}
	}
}
