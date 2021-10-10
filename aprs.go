// Package aprs provides an Amateur Packet Radio Service messaging interface.
package aprs

import (
	"bytes"
	"fmt"
	"strings"
)

// Info represents the information payload of an APRS packet.
type Info string

// Frame represents a complete, abstract, APRS frame.
type Frame struct {
	Original string
	Source   Address
	Dest     Address
	Path     []Address
	Body     Info
}

// Type of the message.
func (i Info) Type() PacketType {
	t := PacketType(0)
	if len(i) > 0 {
		t = PacketType(i[0])
	}
	return t
}

// ParseFrame parses an APRS string into an Frame struct.
func ParseFrame(raw string) (Frame, error) {
	parts := strings.SplitN(raw, ":", 2)

	if len(parts) != 2 {
		return Frame{}, fmt.Errorf("Missing the header delimiter ':' in the string %q", raw)
	}
	srcparts := strings.SplitN(parts[0], ">", 2)
	if len(srcparts) < 2 {
		return Frame{}, fmt.Errorf("Missing the path delimiter '>' in the header: %q", parts[0])
	}
	pathparts := strings.Split(srcparts[1], ",")

	return Frame{Original: raw,
		Source: AddressFromString(srcparts[0]),
		Dest:   AddressFromString(pathparts[0]),
		Path:   parseAddresses(pathparts[1:]),
		Body:   Info(parts[1])}, nil
}

// Wire forms an Frame back into its proper wire format.
func (f Frame) Wire() string {
	b := bytes.NewBufferString(f.Source.String())
	b.WriteByte('>')
	b.WriteString(f.Dest.String())
	for _, p := range f.Path {
		b.WriteByte(',')
		b.WriteString(p.String())
	}
	b.WriteByte(':')
	b.WriteString(string(f.Body))
	return b.String()
}

func (f Frame) String() string {
	return f.Wire()
}
