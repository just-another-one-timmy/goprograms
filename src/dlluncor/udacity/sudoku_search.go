package udacity

import(
  "dlluncor/container"
  "container/heap"
)

/* This file deals with doing an A* search using utilities found in
 * sudoku.go.
 */

type SNode struct {
  state *CellState
  cost int32
  f int32 // how many hops to get to this state.
  h int32 // how far am I from the goal.
}


/*
func (s *SFrontier) () {
  
}
*/

type SFrontier struct {
  queue *container.PriorityQueue
}

func (s *SFrontier) IsEmpty() bool {
  return true
}

func (s *SFrontier) RemoveChoice() interface{} {
  node := heap.Pop(s.queue).(*SNode)
  return node
}

func (s *SFrontier) Contains(node interface{}) bool {
  return true
}

func (s *SFrontier) Add(inode interface{}) {
  node := inode.(*SNode)
  item := &container.Item{
    Value: node,
    Priority: node.f + node.h,
  }
  heap.Push(s.queue, item)
  // TODO(dlluncor): Add to list of explored states.
}


type SExplored struct {
  seen map[string]bool
}

/*
func (s *SExplored) () {
  
}
*/

func (s *SExplored) Add(node interface{}) {
  
}

func (s *SExplored) Contains(node interface{}) bool {
  return false
}

type SudokuSolver struct {
  frontier *SFrontier
  explored *SExplored
}

/*
func (s *SudokuSolver) () {
  
}
*/

func (s *SudokuSolver) IsGoal(node interface{}) bool {
  return false 
}

func (s *SudokuSolver) NextActions(node interface{}) []interface{} {
  arr := make([]interface{}, 0)
  return arr
}

func (s *SudokuSolver) Init(s0 *CellState) {
  s.frontier = &SFrontier{
    queue: &container.PriorityQueue{},
  }
  heap.Init(s.frontier.queue)
  s.explored = &SExplored{
    seen: make(map[string]bool),
  }
  node0 := &SNode{
    f: 0,
    h: 0,
    state: s0,
  }
  s.frontier.Add(node0)
}