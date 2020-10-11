package app

type PeepData struct {
	Filename   string
	Sheet      string
	PeepValues []*PeepValue
}

func NewPeepData(filename, sheet string) *PeepData {
	return &PeepData{filename, sheet, []*PeepValue{}}
}

func (p *PeepData) AddPeepValue(v *PeepValue) {
	p.PeepValues = append(p.PeepValues, v)
}

type PeepValue struct {
	Cell  string
	Value string
}

func NewPeepValue(cell, value string) *PeepValue {
	return &PeepValue{cell, value}
}

func (p *PeepData) GetTableHeader() []string {
	header := []string{"FILE", "SHEET"}
	for _, d := range p.PeepValues {
		header = append(header, d.Cell)
	}
	return header
}

func (p *PeepData) GetTableSrc() []string {
	src := []string{p.Filename, p.Sheet}
	for _, d := range p.PeepValues {
		src = append(src, d.Value)
	}
	return src
}
