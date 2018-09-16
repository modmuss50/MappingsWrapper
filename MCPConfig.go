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

	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp_config-%s", version))

	fmt.Println("Cloning MCP Config")

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

	srgPath := filepath.Join(extractedPath, "build/versions/1.13/data/joined.srg")

	copyFile(srgPath, mcpConfigSRGLocation(version))
	return nil
}

func mcpConfigSRGLocation(version string) string {
	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp_config-%s", version))
	srg := filepath.Join(extractedPath, "joined.srg")
	return srg
}

func buildMCPConfigData(version string) MCPData {
	return ReadMCPData(mcpConfigSRGLocation(version))
}
