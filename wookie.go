package wookie

import "strings"

const minLengthMatching = 3

type Part struct {
	Line   string
	Rights []*Part
	Done   bool
}

type Wookie struct {
	Lines    []string
	Parts    []*Part
	Genome   Genome
	Missings int
}

type Genome struct {
	// Start *Part
	Parts []*Part
}

func NewPart(line string) Part {
	return Part{
		Line:   line,
		Rights: make([]*Part, 0),
	}
}

func NewWookie(lines []string) Wookie {
	return Wookie{
		Lines:  lines,
		Parts:  make([]*Part, 0),
		Genome: NewGenome(),
	}
}

func NewGenome() Genome {
	return Genome{
		Parts: make([]*Part, 0),
	}
}

func (p *Part) CoincideLength(part *Part) int {
	rightLen := len(part.Line)

	maxLength := len(p.Line)
	if rightLen < maxLength {
		maxLength = rightLen
	}

	for i := maxLength; i >= minLengthMatching; i-- {
		if idx := strings.Index(p.Line[:i], part.Line[rightLen-i:]); idx != -1 {
			return i
		}
	}
	return -1
}

func (p *Part) CoincideWith(part *Part) bool {
	return p.CoincideLength(part) > 0
}

func (w *Wookie) Compute() {
	// prepare structures
	for _, line := range w.Lines {
		part := NewPart(line)
		w.Parts = append(w.Parts, &part)
	}
	w.Missings = len(w.Parts)

	for _, part := range w.Parts {
		for _, existingPart := range w.Parts {
			if part == existingPart {
				continue
			}
			if existingPart.CoincideWith(part) {
				part.Rights = append(part.Rights, existingPart)
			}
		}
	}

	// run recursive loop
	startPart := Part{
		Rights: w.Parts,
	}
	w.buildGenomeRec(&startPart)
}

func (g *Genome) String() string {
	output := g.Parts[1].Line
	for i := 1; i < len(g.Parts)-1; i++ {
		left := g.Parts[i]
		right := g.Parts[i+1]
		intersectionLength := right.CoincideLength(left)
		output += right.Line[intersectionLength:]
	}
	return output
}

func (w *Wookie) buildGenomeRec(part *Part) bool {
	w.Genome.Parts = append(w.Genome.Parts, part)
	if w.Missings < 1 {
		return true
	}
	part.Done = true
	w.Missings--

	for _, right := range part.Rights {
		if right.Done {
			continue
		}
		if finished := w.buildGenomeRec(right); finished {
			return true
		}
	}

	w.Genome.Parts = w.Genome.Parts[:len(w.Genome.Parts)-1]

	w.Missings++
	part.Done = false
	return false
}
