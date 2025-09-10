package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/g-restante/GopeherKit.Test/internal"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	
	switch command {
	case "generate-mock":
		if len(os.Args) < 4 {
			fmt.Println("Usage: gopherkit-test generate-mock <interface-file> <output-dir>")
			os.Exit(1)
		}
		generateMock(os.Args[2], os.Args[3])
		
	case "generate-test":
		if len(os.Args) < 4 {
			fmt.Println("Usage: gopherkit-test generate-test <package-path> <output-dir>")
			os.Exit(1)
		}
		generateTestBoilerplate(os.Args[2], os.Args[3])
		
	case "generate-assertions":
		if len(os.Args) < 4 {
			fmt.Println("Usage: gopherkit-test generate-assertions <output-dir> <spec1> [spec2] ...")
			fmt.Println("Spec format: name:params:condition:defaultMessage")
			os.Exit(1)
		}
		generateAssertions(os.Args[2], os.Args[3:])
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("GopherKit.Test Code Generator")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  gopherkit-test generate-mock <interface-file> <output-dir>")
	fmt.Println("  gopherkit-test generate-test <package-path> <output-dir>")
	fmt.Println("  gopherkit-test generate-assertions <output-dir> <spec1> [spec2] ...")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  gopherkit-test generate-mock ./example/user_service.go ./mocks")
	fmt.Println("  gopherkit-test generate-test mypackage ./tests")
	fmt.Println("  gopherkit-test generate-assertions ./assert \"IsPositive:value int:value > 0:expected positive value\"")
}

func generateMock(interfaceFile, outputDir string) {
	packageName := filepath.Base(filepath.Dir(interfaceFile))
	generator := internal.NewGenerator(packageName, outputDir)
	
	fmt.Printf("Generating mock for interface in %s...\n", interfaceFile)
	
	err := generator.GenerateMocks([]string{interfaceFile})
	if err != nil {
		fmt.Printf("Error generating mock: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Mock generated successfully in %s\n", outputDir)
}

func generateTestBoilerplate(packagePath, outputDir string) {
	packageName := filepath.Base(packagePath)
	generator := internal.NewGenerator(packageName, outputDir)
	
	fmt.Printf("Generating test boilerplate for package %s...\n", packagePath)
	
	err := generator.GenerateTestBoilerplate(packagePath)
	if err != nil {
		fmt.Printf("Error generating test boilerplate: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Test boilerplate generated successfully in %s\n", outputDir)
}

func generateAssertions(outputDir string, specs []string) {
	generator := internal.NewGenerator("assert", outputDir)
	
	fmt.Printf("Generating custom assertions...\n")
	
	err := generator.GenerateAssertions(specs)
	if err != nil {
		fmt.Printf("Error generating assertions: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Custom assertions generated successfully in %s\n", outputDir)
}
