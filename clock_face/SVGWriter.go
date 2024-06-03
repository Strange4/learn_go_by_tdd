package clock_face

import (
	"fmt"
	"io"
)

func SVGWriter(writer io.Writer, clock *Clock) {
	secondHand := clock.SecondHand()

	seconds := fmt.Sprintf(`
	<!-- second hand -->
	<line id="second_hand" x1="150" y1="150" x2="%.3f" y2="%.3f"
		  style="fill:none;stroke:#f00;stroke-width:3px;"/>`, secondHand.X, secondHand.Y)

	writer.Write([]byte(svgStart))
	writer.Write([]byte(seconds))
	writer.Write([]byte(svgEnd))
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

	<!-- bezel -->
	<circle id="bezel" cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

	<!-- hour hand -->
	<line id="hour_hand" x1="150" y1="150" x2="114.150000" y2="132.260000"
		style="fill:none;stroke:#000;stroke-width:7px;"/>

	<!-- minute hand -->
	<line id="minute_hand" x1="150" y1="150" x2="101.290000" y2="99.730000"
		style="fill:none;stroke:#000;stroke-width:7px;"/>
`

const svgEnd = `</svg>`
