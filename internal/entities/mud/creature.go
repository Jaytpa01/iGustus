package mud

type Creature struct {
	Species   string
	BodyParts map[string]Object
}

type Humanoid struct {
	Creature
	Race string
}

const (
	TORSO = "torso"
	HEAD  = "head"
)

func NewHumanoid() *Humanoid {
	return &Humanoid{
		Creature: Creature{
			Species:   "Human",
			BodyParts: HumanoidBodyParts(),
		},
		Race: "Nord",
	}
}

func HumanoidBodyParts() map[string]Object {
	m := map[string]Object{
		TORSO: NewContainer(),
		HEAD:  NewContainer(),
	}

	return m
}
