package plot

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func Open(dbPath string) error {
	gnucmd := exec.Command("gnuplot", "-p")

	t, err := template.New("").Parse(plotConfigTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	stdin := bytes.Buffer{}

	err = t.Execute(&stdin, struct {
		DatabasePath string
	}{
		DatabasePath: dbPath,
	})
	if err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	gnucmd.Stdout = os.Stdout
	gnucmd.Stderr = os.Stderr
	gnucmd.Stdin = &stdin

	err = gnucmd.Run()
	if err != nil {
		return fmt.Errorf("running command: %w", err)
	}

	return nil
}

const plotConfigTemplate = `
# https://raymii.org/s/tutorials/GNUplot_tips_for_nice_looking_charts_from_a_CSV_file.html

set datafile separator ','

set title "Workouts"
set key autotitle columnhead
set autoscale

set xlabel "time"
set xtic auto
set xdata time
set timefmt "%d.%m.%Y"
set format x "%d.%m.%y"

set ylabel "repetitions"
set ytic auto
set yrange[0:30]

plot "{{ .DatabasePath }}" \
	using 1:2 with lines, \
	'' using 1:3 with lines, \
	'' using 1:4 with lines, \
	'' using 1:5 with lines,
`
