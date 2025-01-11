package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runCommand(cmd string) error {
	// Split command string into arguments and run it
	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func pushToGit(outputDir, branchName string) error {
	// Change to the output directory
	if err := os.Chdir(outputDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", outputDir, err)
	}

	// Initialize a git repository in the output directory
	if err := runCommand("git init"); err != nil {
		return fmt.Errorf("failed to initialize git repository: %v", err)
	}

	// Add all files to git staging area
	if err := runCommand("git add ."); err != nil {
		return fmt.Errorf("failed to add files to git: %v", err)
	}

	// Commit the changes
	commitMessage := fmt.Sprintf("Deploy-%s", time.Now().Format("2006-01-02T15:04:05"))
	if err := runCommand(fmt.Sprintf("git commit -m '%s'", commitMessage)); err != nil {
		return fmt.Errorf("failed to commit files: %v", err)
	}

	remoteURL := "git@github.com:IndieCoderMM/indiecoder.git"
	if err := runCommand(fmt.Sprintf("git remote add origin %s", remoteURL)); err != nil {
		return fmt.Errorf("failed to add remote URL: %v", err)
	}

	// Push the commit to the specified branch
	if err := runCommand(fmt.Sprintf("git push -f origin main:%s", branchName)); err != nil {
		return fmt.Errorf("failed to push to git branch %s: %v", branchName, err)
	}

	return nil
}

func main() {
	// Parse command line flags
	buildCmd := flag.String("build", "make build", "Command to build the project")
	outputDir := flag.String("output", "dist", "Output directory to push to git")
	branchName := flag.String("branch", "gh-pages", "Git branch to push the output")
	flag.Parse()

	// Run the build command
	fmt.Println("Running build command...")
	if err := runCommand(*buildCmd); err != nil {
		log.Fatalf("Build failed: %v\n", err)
	}

	// Push the build output to the specified Git branch
	fmt.Printf("Pushing output directory %s to branch %s...\n", *outputDir, *branchName)
	if err := pushToGit(*outputDir, *branchName); err != nil {
		log.Fatalf("Error pushing to Git: %v\n", err)
	}

	fmt.Println("Deployment successful! 🎉")
}
