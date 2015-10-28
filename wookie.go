package wookie

import (
	"fmt"
	"strings"
)

const minLengthMatching = 3

type Part struct {
	Line   string
	Rights map[int][]*Part
	Lefts  map[int][]*Part
}

type Wookie struct {
	Lines []string
	Parts []*Part
}

func NewPart(line string) Part {
	return Part{
		Line:   line,
		Rights: make(map[int][]*Part, 0),
		Lefts:  make(map[int][]*Part, 0),
	}
}

func NewWookie(lines []string) Wookie {
	return Wookie{
		Lines: lines,
		Parts: make([]*Part, 0),
	}
}

type Genome struct {
	Start *Part
}

func getEndLength(left, right string) int {
	leftLen := len(left)
	rightLen := len(right)
	maxLength := leftLen
	if rightLen < maxLength {
		maxLength = rightLen
	}

	for i := maxLength; i >= minLengthMatching; i-- {
		if strings.Index(left[:i], right[rightLen-i:]) != -1 {
			//fmt.Println(left, right, left[:i], right[rightLen-i:], i, idx)
			return i
		}
	}
	return -1
}

func (w *Wookie) Compute() {
	for _, line := range w.Lines {
		part := NewPart(line)

		for _, existingPart := range w.Parts {
			if endLength := getEndLength(line, existingPart.Line); endLength > -1 {
				part.Lefts[endLength] = append(part.Lefts[endLength], existingPart)
				existingPart.Rights[endLength] = append(existingPart.Rights[endLength], &part)
			}
		}

		w.Parts = append(w.Parts, &part)
	}
}

func NewGenome(start *Part) Genome {
	return Genome{
		Start: start,
	}
}

func (p *Part) MaxLength() int {
	maxLength := 0
	for length, _ := range p.Rights {
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func (p *Part) Next() (*Part, int) {
	maxLength := p.MaxLength()
	if maxLength == -1 {
		return nil, 0
	}
	maxRights := p.Rights[maxLength]
	switch len(maxRights) {
	case 0:
		return nil, 0
	case 1:
		return maxRights[0], maxLength
	default:
		fmt.Println(maxLength, maxRights)
		//panic("Multiple rights with max length, not handled")
		return nil, 0
	}
}

func (g *Genome) String() string {
	output := ""

	leftPart := g.Start
	rightPart := g.Start
	maxLength := 0
	for rightPart != nil {
		//fmt.Println(leftPart.Line, rightPart.Line, maxLength)
		// output += fmt.Sprintf(" %q %d", rightPart.Line, maxLength)
		output += rightPart.Line[maxLength:]
		leftPart = rightPart
		rightPart, maxLength = leftPart.Next()
		// output += g.Start.Line
		//output += fmt.Sprintf(" %d", g.Start.MaxLength())
		//output += fmt.Sprintf(" %p", g.Start.Next())
	}

	return output
}

func (w *Wookie) GetGenomes() []Genome {
	genomes := make([]Genome, 0)
	for _, part := range w.Parts {
		if len(part.Lefts) == 0 {
			genome := NewGenome(part)
			genomes = append(genomes, genome)
		}
	}
	return genomes
}
