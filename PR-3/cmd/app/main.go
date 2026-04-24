package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gitlab.com/msv/pr3-calculator-v12/internal/app"
	"gitlab.com/msv/pr3-calculator-v12/internal/domain"
	"gitlab.com/msv/pr3-calculator-v12/internal/patterns"
	"gitlab.com/msv/pr3-calculator-v12/internal/service"
)

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(stderr)

	op := fs.String("op", "", "операция: add|sub|mul|div|pow")
	left := fs.Float64("a", 0, "первый операнд")
	right := fs.Float64("b", 0, "второй операнд")
	precision := fs.Int("precision", 2, "количество знаков после запятой")
	verbose := fs.Bool("verbose", false, "включить логирование вычисления")
	demo := fs.Bool("demo", false, "запустить демонстрационный сценарий")

	fs.Usage = func() {
		fmt.Fprintf(stderr, "Использование: %s [опции]\n", fs.Name())
		fmt.Fprintln(stderr, "Выполняет арифметическую операцию над двумя числами.")
		fmt.Fprintln(stderr, "\nОпции:")
		fs.PrintDefaults()
		fmt.Fprintln(stderr, "\nПримеры:")
		fmt.Fprintln(stderr, "  app -op add -a 10 -b 5")
		fmt.Fprintln(stderr, "  app -op div -a 7 -b 2 -precision 3 -verbose")
		fmt.Fprintln(stderr, "  app -demo")
	}

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		fmt.Fprintf(stderr, "ошибка разбора аргументов: %v\n", err)
		return 2
	}

	logger := log.New(stdout, "[calc] ", 0)
	factory := patterns.NewOperationFactory()
	calculator := service.NewCalculator(factory, logger)

	if *demo {
		if err := app.RunDemo(calculator, stdout); err != nil {
			fmt.Fprintf(stderr, "ошибка демонстрационного сценария: %v\n", err)
			return 1
		}
		return 0
	}

	if fs.NFlag() == 0 || *op == "" {
		fs.Usage()
		return 1
	}

	request := domain.Request{
		Operation: *op,
		Left:      *left,
		Right:     *right,
		Precision: *precision,
		Verbose:   *verbose,
	}

	result, err := calculator.Calculate(request)
	if err != nil {
		fmt.Fprintf(stderr, "ошибка: %v\n", err)
		return 1
	}

	fmt.Fprintln(stdout, result.String())
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
