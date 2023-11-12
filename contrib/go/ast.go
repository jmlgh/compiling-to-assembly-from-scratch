package main

type AST interface {
	Equals(node AST) bool
}

type Number struct {
	Value uint
}

func (n *Number) Equals(other AST) bool {
	if _, ok := other.(*Number); !ok {
		return false
	}
	return n.Value == other.(*Number).Value
}

type Id struct {
	Value string
}

func (i *Id) Equals(other AST) bool {
	if _, ok := other.(*Id); !ok {
		return false
	}
	return i.Value == other.(*Id).Value
}

type Not struct {
	Term AST
}

func (n *Not) Equals(other AST) bool {
	if _, ok := other.(*Not); !ok {
		return false
	}
	return n.Term == other.(*Not).Term
}

type Equals struct {
	Left, Right AST
}

func (e *Equals) Equals(other AST) bool {
	if _, ok := other.(*Equals); !ok {
		return false
	}
	return e.Left == other.(*Equals).Left &&
		e.Right == other.(*Equals).Right
}

type NotEqual struct {
	Left, Right AST
}

func (ne *NotEqual) Equals(other AST) bool {
	if _, ok := other.(*NotEqual); !ok {
		return false
	}
	return ne.Left == other.(*NotEqual).Left &&
		ne.Right == other.(*NotEqual).Right
}

type Add struct {
	Left, Right AST
}

func (a *Add) Equals(other AST) bool {
	if _, ok := other.(*Add); !ok {
		return false
	}
	return a.Left == other.(*Add).Left &&
		a.Right == other.(*Add).Right
}

type Subtract struct {
	Left, Right AST
}

func (s *Subtract) Equals(other AST) bool {
	if _, ok := other.(*Add); !ok {
		return false
	}
	return s.Left == other.(*Add).Left &&
		s.Right == other.(*Add).Right
}

type Multiply struct {
	Left, Right AST
}

func (m *Multiply) Equals(other AST) bool {
	if _, ok := other.(*Add); !ok {
		return false
	}
	return m.Left == other.(*Add).Left &&
		m.Right == other.(*Add).Right
}

type Divide struct {
	Left, Right AST
}

func (d *Divide) Equals(other AST) bool {
	if _, ok := other.(*Add); !ok {
		return false
	}
	return d.Left == other.(*Add).Left &&
		d.Right == other.(*Add).Right
}
