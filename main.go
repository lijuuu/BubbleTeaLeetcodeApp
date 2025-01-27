package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Leetcode struct {
	No       int    `json:"no"`
	Question string `json:"question"`
	Done     bool   `json:"done"`
}

type model struct {
	LeetcodeArr []Leetcode
	cursor      int
	page        int
	pageSize    int
}

var (
	styleTitle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF6F61")).PaddingBottom(1) // Coral color for title
	styleQuestion    = lipgloss.NewStyle().Foreground(lipgloss.Color("#A2D5F2")).Width(60)                   // Light blue for questions
	styleCursor      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700"))                  // Gold for cursor
	styleDone        = lipgloss.NewStyle().Foreground(lipgloss.Color("#77DD77"))                             // Pastel green for done
	styleNotDone     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6961"))                             // Pastel red for not done
	styleInstruction = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#F5F5F5")).PaddingTop(1)  // Light gray for instructions
	stylePagination  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).PaddingTop(1)               // Gold for pagination
	styleSeparator   = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Render("──────────────────") // Separator line
)

func main() {
	// Step 1: Read the JSON file to get the Leetcode questions
	fileData, err := os.ReadFile("/usr/local/bin/leetcode.json")
	if err != nil {
		log.Fatal("failed to read file data:", err)
	}

	// Step 2: Parse the JSON data into an array of Leetcode questions
	var leetcodeArr []Leetcode
	err = json.Unmarshal(fileData, &leetcodeArr)
	if err != nil {
		log.Fatal("failed to unmarshal JSON data:", err)
	}

	// Step 3: Initialize the model with a default pageSize of 20
	m := model{
		LeetcodeArr: leetcodeArr,
		cursor:      0,
		page:        0,
		pageSize:    20, // Default page size
	}

	// Step 4: Start the Bubble Tea program
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}

func (m model) Init() tea.Cmd {
	// Initial command (optional)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up": // Move cursor up
			if m.cursor > 0 {
				m.cursor--
			}
		case "down": // Move cursor down
			if m.cursor < m.pageSize-1 && m.cursor < len(m.LeetcodeArr)-1-(m.page*m.pageSize) {
				m.cursor++
			}
		case "enter":
			// Toggle "done" status of the current question
			idx := m.page*m.pageSize + m.cursor
			if idx < len(m.LeetcodeArr) {
				m.LeetcodeArr[idx].Done = !m.LeetcodeArr[idx].Done
			}

			// Update the JSON file with the new data
			err := updateJSONFile(m.LeetcodeArr)
			if err != nil {
				log.Fatal("Failed to update JSON file:", err)
			}
		case "left": // Previous page
			if m.page > 0 {
				m.page--
				m.cursor = 0 // Reset cursor to the top of the page
			}
		case "right": // Next page
			if (m.page+1)*m.pageSize < len(m.LeetcodeArr) {
				m.page++
				m.cursor = 0 // Reset cursor to the top of the page
			}
		}
	}

	// Return the updated model
	return m, nil
}

func (m model) View() string {
	var sb strings.Builder

	// Print the title
	sb.WriteString(styleTitle.Render("Leetcode Questions"))
	sb.WriteString("\n" + styleSeparator + "\n")

	// Calculate the start and end indices for the current page
	startIdx := m.page * m.pageSize
	endIdx := startIdx + m.pageSize
	if endIdx > len(m.LeetcodeArr) {
		endIdx = len(m.LeetcodeArr)
	}

	// Display questions for the current page
	for i := startIdx; i < endIdx; i++ {
		q := m.LeetcodeArr[i]
		cursor := "  "
		if i == m.page*m.pageSize+m.cursor {
			cursor = styleCursor.Render("➤") // Dynamic cursor
		}

		done := styleNotDone.Render("✘")
		if q.Done {
			done = styleDone.Render("✔")
		}

		// Add spacing and separators between questions
		sb.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, done, styleQuestion.Render(q.Question)))
		if i < endIdx-1 {
			sb.WriteString(styleSeparator + "\n")
		}
	}

	// Instructions and pagination
	sb.WriteString("\n")
	sb.WriteString(styleInstruction.Render("↑/↓: Navigate • Enter: Toggle Done • ←/→: Change Page • q: Quit\n"))
	sb.WriteString(stylePagination.Render(fmt.Sprintf("Page %d of %d", m.page+1, (len(m.LeetcodeArr)+m.pageSize-1)/m.pageSize)))

	return sb.String()
}

// Helper function to update the leetcode.json file with the updated questions list
func updateJSONFile(data []Leetcode) error {
	// Marshal the updated data into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated data to the file
	err = os.WriteFile("leetcode.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}