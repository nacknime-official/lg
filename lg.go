// Сreated a new structure based on the logger structure of the standard Lg package.
package lg

import (
	"fmt"
	"log"
	"os"
)

type lg struct {
	c   *color
	log *log.Logger
}

func loggerCreateNew(std *os.File, colorVal string, prefix string) *lg {
	log := &lg{
		c:   newColor(),
		log: log.New(std, colorVal+prefix+"\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
	if colorVal == "" {
		log.c = &color{}
	}
	return log
}

func (l *lg) print(v ...interface{}) error {
	return l.log.Output(3, fmt.Sprint(v...)+l.c.Reset)
}
func (l *lg) println(v ...interface{}) error {
	if l.c.Reset == "" {
		l.c.Reset = "\n"
	}
	return l.log.Output(3, fmt.Sprintln(v...)+l.c.Reset)
}
func (l *lg) printf(format string, v ...interface{}) error {
	return l.log.Output(3, fmt.Sprintf(format, v...)+l.c.Reset)
}

func (l *lg) fatal(v ...interface{}) error {
	if err := l.log.Output(3, fmt.Sprint(v...)+l.c.Reset); err != nil {
		return err
	}
	os.Exit(1)
	return nil
}
func (l *lg) fatalln(v ...interface{}) error {
	if l.c.Reset == "" {
		l.c.Reset = "\n"
	}
	if err := l.log.Output(3, fmt.Sprintln(v...)+l.c.Reset); err != nil {
		return err
	}
	os.Exit(1)
	return nil
}
func (l *lg) fatalf(format string, v ...interface{}) error {
	if err := l.log.Output(3, fmt.Sprintf(format, v...)+l.c.Reset); err != nil {
		return err
	}
	os.Exit(1)
	return nil
}

func (l *lg) panic(v ...interface{}) error {
	s := fmt.Sprint(v...) + l.c.Reset
	if err := l.log.Output(3, s); err != nil {
		return err
	}
	panic(s)
}
func (l *lg) panicln(v ...interface{}) error {
	if l.c.Reset == "" {
		l.c.Reset = "\n"
	}
	s := fmt.Sprintln(v...) + l.c.Reset
	if err := l.log.Output(3, s); err != nil {
		return err
	}
	panic(s)
}
func (l *lg) panicf(format string, v ...interface{}) error {
	s := fmt.Sprintf(format, v...) + l.c.Reset
	if err := l.log.Output(3, s); err != nil {
		return err
	}
	panic(s)
}
