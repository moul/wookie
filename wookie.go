package wookie

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/awalterschulze/gographviz"
)

// minimal length of a sequence
const MinSequenceLength = 2

type Part struct {
	Line   string
	Rights map[int][]*Part
	Done   bool
}

type Wookie struct {
	Sequences []string
	Parts     []*Part
	Genome    Genome
	Missings  int
	Length    int
}

type Genome struct {
	// Start *Part
	Parts []*Part
}

func NewPart(line string) Part {
	return Part{
		Line:   line,
		Rights: make(map[int][]*Part, 0),
	}
}

func NewWookie(input string) Wookie {
	sequences := regexp.MustCompile("\\w+").FindAllString(input, -1)
	return Wookie{
		Sequences: sequences,
		Parts:     make([]*Part, 0),
		Genome:    NewGenome(),
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

	for i := maxLength; i >= MinSequenceLength; i-- {
		if idx := strings.Index(p.Line[:i], part.Line[rightLen-i:]); idx != -1 {
			return i
		}
	}
	return -1
}

func (w *Wookie) Compute() error {
	minLength := len(w.Sequences[0])

	// prepare structures
	for _, line := range w.Sequences {
		part := NewPart(line)
		if len(line) < minLength {
			minLength = len(line)
		}
		w.Parts = append(w.Parts, &part)
	}

	for _, part := range w.Parts {
		for _, otherPart := range w.Parts {
			if part == otherPart {
				continue
			}
			coincideLength := otherPart.CoincideLength(part)
			if coincideLength > 0 {
				part.Rights[coincideLength] = append(part.Rights[coincideLength], otherPart)
			}
		}
	}

	for w.Length = minLength; w.Length >= MinSequenceLength; w.Length-- {
		// reset parts
		startPart := Part{
			Rights: map[int][]*Part{
				w.Length: w.Parts,
			},
		}
		w.Missings = len(w.Parts)
		for _, part := range w.Parts {
			part.Done = false
		}

		// run recursive loop
		if ret := w.buildGenomeRec(&startPart, w.Length); ret {
			return nil
		}
	}
	return fmt.Errorf("No such genome with these sequences")
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

func (w *Wookie) Graphviz() string {
	graph := gographviz.NewGraph()

	graph.SetName("G")
	graph.SetDir(true)

	graph.AddNode("G", "start", map[string]string{"shape": "Mdiamond"})
	graph.AddNode("G", "end", map[string]string{"shape": "Msquare"})
	graph.AddEdge("start", w.Genome.Parts[1].Line, true, nil)
	graph.AddEdge(w.Genome.Parts[len(w.Genome.Parts)-1].Line, "end", true, nil)

	for _, part := range w.Genome.Parts {
		if part.Line == "" {
			continue
		}
		attrs := map[string]string{"color": "blue"}
		graph.AddNode("G", part.Line, attrs)
	}

	for i := 1; i < len(w.Genome.Parts)-1; i++ {
		left := w.Genome.Parts[i]
		right := w.Genome.Parts[i+1]
		attrs := map[string]string{
			"color": "red",
			"label": fmt.Sprintf("idx%d_len%d", i, right.CoincideLength(left)),
		}
		graph.AddEdge(left.Line, right.Line, true, attrs)
	}

	for _, part := range w.Parts {
		for length, rights := range part.Rights {
			if length < w.Length {
				continue
			}

			attrs := map[string]string{
				"label": fmt.Sprintf("len%d", length),
				//"iterations": "nslimit1",
				//"splines":    "line",
				//"maxiter":    "1",
				//"mclimit":    "1",
			}
			for _, right := range rights {
				isSolution := false
				for idx, solutionPart := range w.Genome.Parts {
					if solutionPart == part {
						if w.Genome.Parts[idx+1] == right {
							isSolution = true
						}
						break
					}
				}
				if isSolution {
					continue
				}

				graph.AddEdge(part.Line, right.Line, true, attrs)
			}
		}
	}
	return graph.String()
}

func (w *Wookie) buildGenomeRec(part *Part, minLength int) bool {
	w.Genome.Parts = append(w.Genome.Parts, part)
	if w.Missings == 0 { // genome is complete
		return true
	}
	part.Done = true
	w.Missings--

	for length, rights := range part.Rights {
		if length < minLength {
			continue
		}
		for _, right := range rights {
			if right.Done {
				continue
			}
			if finished := w.buildGenomeRec(right, minLength); finished {
				return true
			}
		}
	}

	w.Genome.Parts = w.Genome.Parts[:len(w.Genome.Parts)-1]

	w.Missings++
	part.Done = false
	return false
}
