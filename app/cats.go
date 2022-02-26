package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Run the application.
func TheCatMeows() {
	if os.Getenv("KOSHKA_DEBUG") != "" {
		if f, err := tea.LogToFile("koshka.log", "help"); err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		} else {
			defer f.Close()
		}
	}

	if err := tea.NewProgram(
		Кошка(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	).Start(); err != nil {
		fmt.Printf("The cat did not wake up. ERROR: %v", err)
		os.Exit(1)
	}
}
