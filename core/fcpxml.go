package core

import (
	"encoding/xml"
	"strings"

	"github.com/asticode/go-astisub"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event/Project"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event/Project/Sequence"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event/Project/Sequence/Spine"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event/Project/Sequence/Spine/Gap"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Library/Event/Project/Sequence/Spine/Gap/Title"
	"github.com/hnlq715/srt2fcpxml/core/FcpXML/Resources"
)

func Srt2FcpXmlExport(projectName string, frameDuration interface{}, subtitles *astisub.Subtitles, width, height int) ([]byte, error) {
	fcpxml := FcpXML.New()
	res := Resources.NewResources()
	res.SetEffect(Resources.NewBackGroundEffect())
	format := Resources.NewFormat().
		SetWidth(width).
		SetHeight(height).
		SetFrameRate(frameDuration).Render()
	res.SetFormat(format)
	fcpxml.SetResources(res)
	gap := Gap.NewGap(subtitles.Duration().Seconds())

	for index, item := range subtitles.Items {
		textStyleDef := Title.NewTextStyleDef(index + 1)
		text := Title.NewContent(index+1, func(lines []astisub.Line) string {
			var os []string
			for _, l := range lines {
				os = append(os, l.String())
			}
			return strings.Trim(strings.Join(os, "\n"), "\n")
		}(item.Lines))
		title := Title.NewTitle(item.String(), item.StartAt.Seconds(), item.EndAt.Seconds()).SetTextStyleDef(textStyleDef).SetText(text)
		title.AddParam(Title.NewParams("Position", "9999/999166631/999166633/1/100/101", "0 -450"))
		title.AddParam(Title.NewParams("Alignment", "9999/999166631/999166633/2/354/999169573/401", "1 (Center)"))
		title.AddParam(Title.NewParams("Flatten", "9999/999166631/999166633/2/351", "1"))
		title.AddParam(Title.NewParams("Build In", "9999/10000/2/101", "0"))
		title.AddParam(Title.NewParams("Build Out", "9999/10000/2/102", "0"))
		title.AddParam(Title.NewParams("ScaleY", "9999/1825768479/100/1825768480/2/100", "0"))
		title.AddParam(Title.NewParams("Opacity", "9999/1825768325/10003/10045/1/200/202", "0.7"))
		gap.AddTitle(title)
	}

	spine := Spine.NewSpine().SetGap(gap)
	seq := Sequence.NewSequence(subtitles.Duration().Seconds()).SetSpine(spine)
	project := Project.NewProject(projectName).SetSequence(seq)
	event := Event.NewEvent().SetProject(project)
	library := Library.NewLibrary(projectName).SetEvent(event)
	fcpxml.SetLibrary(library)

	return xml.MarshalIndent(fcpxml, "", "    ")
}
