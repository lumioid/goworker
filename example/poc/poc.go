package main

import (
	"fmt"
	"strings"
)

type process struct {
	Hostname string
	Pid      int
	ID       string
	Queues   []string
}

func (p *process) String() string {
	return fmt.Sprintf("%s:%d-%s:%s", p.Hostname, p.Pid, p.ID, strings.Join(p.Queues, ","))
}

func main() {
	p1 := process{
		Hostname: "hostname",
		Pid:      12345,
		ID:       "123",
		Queues:   []string{"high", "low"},
	}

	fmt.Println(p1.String())
	fmt.Printf(fmt.Sprintf("%sstat:processed:%v", "queue", p1))
	fmt.Printf(fmt.Sprintf("%sstat:processed:%s", "queue", p1))
}
