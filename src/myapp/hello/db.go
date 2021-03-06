package hello

import (
    "time"
)

// Database model for the entire Game state.

// This map will now contain everything b/c I don't know how to save
// or retrieve any other fields from this stupid struct besides the
// map values WTF!!!
// TODO(dlluncor): Figure out how to get a map type in here. For now,
// not worth it!!
type MyGame struct {
   States []string
   Tables []string  // Only 4.
   // Running total for the game. Token[0] corresponds to Points[0] for a user,
   // as does Users[0]
   // If a user leaves, then oh well no one can get those points.
   Users  []string
   Tokens []string
   Points []int
   // Current Table Info.
   CurTable string
   CurWords []string // List of words found this round.
   CurRound int // Starts off at 1, ends at 4. (not used kept by client!)

   LastRoundFetched int64  // When was the last round info fetched.
   Now int64  // When am I sending back the entire game state.

   // Immutable information about a table.
   Language string // e.g., 'en', 'es'
}

func defaultGame() *MyGame {
    g := &MyGame{
        Users:  []string{},
        Tokens: []string{},
        Points: []int{},
        States: []string{},
        Tables: []string{},
    }
    g.AddState("notStarted")
    return g
}

// Delete the session state of the game but keep the users and tokens.
func (g *MyGame) Clear() {
  g.Points = []int{}
  g.Tables = []string{}
  g.States = []string{}
  g.CurTable = ""
  g.CurWords = []string{}
  g.CurRound = -1
  for _, _ = range g.Users {
    g.Points = append(g.Points, 0)
  }
}

// TODO(dlluncor): Pretty inefficient, but oh well! Go is fast I think...
func inArr(items []string, item string) bool {
  has := false
  for _, aItem := range items {
    if aItem == item {
        return true
    }
  }
  return has
}

func (g *MyGame) AddState(state string) {
  g.States = append(g.States, state)
}

func (g *MyGame) HadState(state string) bool {
  return inArr(g.States, state)
}

// Returns whether we've seen the user or not already.
func (g *MyGame) AddUserToken(user, token string) bool {
  // If this was called on reload, the user already exists here but has
  // a different token. So let's just update their token and keep everything
  // the same. This is only if the user is already in our pool.
  index := indexOf(g.Users, user)
  if index != -1 {
    g.Tokens[index] = token
    return true
  } else {
    g.Tokens = append(g.Tokens, token)
    g.Points = append(g.Points, 0)
    g.Users = append(g.Users, user)
    return false
  }
}

func removeEl(els []string, deleteIndex int) []string {
  newEls := []string{}
  for index, el := range els {
    if index == deleteIndex {
        continue
    }
    newEls = append(newEls, el)
  }
  return newEls
}

func removeElInt(els []int , deleteIndex int) []int {
  newEls := []int{}
  for index, el := range els {
    if index == deleteIndex {
        continue
    }
    newEls = append(newEls, el)
  }
  return newEls
}

func (g *MyGame) RemoveUser(user string) {
  index := indexOf(g.Users, user)
  g.Users = removeEl(g.Users, index)
  g.Tokens = removeEl(g.Tokens, index)
  g.Points = removeElInt(g.Points, index)
}

// SetTable("table1", "dsdfsdfs\nXXXX\nDDFD")
// We'll use the index for now to generate which table to serve up.
func (g *MyGame) SetTable(tableRoundKey, tableVal string) {
  g.Tables = append(g.Tables, tableVal)
}

func (g *MyGame) GetUserTokens() []string {
  return g.Tokens
}

type TableInfo struct {
  Table string
}

func (g *MyGame) CreateTableInfo(round int) {
  // Keep track of a new slate of data for this particular round.
  g.CurTable = g.Tables[round-1]
  g.CurRound = round
  g.CurWords = []string{}
}

// Get the length of our table encoded as 'abcdX\ndfdsX'
func getTableLen(table string) int {
  i := 0
  for _, char := range table {
    if char == '*' {
        break
    }
    i += 1
  }
  return i
}

// Gets the table information for this round.
func (g *MyGame) GetTableInfo() *TableInfo {
  // Run the algorithm to generate a list of solutions!!!
  //length := getTableLen(g.CurTable)
  // TODO(dlluncor): Can't get solutions in this thread for some reason
  // so going back to the old GET request way of doing things, that seems
  // to work...
  //answers := "cheese,potatoes"
  //tableToSolve := strings.Replace(g.CurTable, "*", "", -1)
  //solveForWords(tableToSolve, length)
  info := &TableInfo{
    Table: g.CurTable,
  }
  return info
}

func (g *MyGame) HasWord(word string) bool {
  return inArr(g.CurWords, word)
}


func indexOf(els []string, el string) int {
  for index, aEl := range els {
    if aEl == el {
        return index
    }
  }
  return -1
}

func (g *MyGame) AddWord(user, word string, points int) int {
  // TODO(dlluncor): Store the entire state of round for an incoming observer.
  // aka, the words and who has found them.
  index := indexOf(g.Users, user)
  g.Points[index] = g.Points[index] + points
  g.CurWords = append(g.CurWords, word)
  return g.Points[index]
}

func (g *MyGame) SetNow() {
  g.Now = time.Now().Unix()
}

func (g *MyGame) SetRoundFetched() {
  g.LastRoundFetched = time.Now().Unix()
}