package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: ub <command> [args]")
		return
	}

	// Capturar o comando passado como argumento
	//cmd := os.Args[1]
	//args := os.Args[2:]

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

// Função para executar comandos Git
func runGitCommand(args ...string) {
	runCommand("git", args...)
}

// Função para executar comandos do sistema
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
	fmt.Printf("Nome: %s\nIdade: %d\nTelefone: %s\n", answers.Name, answers.Age, answers.Phone)
}
