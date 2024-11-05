package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

func getDependencies(pkg string) ([]string, error) {
	cmd := exec.Command("apk", "info", "--depends", pkg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(&out)
	dependencies := []string{}
	for scanner.Scan() {
		dependencies = append(dependencies, scanner.Text())
	}

	return dependencies, scanner.Err()
}

func buildGraph(pkg string, graph *simple.DirectedGraph, visited map[string]struct{}) {
	if _, exists := visited[pkg]; exists {
		return
	}
	visited[pkg] = struct{}{}

	node := graph.NewNode()
	graph.AddNode(node)

	dependencies, err := getDependencies(pkg)
	if err != nil {
		fmt.Printf("Error getting dependencies for %s: %v\n", pkg, err)
		return
	}

	for _, dep := range dependencies {
		depNode := graph.NewNode()
		graph.AddNode(depNode)
		graph.SetEdge(graph.NewEdge(node, depNode))

		buildGraph(dep, graph, visited)
	}
}

func saveGraphAsPNG(graph *simple.DirectedGraph, filename string) error {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	if err := dot.Write(w, graph); err != nil {
		return err
	}
	w.Flush()

	cmd := exec.Command("dot", "-Tpng", "-o", filename)
	cmd.Stdin = &buf
	return cmd.Run()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <package_name>")
		return
	}

	pkg := os.Args[1]
	graph := simple.NewDirectedGraph(0, 0)
	visited := make(map[string]struct{})

	buildGraph(pkg, graph, visited)

	err := saveGraphAsPNG(graph, "dependencies.png")
	if err != nil {
		fmt.Printf("Error saving graph: %v\n", err)
		return
	}

	fmt.Println("Graph saved as dependencies.png")
}
