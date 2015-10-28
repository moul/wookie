package wookie

import "strings"

const minLengthMatching = 3

type Part struct {
	Line   string
	Rights map[int][]*Part
	Lefts  map[int][]*Part
	Done   bool
}

type Wookie struct {
	Lines    []string
	Parts    []*Part
	Genome   Genome
	Missings int
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
		Lines:  lines,
		Parts:  make([]*Part, 0),
		Genome: NewGenome(),
	}
}

type Genome struct {
	// Start *Part
	Parts []*Part
}

func getEndLength(left, right string) int {
	leftLen := len(left)
	rightLen := len(right)
	maxLength := leftLen
	if rightLen < maxLength {
		maxLength = rightLen
	}

	for i := maxLength; i >= minLengthMatching; i-- {
		if idx := strings.Index(left[:i], right[rightLen-i:]); idx != -1 {
			return i
		}
	}
	return -1
}

func (w *Wookie) Compute() {
	for _, line := range w.Lines {
		part := NewPart(line)
		w.Parts = append(w.Parts, &part)
	}

	//for _, line := range w.Lines {
	for _, part := range w.Parts {
		line := part.Line
		//part := NewPart(line)

		for _, existingPart := range w.Parts {
			if part == existingPart {
				continue
			}
			if endLength := getEndLength(line, existingPart.Line); endLength > -1 {
				part.Lefts[endLength] = append(part.Lefts[endLength], existingPart)
				existingPart.Rights[endLength] = append(existingPart.Rights[endLength], part)
			}
		}

		//w.Parts = append(w.Parts, &part)
	}

	w.Missings = len(w.Parts)

}

func NewGenome() Genome {
	return Genome{
		Parts: make([]*Part, 0),
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

func (g *Genome) String() string {
	output := g.Parts[1].Line
	for i := 1; i < len(g.Parts)-1; i++ {
		leftPart := g.Parts[i]
		rightPart := g.Parts[i+1]
		intersectionLength := getEndLength(rightPart.Line, leftPart.Line)
		output += rightPart.Line[intersectionLength:]
	}
	return output
}

func (w *Wookie) BuildGenomesRec(part *Part) bool {
	w.Genome.Parts = append(w.Genome.Parts, part)
	if w.Missings < 1 {
		return true
	}
	part.Done = true
	w.Missings--

	for _, rights := range part.Rights {
		for _, right := range rights {
			if right.Done {
				continue
			}
			if finished := w.BuildGenomesRec(right); finished {
				return true
			}
		}
	}

	w.Genome.Parts = w.Genome.Parts[:len(w.Genome.Parts)-1]

	w.Missings++
	part.Done = false
	return false
}

func (w *Wookie) GetGenome() *Genome {
	startPart := Part{
		Rights: map[int][]*Part{
			0: w.Parts,
		},
	}
	if finished := w.BuildGenomesRec(&startPart); !finished {
		return nil
	}
	return &(w.Genome)
}
