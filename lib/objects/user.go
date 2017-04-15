package objects

import "strconv"

//User structure for the user
type User struct {
	X, Y, Z        int64
	C              chan string
	Name           string
	SentTimeUpdate int
}

//ToString Function to convert the user to a string object
func (u *User) ToString() string {
	return strconv.FormatInt(u.X, 10) + ", " + strconv.FormatInt(u.Y, 10) + ", " + strconv.FormatInt(u.Z, 10)
}

//Move Function to move the user's position
func (u *User) Move(direction int) {
	switch direction {
	case 0:
		u.Y = u.Y + 1
	case 1:
		u.X = u.X + 1
	case 2:
		u.X = u.X - 1
	case 3:
		u.Y = u.Y - 1
	case 4:
		u.Z = u.Z + 1
	case 5:
		u.Z = u.Z - 1
	}
}

/*
 Strength - A measure of the characters ability to exert physical
  force. The attribute is also referred to as Physical Strength. The
  Strength attribute factors into the ability to lift and carry heavy
  loads; the amount of damage inflicted through a blow; and restricts
  the types of manual tools that can be operated easily.

  Dexterity - The relative ability to react physically to a brief
  event. Dexterity is also called Agility, Reflexes, or Physical
  Prowess. It differs from pure exertion of strength, and requires tight
  coordination between the central nervous system and the
  musculature. This attribute is used to evaluate most competitive
  physical skills, such as striking a target with a missile, or dodging
  a blow.


  Fatigue - The short term physical energy that the character can expend
  before becoming exhausted. It is also called Energy. This attribute
  places a more realistic cap on the characters activities during a
  battle, and is also used to limit magical powers.

  Endurance - A rating of how much cumulative damage a creature can
  withstand before dying. Endurance is also referred to as Hit Points or
  Body. In systems that do not include defensive skills, Hit Points are
  also used to measure the amount of combat experience. This mixture of
  unrelated physiological and combat capabilities within a single
  parameter can have a distorting effect, resulting in an unrealistic
  model. However, it does have the side benefits of making combat less
  bloody and more protracted at higher levels, as the characters are
  rarely eliminated by a single blow.

  Appearance - The physical attractiveness of the character,
  particularly with regards to the opposite sex. This attribute has also
  been called Comeliness and Beauty. The effect on game play is less
  useful than most physical attributes, being required primarily during
  social situations. However, physical beauty has been known to enjoy
  subtle effects on the human psyche, inspiring loyalty and trust beyond
  the norm.

  Intelligence - The mental ability to remember facts and employ
  reasoning to analyze a problem. Known also as I.Q., Memory, and
  Reasoning. Intelligence is also a crude rating of the ability to use
  mathematics, solve puzzles, learn a new language, create music, etc.

  Quickness - The speed of reaction to a rapidly changing
  situation. Quickness is also known as Reaction or Initiative. It
  determines who reacts first during a crisis, allowing a quick witted
  character to get the jump on a foe. Quickness is closely related to
  the Dexterity attribute.

  Size - The gross physical proportions of a creature. Also known as
  Height or Growth. This attribute is often used to handle scaling
  issues, such as visibility, melee reach, minimum opening required to
  enter a room, etc.

  Cool - This measures the ability of the character to remain calm under
  conditions of duress, such as during combat or when facing a truly
  terrifying sight. It is also called Morale, and is a measure of a
  creatures steady reaction to a panic situation. In AD&D, this
  Attribute is only used for creatures, and the morale aspects of the
  characters are handled directly by the players.

  Wisdom - An Attribute that measures a characters worldly knowledge and
  common sense. Also known as empathy. It is used, probably incorrectly,
  as a magical attribute for priestly magic. Most modern systems use
  willpower or piety, and leave knowledge evaluation to a skills system,
  various advantages, and/or role-playing.

  Cunning - A primal form of intelligence available to many animals,
  cunning measures a creatures ability to quickly formulate good
  reactions to stressful conditions. For humans, this measures wits, and
  is closely related (if not identical) to quickness. Cunning is useful
  for evaluating a changing tactical situation, coming up with a witty
  rejoinder during a discussion, taking advantage of a lucky break, and
  so on.

  Willpower - The amount of self-control a character has over his own
  mind and body. This attribute is also called Mental Endurance, Self
  Discipline, and Ego. It is often used as a measure of a characters
  control over arcane forces, the ability to resist the imposition of
  another's will, and the degree of vulnerability to fearful thoughts
  and experiences. Willpower is closely related to cool.

  Leadership - The ability to influence the behavior of others using a
  commanding presence, persuasive dialogue, and appealing behavior. This
  attribute has also called Charisma, Mental Affinity, Power, or
  Presence. It does not imply an ability to lead sensibly, but does
  enhance the loyalty and morale of friends and allies under
  discouraging conditions. Some systems use a separate set of leadership
  skills.

  Fellowship - The ease with which a character associates with others in
  a social environment. Many systems use a skill to handle this ability,
  although some people do seem to have an innate ability to get along
  well with others.

  Movement - This derived attribute is used to determine how far a
  creature can move during an interval of time. Movement is also called
  Speed. Typically the movement rate is a fixed value for each race,
  with a modifier based on the Dexterity Attribute. Separate factors are
  also used for measuring different means of movement, such a swimming,
  flying, tunneling, etc.



*/
