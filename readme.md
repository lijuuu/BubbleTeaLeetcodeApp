## Repository: Leetcode Todo Terminal

A terminal-based interactive tool to manage and track your progress on Leetcode questions. (Made using DeepSeek V3)

---

### Key Features:
- **Interactive Navigation**: Easily browse through your Leetcode questions with pagination support.
- **Progress Tracking**: Mark questions as "done" or "not done" with just a keystroke.
- **User-Friendly Interface**: Dynamic and visually appealing interface powered by the `bubbletea` and `lipgloss` libraries.
- **JSON Persistence**: Your progress is saved in a JSON file (`leetcode.json`), ensuring your data is never lost.
- **Custom Setup Script**: Quickly set up the application with the included `run.sh` script.

---

### Installation:

1. Clone this repository:
   ```bash
   git clone https://github.com/lijuu/BubbleTeaLeetcodeApp.git
   cd BubbleTeaLeetcodeApp
   ```

2. Run the `run.sh` script to install:
   ```bash
   chmod +x run.sh
   ./run.sh
   ```

---

### Usage:

After setup, simply type `leetcode` in the terminal to launch the app.

---

### Controls:

- **Navigation**: Use `↑` and `↓` to move through the list.
- **Toggle Status**: Press `Enter` to mark a question as "done" or "not done".
- **Pagination**: Use `←` and `→` to navigate through pages.
- **Quit**: Press `q` to exit the application.

---

### Example Setup Output:

Running `run.sh` will copy the `leetcode` binary and `leetcode.json` file to `/usr/local/bin`, enabling global access to the app.

---

### Dependencies:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea): For the terminal UI.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss): For styling the terminal output.

---

Feel free to contribute by opening issues or submitting pull requests!