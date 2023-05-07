package Resources

type Effect struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Uid  string `xml:"uid,attr"`
}

func NewEffect() *Effect {
	return &Effect{
		Text: "",
		ID:   "r2",
		Name: "Basic Title",
		Uid:  ".../Titles.localized/Bumper:Opener.localized/Basic Title.localized/Basic Title.moti",
	}
}

func NewBackGroundEffect() *Effect {
	return &Effect{
		Text: "",
		ID:   "r2",
		Name: "自适应背景",
		Uid:  "~/Titles.localized/自适应背景/自适应背景.moti",
	}
}
