// Copyright 2015 Vladimír Chalupecký. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/vladimir-ch/surface"
)

type ParseError struct {
	Line int
	Err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.Err)
}

var (
	ErrCoordinateCount = errors.New("wrong number of vertex coordinates")
	ErrFaceNodeCount   = errors.New("wrong number of face nodes")
	ErrInvalidNodeID   = errors.New("invalid node ID")
)

type Reader struct {
	s    *bufio.Scanner
	line int
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		s: bufio.NewScanner(r),
	}
}

func (r *Reader) error(err error) error {
	return &ParseError{
		Line: r.line,
		Err:  err,
	}
}

func (r *Reader) ReadGeometry(m *surface.Mesh) (err error) {
	var vertices []*surface.Node // Mapping from file node IDs to mesh nodes.
	for r.s.Scan() {
		line := bytes.TrimSpace(r.s.Bytes())
		r.line++
		if len(line) == 0 || line[0] == '#' {
			// Skip empty lines and comments.
			continue
		}
		fields := bytes.Fields(line)
		switch string(fields[0]) {
		case "v":
			p, err := r.parseVertex(fields[1:])
			if err != nil {
				return r.error(err)
			}
			u := m.AddNode(p)
			vertices = append(vertices, u)
		case "f":
			nodeIDs, err := r.parseFace(fields[1:])
			if err != nil {
				return r.error(err)
			}
			if len(nodeIDs) > 3 {
				// Skip non-triangle faces.
				continue
			}
			for i, id := range nodeIDs {
				switch {
				case id < 0:
					// Convert negative relative indices to positive absolute ones.
					nodeIDs[i] = len(vertices) + id
				case id > 0:
					// Convert 1-based indices to 0-based.
					nodeIDs[i]--
				}
			}
			for _, id := range nodeIDs {
				// Node indices are 1-based, so zero index is invalid.
				if id < 0 || id >= len(vertices) {
					return r.error(ErrInvalidNodeID)
				}
			}
			v1 := vertices[nodeIDs[0]]
			v2 := vertices[nodeIDs[1]]
			v3 := vertices[nodeIDs[2]]
			err = m.AddFace(v1, v2, v3)
			if err != nil {
				return r.error(err)
			}
		}
	}
	if err := r.s.Err(); err != nil {
		return r.error(err)
	}
	return nil
}

func (r *Reader) parseVertex(fields [][]byte) (p surface.Point, err error) {
	if len(fields) < 3 {
		return p, ErrCoordinateCount
	}
	for i := 0; i < 3; i++ {
		p[i], err = strconv.ParseFloat(string(fields[i]), 64)
		if err != nil {
			return p, err
		}
	}
	return p, err
}

func (r *Reader) parseFace(fields [][]byte) ([]int, error) {
	if len(fields) < 3 {
		return nil, ErrFaceNodeCount
	}
	var (
		refs  [][]byte
		nodes []int
	)
	for _, field := range fields {
		refs = bytes.Split(field, []byte{'/'})             // Split each field on '/'.
		id, err := strconv.ParseInt(string(refs[0]), 0, 0) // Convert the node ID reference to int.
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, int(id))
	}
	return nodes, nil
}
