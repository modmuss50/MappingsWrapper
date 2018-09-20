package mcpwrapper

import (
	"bytes"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getMCPConfigData(version string) (MCPData, error) {
	var data MCPData
	if !fileExists(mcpConfigSRGLocation(version)) {
		err := prepareMCPConfig(version)
		if err != nil {
			return data, err
		}
	}
	data = buildMCPConfigData(version)
	return data, nil
}

func prepareMCPConfig(version string) error {

	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp-%s-config", version))

	fmt.Println("Cloning MCP Config")

	//Handles errors that can be caused if a previous build failed
	if fileExists(extractedPath) {
		deleteDir(extractedPath)
	}

	_, err := git.PlainClone(extractedPath, false, &git.CloneOptions{
		URL:      "https://github.com/MinecraftForge/MCPConfig",
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	fmt.Println("Making SRG")

	//TODO linux
	cmd := exec.Command("cmd", "/c", "gradlew", version+":makeSrg")
	cmd.Dir = extractedPath

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	log.Println(stdBuffer.String())

	srgPath := filepath.Join(extractedPath, "build", "versions", version, "data", "joined.srg")

	fmt.Println(srgPath)
	fmt.Println(mcpConfigSRGLocation(version))

	return copyFile(srgPath, mcpConfigSRGLocation(version))
}

func mcpConfigSRGLocation(version string) string {
	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp-%s-config", version), "joined.srg")
	return extractedPath
}

func buildMCPConfigData(version string) MCPData {
	return ReadMCPData(mcpConfigSRGLocation(version))
}
