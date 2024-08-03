package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"os"
	"os/exec"
)

type Answers struct {
	Name  string
	Age   int
	Phone string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: ub <command> [args]")
		return
	}

	cmd := os.Args[1]
	var subCmd string
	if len(os.Args) > 2 {
		subCmd = os.Args[2]
	}
	args := os.Args[2:]

	switch cmd {
	case "tg":
		switch subCmd {
		case "status":
			runGitCommand("status")
		case "add":
			runGitCommand("add", ".")
		default:
			fmt.Printf("Subcommand not found for: 'tg': %s\n", subCmd)
		}

	case "tkdir":
		runCommand("mkdir", args...)
	case "ls":
		runCommand("ls", args...)
	case "rm":
		runCommand("rm", args...)
	case "form":
		runForm()
	default:
		fmt.Printf("Command not found: %s\n", cmd)
	}
}

func runGitCommand(args ...string) {
	runCommand("git", args...)
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error to execute the command '%s %v': %v\n%s", name, args, err, output)
	}
	fmt.Println(string(output))
}

func runForm() {

	questions := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Enter your name:"},
			Validate: survey.Required,
		},
		{
			Name:     "age",
			Prompt:   &survey.Input{Message: "Idade:"},
			Validate: survey.Required,
		},
		{
			Name:     "phone",
			Prompt:   &survey.Input{Message: "Telefone:"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Name  string
		Age   int
		Phone string
	}{}
	err := survey.Ask(questions, &answers)
	if err != nil {
		log.Fatalf("Error to fill up the form: %v", err)
	}

	answersText := Answers{
		Name:  answers.Name,
		Age:   answers.Age,
		Phone: answers.Phone,
	}

	jsonData, err := json.MarshalIndent(answersText, "", "    ")
	if err != nil {
		fmt.Println("Erro ao codificar a struct em JSON:", err)
		return
	}
	create(jsonData)
}

func create(text []byte) {
	file, err := os.Create("answers.json")
	if err != nil {
		fmt.Printf("Erro ao criar arquivo: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(text)
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
	}
	fmt.Println("File created successfully")
}
