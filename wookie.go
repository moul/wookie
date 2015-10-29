package wookie

import "strings"

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

func (p *Part) CoincideLength(part *Part, minLength int) int {
	rightLen := len(part.Line)

	maxLength := len(p.Line)
	if rightLen < maxLength {
		maxLength = rightLen
	}

	for i := maxLength; i >= minLength; i-- {
		if idx := strings.Index(p.Line[:i], part.Line[rightLen-i:]); idx != -1 {
			return i
		}
	}
	return -1
}

func (p *Part) CoincideWith(part *Part, minLength int) bool {
	return p.CoincideLength(part, minLength) > -1
}

func (w *Wookie) Compute() bool {

	minLength := len(w.Lines[0])

	// prepare structures
	for _, line := range w.Lines {
		part := NewPart(line)
		if len(line) < minLength {
			minLength = len(line)
		}
		w.Parts = append(w.Parts, &part)
	}

	startPart := Part{
		Rights: w.Parts,
	}

	for length := minLength; length > 0; length-- {
		w.Missings = len(w.Parts)
		for _, part := range w.Parts {
			//part.Rights = part.Rights[:]
			part.Rights = make([]*Part, 0)
			part.Done = false
		}
		for _, part := range w.Parts {
			for _, otherPart := range w.Parts {
				if part == otherPart {
					continue
				}
				if otherPart.CoincideWith(part, length) {
					part.Rights = append(part.Rights, otherPart)
				}
			}
		}

		// run recursive loop
		if ret := w.buildGenomeRec(&startPart); ret {
			return true
		}
	}
	return false
}

func (g *Genome) String() string {
	output := g.Parts[1].Line
	for i := 1; i < len(g.Parts)-1; i++ {
		left := g.Parts[i]
		right := g.Parts[i+1]
		intersectionLength := right.CoincideLength(left, 0)
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
