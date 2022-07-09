package game

import (
	"log"
)

type Vampire struct {
	Name       string
	Characters []string
	Memories   [5]Memory
	Resources  []Resource
	Marks      []string
}

type Memory struct {
	Experiences [3]string
}

type Resource struct {
	Description string
	Diary       [4]Memory
}

func AddMemory(vampire *Vampire, memory Memory) {
	for i := 0; i < 5; i++ {
		if vampire.Memories[i].Experiences[0] == "" {
			vampire.Memories[i] = memory
			return
		}
	}

	log.Fatal("Tried to add a memory, but memories already full")
}

func AddExperience(vampire *Vampire, memory int, experience string) {
	for i := 0; i < 3; i++ {
		if vampire.Memories[memory].Experiences[i] == "" {
			vampire.Memories[memory].Experiences[i] = experience
			return
		}
	}
	log.Fatal("Tried to add an experience, but the memory was already full")
}

func AddMark(vampire *Vampire, mark string) {
	vampire.Marks = append(vampire.Marks, mark)
}

func AddResource(vampire *Vampire, resource Resource) {
	vampire.Resources = append(vampire.Resources, resource)
}
