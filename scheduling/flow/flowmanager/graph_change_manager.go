// Copyright 2016 The ksched Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flowmanager

import (
	"github.com/coreos/ksched/scheduling/flow/dimacs"
	"github.com/coreos/ksched/scheduling/flow/flowgraph"
)

// The GraphChangeManager bridges GraphManager and Graph. Every
// graph change done by the GraphManager should be conducted via
// FlowGraphChangeManager's methods.
// The class stores all the changes conducted in-between two scheduling rounds.
// Moreover, FlowGraphChangeManager applies various algorithms to reduce
// the number of changes (e.g., merges idempotent changes, removes superfluous
// changes).
type GraphChangeManager interface {
	AddArc(src, dst *flowgraph.Node,
		capLowerBound, capUpperBound uint64,
		cost int64,
		arcType flowgraph.ArcType,
		changeType dimacs.ChangeType,
		comment string) *flowgraph.Arc

	AddNode(nodeType flowgraph.NodeType,
		excess int64,
		changeType dimacs.ChangeType,
		comment string) *flowgraph.Node

	ChangeArc(arc flowgraph.Arc, capLowerBound uint64,
		capUpperBound uint64, cost int64,
		changeType dimacs.ChangeType, comment string)

	ChangeArcCapacity(arc flowgraph.Arc, capacity uint64,
		changeType dimacs.ChangeType, comment string)

	ChangeArcCost(arc flowgraph.Arc, cost int64,
		changeType dimacs.ChangeType, comment string)

	DeleteArc(arc flowgraph.Arc, changeType dimacs.ChangeType, comment string)

	DeleteNode(arc flowgraph.Node, changeType dimacs.ChangeType, comment string)

	GetGraphChanges() []*dimacs.Change

	GetOptimizedGraphChanges() []*dimacs.Change

	ResetChanges()

	CheckNodeType(nodeID uint64, typ flowgraph.NodeType) bool

	// FlowGraph getter: Returns flow graph instance for this manager
	Graph() *flowgraph.Graph
	// Node getter
	Node(nodeID uint64) *flowgraph.Node
}

// The change manager that should implement the ChangeMangerInterface
type changeManager struct {
	flowGraph *flowgraph.Graph
	// Vector storing the graph changes occured since the last scheduling round.
	graphChanges []*dimacs.Change
	dimacsStats  *dimacs.ChangeStats
}

// Public Interface functions
func (cm *changeManager) AddArc(src, dst *flowgraph.Node,
	capLowerBound, capUpperBound uint64,
	cost int64,
	arcType flowgraph.ArcType,
	changeType dimacs.ChangeType,
	comment string) *flowgraph.Arc {

	arc := cm.flowGraph.AddArc(src, dst)
	arc.CapLowerBound = capLowerBound
	arc.CapUpperBound = capUpperBound
	arc.Cost = cost
	arc.Type = arcType

	// TODO: add dimacs increamental change

	return nil
}

func (cm *changeManager) AddNode(t flowgraph.NodeType, excess int64, changet dimacs.ChangeType, comment string) *flowgraph.Node {
	n := cm.flowGraph.AddNode()
	n.Type = t
	n.Excess = excess
	n.Comment = comment

	// TODO: add dimacs increamental change

	return n
}

func (cm *changeManager) DeleteNode(n *flowgraph.Node, changeType dimacs.ChangeType, comment string) {
	cm.flowGraph.DeleteNode(n)

	// TODO: add dimacs increamental change
}

func (cm *changeManager) ChangeArc(arc flowgraph.Arc, capLowerBound uint64,
	capUpperBound uint64, cost int64,
	changeType dimacs.ChangeType, comment string) {

}

func (cm *changeManager) ChangeArcCapacity(arc flowgraph.Arc, capacity uint64,
	changeType dimacs.ChangeType, comment string) {

}

func (cm *changeManager) ChangeArcCost(arc flowgraph.Arc, cost int64,
	changeType dimacs.ChangeType, comment string) {

}

func (cm *changeManager) DeleteArc(arc flowgraph.Arc, changeType dimacs.ChangeType, comment string) {

}

func (cm *changeManager) GetGraphChanges() []*dimacs.Change {
	return nil
}

func (cm *changeManager) GetOptimizedGraphChanges() []*dimacs.Change {
	return nil
}

func (cm *changeManager) ResetChanges() {

}

func (cm *changeManager) CheckNodeType(nodeID uint64, typ flowgraph.NodeType) bool {
	return false
}

func (cm *changeManager) Graph() *flowgraph.Graph {
	return nil
}

func (cm *changeManager) Node(nodeID uint64) *flowgraph.Node {
	return nil
}

// Private helper methods for change_manager internal use
func (cm *changeManager) addGraphChange(change *dimacs.Change) {

}

func (cm *changeManager) optimizeChanges() {

}

func (cm *changeManager) mergeChangesToSameArc() {

}

// Checks if there's already a change for the (src_id, dst_id) arc.
// If there's no change then it adds one to the state, otherwise
// it updates the existing change.
func (cm *changeManager) mergeChangesToSameArcHelper(
	srcID, dstID, capLowerBound, capUpperBound uint64,
	cost int64, typ flowgraph.ArcType,
	change *dimacs.Change, newGraphChanges []*dimacs.Change,
	arcsSrcChanges map[uint64]map[uint64]*dimacs.Change,
	arcsDstChanges map[uint64]map[uint64]*dimacs.Change) {

}

func (cm *changeManager) purgeChangesBeforeNodeRemoval() {

}

func (cm *changeManager) removeDuplicateChanges() {

}

// Checks if there's already an identical change for the (src_id, dst_id) arc.
// If there's no change the it updates the state, otherwise it just ignores
// the change we're currently processing because it's duplicate.
func (cm *changeManager) removeDuplicateChangesHelper(
	srcID, dstID uint64, change *dimacs.Change,
	newGraphChanges []*dimacs.Change,
	node_to_change map[uint64]map[string]*dimacs.Change) {

}

func (cm *changeManager) removeDuplicateChangesUpdateState(
	nodeID uint64, change *dimacs.Change,
	node_to_change map[uint64]map[string]*dimacs.Change) bool {
	return false
}

// Method to be called upon node addition. This method makes sure that the
// state is cleaned when we re-use a node id.
func (cm *changeManager) removeDuplicateCleanState(
	newNodeID, srcID, dstID uint64,
	changeDesc string,
	node_to_change map[uint64]map[string]*dimacs.Change) {

}