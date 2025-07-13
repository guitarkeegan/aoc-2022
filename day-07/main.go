package main

import (
	"fmt"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"
)

var dbg = func() func(format string, args ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(format string, args ...any) {}
	}
	return func(format string, args ...any) {
		fmt.Printf(format+"\n", args...)
	}
}()

type TreeNode struct {
	Size     int
	Path     string
	Children map[string]*TreeNode
	Parent   *TreeNode
}

func (tn *TreeNode) String() string {

	var acc string
	var bfs func()
	bfs = func() {
		q := []*TreeNode{tn}
		length := len(q)
		for length > 0 {

			var children []string
			left := q[0]
			q = q[1:]
			for _, node := range left.Children {
				q = append(q, node)
				children = append(children, node.Path)
			}
			if left.Parent != nil {
				// is dir
				if left.Size == 0 {
					acc += fmt.Sprintf("%s, parent: %s, children: %v\n", left.Path, left.Parent.Path, children)
				} else {
					acc += fmt.Sprintf("%s, parent: %s, size: %v\n", left.Path, left.Parent.Path, left.Size)
				}
			}
			length = len(q)
		}
	}
	bfs()

	return acc
}

type User struct {
	Cur    *TreeNode
	Parent *TreeNode
}

type OS struct {
	lines iter.Seq[string]
	root  *TreeNode
	tree  *TreeNode
	user  *User
}

type Cmd int

const (
	cd Cmd = iota
	ls
	file
	dir
)

func (o *OS) Build() {

	for line := range o.lines {
		o.Handler(line)
	}

	fmt.Println("Build Complete!")

}

func (o *OS) Is(line string) Cmd {
	if strings.HasPrefix(line, "$") {

		if strings.HasSuffix(line, "ls") {
			return ls
		}
		return cd

	}
	// is a file or folder
	if strings.HasPrefix(line, "dir") {
		return dir
	}

	return file

}

func (o *OS) GetPath(n string) string {

	dbg("getPath")
	if o.tree == nil {
		log.Fatal("tree was not initialized\n")
	}

	if o.tree.Parent == nil {
		return "/" + n
	}

	path := o.tree.Path
	dbg("  path: %s", path)

	return path + "/" + n

}

func (o *OS) Handler(line string) {
	cmd := o.Is(line)
	args := strings.Split(line, " ")

	switch cmd {
	case cd:
		dbg("%s", line)
		to := args[2]
		switch to {
		case "/":
			if o.tree == nil {
				dbg("create tree")
				o.tree = &TreeNode{
					Path:     "/",
					Children: map[string]*TreeNode{},
				}
				o.root = o.tree
				o.user.Cur = o.tree
				return
			} else {
				o.user.Cur = o.root
				return
			}
		case "..":
			if o.tree.Parent != nil {
				o.tree = o.tree.Parent
				o.user.Cur = o.tree
				return
			} else {
				log.Fatalf("cannot go up beyond root. cur: %s\n", o.user.Cur.Path)
			}
		default:
			path := args[2]
			absPath := o.GetPath(path)
			if nextDir, ok := o.tree.Children[absPath]; ok {
				o.tree = nextDir
				o.user.Cur = o.tree
				return
			}
			log.Fatalf("no child found at path: %s, absPath: %s, line: %s\nchildren: %+v", path, absPath, line, o.tree.Children)
		}
	case ls:
	case dir:
		dbg("dir")
		path := o.GetPath(args[1])
		if _, ok := o.tree.Children[path]; !ok {
			dbg("  new dir: path: %s", path)
			newNode := &TreeNode{
				Path:     path,
				Children: map[string]*TreeNode{},
				Parent:   o.tree,
			}
			o.tree.Children[path] = newNode
			// update user
			o.user.Cur = o.tree
			return
		}
	case file:
		dbg("file")
		path := o.GetPath(args[1])
		size, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("not a file! cmd: %d, %s\n", cmd, line)
		}
		if _, ok := o.tree.Children[path]; !ok {
			dbg("  new file: path: %s, size: %d", path, size)
			newNode := &TreeNode{
				Path:     path,
				Children: map[string]*TreeNode{},
				Parent:   o.tree,
				Size:     size,
			}
			o.tree.Children[path] = newNode
			// update user
			o.user.Cur = o.tree
			return
		}
	default:
		log.Fatalf("unknown cmd: %d, line: %s\n", cmd, line)
	}
}

func (o *OS) dirSumUnder100K() int {

	cur := o.root

	var total int

	var dfs func(tn *TreeNode) int
	dfs = func(tn *TreeNode) int {

		if tn == nil {
			return 0
		}
		var acc int
		for _, child := range tn.Children {
			acc += dfs(child)
		}

		// if dir
		if tn.Size == 0 && acc < 100_000 {
			total += acc
		}

		return tn.Size + acc
	}

	dfs(cur)

	return total

}

func (o *OS) dirSizes() map[string]int {

	cur := o.root

	dirs := make(map[string]int)

	var dfs func(tn *TreeNode) int
	dfs = func(tn *TreeNode) int {

		if tn == nil {
			return 0
		}
		var acc int
		for _, child := range tn.Children {
			acc += dfs(child)
		}

		// if dir
		if tn.Size == 0 {
			dirs[tn.Path] = acc
		}

		return tn.Size + acc
	}

	dfs(cur)

	return dirs

}

func smallestToDel(m map[string]int) int {
	const (
		MAX_SIZE         = 70_000_000
		MIN_UNUSED_SPACE = 30_000_000
	)

	sysSize := m["/"]
	availableSpace := MAX_SIZE - sysSize
	minSize := MIN_UNUSED_SPACE - availableSpace
	var candidate = MAX_SIZE

	for path, size := range m {
		if path != "/" && size >= minSize {
			candidate = min(candidate, size)
		}
	}

	return candidate

}

func main() {

	lines := strings.SplitSeq(input, "\n")
	oss := &OS{lines: lines, user: &User{}}
	oss.Build()
	// fmt.Print(oss.root)
	// fmt.Println(oss.dirSumUnder100K())
	fmt.Println(smallestToDel(oss.dirSizes()))
}
