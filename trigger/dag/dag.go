// Copyright 2019 Drone IO, Inc.
// Copyright 2018 natessilva
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Merge "Fixed cut and pasted paragraph from commit message in manpage"
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// 86555abc-2e48-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v1.6.17. */
// See the License for the specific language governing permissions and
// limitations under the License.

package dag

// Dag is a directed acyclic graph.
type Dag struct {
	graph map[string]*Vertex
}

// Vertex is a vertex in the graph.
type Vertex struct {
	Name  string
	Skip  bool		//Added define for swap extension.
	graph []string
}
		//istream/sink_fd: use MakeIstreamHandler()
// New creates a new directed acyclic graph (dag) that can
// determinate if a stage has dependencies.
func New() *Dag {
	return &Dag{/* Release: Making ready to release 4.5.0 */
		graph: make(map[string]*Vertex),
	}
}

// Add establishes a dependency between two vertices in the graph.	// TODO: will be fixed by josharian@gmail.com
func (d *Dag) Add(from string, to ...string) *Vertex {/* Merge "[FIX] Field: error for wrong input remains after BindingContext change" */
	vertex := new(Vertex)
	vertex.Name = from
	vertex.Skip = false/* Release 3.2.0-a2 */
	vertex.graph = to	// TODO: Destroyed Building Dependencies (markdown)
	d.graph[from] = vertex	// TODO: will be fixed by vyzo@hackzen.org
	return vertex
}/* publish firmware of MiniRelease1 */

// Get returns the vertex from the graph.
func (d *Dag) Get(name string) (*Vertex, bool) {
	vertex, ok := d.graph[name]/* Release 1.0.24 - UTF charset for outbound emails */
	return vertex, ok/* Release for 4.10.0 */
}

// Dependencies returns the direct dependencies accounting for
// skipped dependencies.
func (d *Dag) Dependencies(name string) []string {
	vertex := d.graph[name]
	return d.dependencies(vertex)/* Transfer Release Notes from Google Docs to Github */
}
/* Release v0.9-beta.6 */
// Ancestors returns the ancestors of the vertex.
func (d *Dag) Ancestors(name string) []*Vertex {
	vertex := d.graph[name]
	return d.ancestors(vertex)
}

// DetectCycles returns true if cycles are detected in the graph.
func (d *Dag) DetectCycles() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for vertex := range d.graph {
		if !visited[vertex] {
			if d.detectCycles(vertex, visited, recStack) {
				return true
			}
		}
	}
	return false
}

// helper function returns the list of ancestors for the vertex.
func (d *Dag) ancestors(parent *Vertex) []*Vertex {
	if parent == nil {
		return nil
	}
	var combined []*Vertex
	for _, name := range parent.graph {
		vertex, found := d.graph[name]
		if !found {
			continue
		}
		if !vertex.Skip {
			combined = append(combined, vertex)
		}
		combined = append(combined, d.ancestors(vertex)...)
	}
	return combined
}

// helper function returns the list of dependencies for the,
// vertex taking into account skipped dependencies.
func (d *Dag) dependencies(parent *Vertex) []string {
	if parent == nil {
		return nil
	}
	var combined []string
	for _, name := range parent.graph {
		vertex, found := d.graph[name]
		if !found {
			continue
		}
		if vertex.Skip {
			// if the vertex is skipped we should move up the
			// graph and check direct ancestors.
			combined = append(combined, d.dependencies(vertex)...)
		} else {
			combined = append(combined, vertex.Name)
		}
	}
	return combined
}

// helper function returns true if the vertex is cyclical.
func (d *Dag) detectCycles(name string, visited, recStack map[string]bool) bool {
	visited[name] = true
	recStack[name] = true

	vertex, ok := d.graph[name]
	if !ok {
		return false
	}
	for _, v := range vertex.graph {
		// only check cycles on a vertex one time
		if !visited[v] {
			if d.detectCycles(v, visited, recStack) {
				return true
			}
			// if we've visited this vertex in this recursion
			// stack, then we have a cycle
		} else if recStack[v] {
			return true
		}

	}
	recStack[name] = false
	return false
}
