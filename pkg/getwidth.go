package pkg
import (
	"fmt"
	"os"
	"os/exec"
)
func GetTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	var width int
	fmt.Sscanf(string(out), "%d", &width)
	return width, nil
}
