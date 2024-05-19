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
	defer m.mutex.Unlock()
	m.members = []*discordgo.Member{}
}

func (m *MembersState) Append(member discordgo.Member) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.members = append(m.members, &member)
}

func (m *MembersState) Member(index int) discordgo.Member {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return *m.members[index]
}

func (m *MembersState) Update(index int, member discordgo.Member) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.members[index] = &member
}

func (m *MembersState) Length() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return len(m.members)
}

func (m *MembersState) Delete(index int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.members = slices.Delete(m.members, index, index+1)
}

func (m *MembersState) Exists(memberId string) (bool, int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, member := range m.members {
		if member.User.ID == memberId {
			return true, i
		}
	}

	return false, -1
}
