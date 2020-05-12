package OOD

type Actor interface {
	Move()
	Attack()
	Talk()
}

// type Actor struct {
// }

type Movable interface {
	Move()
}

type Attackable interface {
	Attack()
}

type Talkable interface {
	Talk()
}

// srp를 어긴다.
func MoveTo(a Actor) {
	a.Move()
	a.Attack()
}

// srp를 강제한다
func MoveTo(a Movable) {
	a.Move()
	a.Attack()
}
