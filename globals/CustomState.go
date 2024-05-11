package globals

import (
	"slices"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type MembersState struct {
	mutex   sync.Mutex
	members []*discordgo.Member
}

func (m *MembersState) Clear() {
	m.mutex.Lock()
	m.members = []*discordgo.Member{}
	m.mutex.Unlock()
}

func (m *MembersState) Append(member *discordgo.Member) {
	m.mutex.Lock()
	m.members = append(m.members, member)
	m.mutex.Unlock()
}

func (m *MembersState) Members() []*discordgo.Member {
	return m.members
}

func (m *MembersState) Length() int {
	return len(m.members)
}

func (m *MembersState) Delete(index int) {
	m.mutex.Lock()
	m.members = slices.Delete(m.members, index, index+1)
	m.mutex.Unlock()
}
