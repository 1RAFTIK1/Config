package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
	"gopkg.in/ini.v1"
)

type Config struct {
	VisualizationProgram string
	PackageName          string
	OutputImagePath      string
	MaxDepth             int
	RepositoryURL        string
}

func readConfig(filepath string) (*Config, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}
	return &Config{
		VisualizationProgram: cfg.Section("settings").Key("visualization_program").String(),
		PackageName:          cfg.Section("settings").Key("package_name").String(),
		OutputImagePath:      cfg.Section("settings").Key("output_image_path").String(),
		MaxDepth:             cfg.Section("settings").Key("max_depth").MustInt(),
		RepositoryURL:        cfg.Section("settings").Key("repository_url").String(),
	}, nil
}

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

func buildGraph(pkg string, graph *simple.DirectedGraph, visited map[string]struct{}, depth int, maxDepth int) {
	if depth > maxDepth {
		return
	}

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

		buildGraph(dep, graph, visited, depth+1, maxDepth)
	}
}

func saveGraphAsPNG(graph *simple.DirectedGraph, filename string) error {
	// Получаем DOT-представление графа и проверяем на ошибки
	dotData, err := dot.Marshal(graph, "Graph", "", "  ")
	if err != nil {
		return err
	}

	// Записываем DOT в файл
	dotFile := "graph.dot"
	if err := os.WriteFile(dotFile, dotData, 0644); err != nil {
		return err
	}

	// Генерируем PNG с помощью Graphviz
	cmd := exec.Command("dot", "-Tpng", dotFile, "-o", filename)
	return cmd.Run()
}

func main() {
	config, err := readConfig("E:\\Emu_ter_go\\Config\\Go_task2\\config.ini")
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	graph := simple.NewDirectedGraph()
	visited := make(map[string]struct{})

	buildGraph(config.PackageName, graph, visited, 0, config.MaxDepth)

	err = saveGraphAsPNG(graph, config.OutputImagePath)
	if err != nil {
		fmt.Printf("Error saving graph: %v\n", err)
		return
	}
	fmt.Printf("Graph saved as %s\n", config.OutputImagePath)
}
