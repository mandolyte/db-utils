package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var wnode *csv.Writer
var wedge *csv.Writer
var nodemap map[string]struct{}
var edgemap map[string]struct{}
var nodecount uint64
var edgecount uint64

func main() {
	input := flag.String("i", "", "Input CSV filename; default STDIN")
	node := flag.String("node", "", "Output Node CSV filename; required")
	edge := flag.String("edge", "", "Output Edge CSV filename; required")
	help := flag.Bool("help", false, "Show usage message")
	flag.Parse()

	if *help {
		usage("Help Message")
	}

	if *node == "" {
		usage("Node filename is missing")
	}

	if *edge == "" {
		usage("Edge filename is missing")
	}

	// init the maps
	nodemap = make(map[string]struct{})
	edgemap = make(map[string]struct{})

	// open node output file
	fnode, foerr := os.Create(*node)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fnode.Close()
	wnode = csv.NewWriter(fnode)
	defer wnode.Flush()

	// open edge output file
	fedge, foerr := os.Create(*edge)
	if foerr != nil {
		log.Fatal("os.Create() Error:" + foerr.Error())
	}
	defer fedge.Close()
	wedge = csv.NewWriter(fedge)
	defer wedge.Flush()

	// open input file
	var r *csv.Reader
	fi, fierr := os.Open(*input)
	if fierr != nil {
		log.Fatal("os.Open() Error:" + fierr.Error())
	}
	defer fi.Close()
	r = csv.NewReader(fi)
	allrows, rerr := r.ReadAll() 
	if rerr != nil {
		log.Fatalf("csv.ReadAll() Error: %v\n", rerr)
	}
	
	var row uint64
	for _, line := range allrows {
		row++
		sLine := line[0]
		// the line is a full pathname so the
		// terminating element is a filename leaf
		child := path.Base(sLine)

		// add the node (as needed)
		addNode(child)
	
		// now get the directory
		parent := path.Dir(sLine)
		
		// add the edge between them
		addEdge(parent,child)

		if parent == "/" {
			// at the top already... don't recurse
			continue
		}

		recurse(parent)
	}
	
	fmt.Printf("Processed %v full paths.\n", row)
	fmt.Printf("Number of nodes: %v.\n", nodecount)
	fmt.Printf("Number of edges: %v.\n", edgecount)
}

func addEdge(parent, child string) {
	edgeID := path.Join(parent, child)
	if _, ok := edgemap[edgeID]; ! ok {
		err := wedge.Write([]string{edgeID, parent, child})
		if err != nil {
			log.Fatalf("wedge.Write() Error: %v\n", err)
		}
		edgecount++

		// add edge to map
		edgemap[edgeID] = struct{}{}
	}
}

func addNode(n string) {
	if _, ok := nodemap[n]; !ok {
		// write out the node
		err := wnode.Write([]string{n})
		if err != nil {
			log.Fatalf("csv.Write() Error: %v\n", err)
		}
		nodecount++

		// add node to map
		nodemap[n] = struct{}{}
	}
}

func recurse(child string) {
	if child == "." {
		log.Fatal("Recursion found a dot.\n")
	}
	// get the current terminal element of path
	parent := path.Dir(child)

	// add the node (as needed)
	addNode(parent)

	// add the edge
	addEdge(parent,child)
	
	if parent == "/" {
		// at the top already... don't recurse
		return
	}
	recurse(parent)
}

func usage(msg string) {
	fmt.Println(msg + "\n")
	flag.PrintDefaults()
	os.Exit(0)
}

/* code graveyard
func recurse(parentdir, sChildhash string) {
	//fmt.Printf("Enter recurse() with parent=%v, child=%v\n", parentdir, sChildhash)
	if sChildhash == "." {
		log.Fatal("Recursion found a dot.\n")
	}
	// get the current terminal element of path
	newchild := path.Base(parentdir)
	newchildhash := sha256.Sum256([]byte(newchild))
	sNewChildhash := fmt.Sprintf("%x", newchildhash)
	if _, ok := nodemap[sNewChildhash]; !ok {
		// write out the node
		err := wnode.Write([]string{sNewChildhash, newchild})
		if err != nil {
			log.Fatalf("csv.Write() Error: %v\n", err)
		}
		nodecount++
		// add node to map
		nodemap[sNewChildhash] = struct{}{}
	}
	edgecount++
	// now create a string combining parent and child hash
	edgeValue := sNewChildhash + sChildhash + fmt.Sprintf("%v",edgecount)
	// get the hash
	edgeID := sha256.Sum256([]byte(edgeValue))
	// now I have the edge id and the from and to ids
	// write the edge out
	sEdgeID := fmt.Sprintf("%x", edgeID)
	err := wedge.Write([]string{sEdgeID, sNewChildhash, sChildhash})
	if err != nil {
		log.Fatalf("wedge.Write() Error: %v\n", err)
	}
	// now get the directory
	newparentdir := path.Dir(parentdir)
	if newparentdir == "/" {
		// at the top already... don't recurse
		return
	}
	recurse(newparentdir, sNewChildhash)
}


*/