package npc

import (
	"github.com/Hucaru/Valhalla/common/constants"
	"github.com/Hucaru/Valhalla/common/nx"
	"github.com/Hucaru/gopacket"
)

func SpawnNPC(index uint32, npc nx.Life) gopacket.Packet {
	p := gopacket.NewPacket()
	p.WriteByte(constants.SEND_CHANNEL_NPC_SPAWN_1)
	p.WriteUint32(index)
	p.WriteUint32(npc.ID)
	p.WriteInt16(npc.X)
	p.WriteInt16(npc.Y)

	p.WriteBool(npc.F)

	p.WriteInt16(npc.Fh)
	p.WriteInt16(npc.Rx0)
	p.WriteInt16(npc.Rx1)

	p.WriteByte(constants.SEND_CHANNEL_NPC_SPAWN_2)
	p.WriteByte(0x1)
	p.WriteUint32(npc.ID)
	p.WriteUint32(npc.ID)
	p.WriteInt16(npc.X)
	p.WriteInt16(npc.Y)

	p.WriteBool(npc.F)
	p.WriteBool(npc.F)

	p.WriteInt16(npc.Fh)
	p.WriteInt16(npc.Rx0)
	p.WriteInt16(npc.Rx1)

	return p
}

func ChangeController(index uint32, npc nx.Life) gopacket.Packet {
	p := gopacket.NewPacket()
	p.WriteByte(0x99)
	p.WriteByte(1)

	p.WriteUint32(index)
	p.WriteUint32(npc.ID)
	p.WriteInt16(npc.X)
	p.WriteInt16(npc.Y)

	p.WriteBool(npc.F)

	p.WriteInt16(npc.Fh)
	p.WriteInt16(npc.Rx0)
	p.WriteInt16(npc.Rx1)

	return p
}