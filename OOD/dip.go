package OOD

// 객체에 의존하게되면 계속해서 기능이 늘어나야 한다.

type Player struct {
}

type Monster struct {
}

type Chest struct {
}

// func (p *Player) AttackPM(m *Monster) {

// }

// func (p *Player) AttackPP(p2 *Player) {

// }

// func (m *Monster) AttackMP(p *Player) {

// }

// func (m *Monster) AttackMM(m2 *Monster) {

// }

type Attackble interface {
	Attack()
}

type BeAttackalbe interface {
	BeAttacked()
}

func (p *Player) Attack(target *BeAttackalbe) {
	// ...
	// 데미지를 계산해서
	target.BeAttacked(dmg)
}

func Attack(attacker *Attackable, defender *BeAttackalbe) {
	attacker.Attack(defender)
}
