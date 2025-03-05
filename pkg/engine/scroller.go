package engine

import "fmt"

// iterScroller structure defines the data required to iterate over a Scroller
// instance.
type iterScroller struct {
	index  int
	offset int
	delta  int
}

// newItemScroller function creates a new iterScroller instance.
func newIterScroller(index int, delta int) *iterScroller {
	iter := &iterScroller{
		index:  index,
		offset: index * delta,
		delta:  delta,
	}
	return iter
}

// GetIndex method returns the iter scroller index.
func (i *iterScroller) GetIndex() int {
	return i.index
}

// GetOffset method returns the iter scroller offset.
func (i *iterScroller) GetOffset() int {
	return i.offset
}

// Next method increments the iter scroller.
func (i *iterScroller) Next() {
	i.index++
	i.offset = i.index * i.delta
}

// SetIndex method sets index attribute to the given value.
func (i *iterScroller) SetIndex(index int) {
	i.index = index
}

func (i *iterScroller) ToString() string {
	return fmt.Sprintf("%d:%d:%d", i.index, i.offset, i.delta)
}

// Scroller structure defines the baseline for any widget using a scroller.
// It can be an horizontal or vertical scroller.
// Vertical scroller sets SelectionLength as always ONE, because it is based on
// lines, where MaxLength is the number of lines that can be displayed at one
// time and TotalSelecionLength is the total number of selections.
// StartSelection and EndSelection represent the index values for the first and
// the last selection to be displayed.
// Horizontal scroller in other way is based in the number of characters and
// not in the number of lines. SelectionLength is the length every selection
// should take, where MaxLength is the maximum number of characters for all
// selections to be displayed at any time, where TotalSelectionLength is the
// total number of characters for all possible selections. StartSelection and
// EndSelection represent the number of characters to shift for the first
// selection and the last selection to be displayed.
type Scroller struct {
	TotalSelectionLength int
	MaxLength            int
	SelectionLength      int
	StartSelection       int
	EndSelection         int
	sIter                *iterScroller
}

// NewScroller function creates a new Scroller instance.
func NewScroller(totalSelectionLength int, maxLength int, selectionLength int) *Scroller {
	scroller := &Scroller{
		TotalSelectionLength: totalSelectionLength,
		MaxLength:            maxLength,
		SelectionLength:      selectionLength,
	}
	if totalSelectionLength > maxLength {
		scroller.StartSelection = 0
		scroller.EndSelection = (maxLength / selectionLength) - 1
	} else {
		scroller.StartSelection = 0
		scroller.EndSelection = (totalSelectionLength / selectionLength) - 1
	}
	fmt.Printf("%#+v\n", scroller)
	return scroller
}

// NewVerticalScroller function creates a new vertival Scroller instance, where
// the selection Length is fixed to one.
func NewVerticalScroller(totalNbrOfSelections int, nbrOfSelectionToDisplay int) *Scroller {
	return NewScroller(totalNbrOfSelections, nbrOfSelectionToDisplay, 1)
}

// -----------------------------------------------------------------------------
// Scroller Public methods
// -----------------------------------------------------------------------------

// CreateIter method created a new scroller iterator.
func (s *Scroller) CreateIter() *iterScroller {
	s.sIter = newIterScroller(s.StartSelection, s.SelectionLength)
	return s.sIter
}

// IterHasNext method checks if there are still some entries in the scroller
// iterator.
func (s *Scroller) IterHasNext() bool {
	return s.sIter.GetIndex() <= s.EndSelection
}

// IterGetNext method returns the next entry to interate in the scorller and
// increase all iterator attributes.
func (s *Scroller) IterGetNext() (int, int) {
	index := s.sIter.GetIndex()
	offset := s.sIter.GetOffset()
	s.sIter.Next()
	return index, offset
}

func (s *Scroller) ToString() string {
	return fmt.Sprintf("%d:%d:%d %d:%d", s.TotalSelectionLength, s.MaxLength, s.SelectionLength, s.StartSelection, s.EndSelection)
}

// Update method update the scroller StartSelection and EndSelection based on
// the given selection value.
func (s *Scroller) Update(selectionIndex int) {
	if selectionIndex < s.StartSelection {
		s.StartSelection = selectionIndex
		s.EndSelection = s.StartSelection + (s.MaxLength / s.SelectionLength) - 1
	} else if selectionIndex > s.EndSelection {
		s.EndSelection = selectionIndex
		s.StartSelection = s.EndSelection - (s.MaxLength / s.SelectionLength) + 1
	}
}
