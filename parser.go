package tailwindcssparser

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func getHtmlTargetFileContent(tailwindTags string) string {
	return fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head class="%s"></head>></html>`, tailwindTags)
}

func getTailwindConfigFileContent(targetFile string) string {
	return fmt.Sprintf(`{"content": ["%s"]}`, targetFile)
}

func createFile(name, content string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil, err
	}

	return file, nil
}

func getParsedCss(htmlFileName string, minify *bool) (string, error) {
	minifyOption := ""

	if *minify {
		minifyOption = "--minify"
	}

	configFileName := getRandomString(5) + ".json"
	configFileContent := getTailwindConfigFileContent(htmlFileName)
	_, err := createFile(configFileName, configFileContent)
	if err != nil {
		return "", err
	}
	defer os.Remove(configFileName)

	parsedCssFileName := getRandomString(5)
	cmd := exec.Command("npx", "tailwindcss", "-o", parsedCssFileName, minifyOption, "--config", configFileName)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting:", err)
		return "", err
	}

	result := getFileContent(parsedCssFileName)

	err = os.Remove(parsedCssFileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return "", err
	}

	return result, nil
}

func getFileContent(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(content)
}

func GetParsedTailwind(tailwindTags string, minify *bool) (string, error) {
	htmlTargetFileName := getRandomString(5)
	htmlTargetFileContent := getHtmlTargetFileContent(tailwindTags)
	_, err := createFile(htmlTargetFileName, htmlTargetFileContent)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer os.Remove(htmlTargetFileName)

	parsedCss, err := getParsedCss(htmlTargetFileName, minify)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return parsedCss, nil
}
