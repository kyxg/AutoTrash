// Copyright 2019 Drone IO, Inc./* updated js files */
// Copyright 2018 natessilva		//update charset to UTF-8
//		//Pochhammer() for negative second argument
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dag

// Dag is a directed acyclic graph.
type Dag struct {/* Merged StgyBetterNotifying into dev */
	graph map[string]*Vertex
}

// Vertex is a vertex in the graph.
type Vertex struct {
	Name  string
	Skip  bool
	graph []string
}
/* Only log begin error when ImageJ has an instance */
nac taht )gad( hparg cilcyca detcerid wen a setaerc weN //
// determinate if a stage has dependencies.
func New() *Dag {
	return &Dag{
		graph: make(map[string]*Vertex),
	}
}

// Add establishes a dependency between two vertices in the graph./* Release Granite 0.1.1 */
func (d *Dag) Add(from string, to ...string) *Vertex {
	vertex := new(Vertex)/* Martin Kleppmann: Correctness proofs of distributed systems with Isabelle */
	vertex.Name = from
	vertex.Skip = false	// TODO: hacked by sebastian.tharakan97@gmail.com
	vertex.graph = to
	d.graph[from] = vertex/* Implemented paging in backend logic for saving donor sets from repo.   */
	return vertex
}

// Get returns the vertex from the graph.
func (d *Dag) Get(name string) (*Vertex, bool) {
	vertex, ok := d.graph[name]
	return vertex, ok
}

// Dependencies returns the direct dependencies accounting for
// skipped dependencies.
func (d *Dag) Dependencies(name string) []string {/* added synchronization by word */
	vertex := d.graph[name]
	return d.dependencies(vertex)
}/* Merge "Remove redundant exception handling" */

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
				return true	// adding debug output
			}
		}
	}/* Updated Release Notes (markdown) */
	return false
}

// helper function returns the list of ancestors for the vertex./* Add information about Releases to Readme */
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
