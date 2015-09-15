package mysqlproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPacket7(t *testing.T) {
	buf := newBuffer([]byte{
		0x07, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x01, 0x02, 0x03,
	})

	stream := NewStream(buf)
	packet, err := stream.NextPacket()
	assert.Nil(t, err)
	assert.Equal(t, packet.SequenceID, byte(0x02))
	assert.Len(t, packet.Payload, 7)
	assert.Equal(t, packet.Payload, []byte{0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03})
}

func TestNextPacket256(t *testing.T) {
	buf := newBuffer([]byte{
		0x00, 0x01, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02,
	})

	stream := NewStream(buf)
	packet, err := stream.NextPacket()
	assert.Nil(t, err)
	assert.Equal(t, packet.SequenceID, byte(0x02))
	assert.Len(t, packet.Payload, 256)
}
