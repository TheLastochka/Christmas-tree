package main

import (
	"os"
	"os/exec"
	"time"

	color "github.com/fatih/color"
	"github.com/pterm/pterm"
)

func screenClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	screenClear()
	clr := make(map[byte]func(a ...interface{}) string)
	clr['Y'] = color.New(color.BgYellow).SprintFunc()
	clr['B'] = color.New(color.BgYellow).SprintFunc()
	clr[' '] = color.New(color.BgBlack).SprintFunc()

	clr['G'] = color.New(color.BgGreen).SprintFunc()
	clr['O'] = color.New(color.BgGreen).SprintFunc()

	clr['E'] = color.New(color.BgBlue).SprintFunc()
	clr['R'] = color.New(color.BgRed).SprintFunc()
	clr['M'] = color.New(color.BgMagenta).SprintFunc()
	clr['C'] = color.New(color.BgCyan).SprintFunc()

	buff := ""
	format := "" +
		"                 Y                 \n" +
		"                YYY                \n" +
		"                 O                 \n" +
		"                OOO                \n" +
		"                ORO                \n" +
		"               OOOOO               \n" +
		"               OMOGO               \n" +
		"              OOOOOOO              \n" +
		"              OROCOMO              \n" +
		"             OOOOOOOOO             \n" +
		"             OCOMOROEO             \n" +
		"            OOOOOOOOOOO            \n" +
		"            OMOEOROGOCO            \n" +
		"           OOOOOOGOOOOOO           \n" +
		"           OEOMOOGOOROGO           \n" +
		"          OOOOOOOOOOOOOOO          \n" +
		"                 B                 \n" +
		"                 B                 \n" +
		"                 B                 "

	area, _ := pterm.DefaultArea.WithCenter().Start()
	for {
		h := pterm.GetTerminalHeight()
		head := make([]byte, int((h-21)/2))
		for i := range head {
			head[i] = '\n'
		}
		new_str := []byte(format)
		for _, chr := range new_str {
			if chr == '\n' {
				buff += "\n"
			} else {
				buff += clr[chr]("  ")
			}
		}
		area.Update(string(head) + buff)
		time.Sleep(500 * time.Millisecond)

		for i, chr := range format {
			switch chr {
			case 'E':
				new_str[i] = 'R'
			case 'R':
				new_str[i] = 'G'
			case 'G':
				new_str[i] = 'M'
			case 'M':
				new_str[i] = 'C'
			case 'C':
				new_str[i] = 'E'
			}
		}
		format = string(new_str)
		buff = ""
	}
}
