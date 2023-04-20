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
	res.SetEffect(Resources.NewEffect())
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
			return strings.Join(os, "\n")
		}(item.Lines))
		title := Title.NewTitle(item.String(), item.StartAt.Seconds(), item.EndAt.Seconds()).SetTextStyleDef(textStyleDef).SetText(text)
		title.AddParam(Title.NewParams("Position", "9999/999166631/999166633/1/100/101", "0 -450"))
		title.AddParam(Title.NewParams("Alignment", "9999/999166631/999166633/2/354/999169573/401", "1 (Center)"))
		title.AddParam(Title.NewParams("Flatten", "9999/999166631/999166633/2/351", "1"))
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
