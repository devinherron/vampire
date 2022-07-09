package game

import "testing"

func Test_AddMemory(t *testing.T) {
	var characters = []string{"character one", "character two"}
	var experiences1 = [3]string{"experience 1", "experience 2", "experience 3"}
	var experiences2 = [3]string{"different exp 1", "diff exp 2", "diff exp 3"}

	var memories [5]Memory
	var memory1 = Memory{experiences1}
	var memory2 = Memory{experiences2}
	memories[0] = memory1
	memories[1] = memory2
	var emptyDiary [4]Memory

	var resource1 = Resource{"resource 1", emptyDiary}
	var resources []Resource
	resources = append(resources, resource1)

	var marks = []string{"mark 1"}

	var newExperiences [3]string
	newExperiences[0] = "new experience"
	newMemory := Memory{newExperiences}

	var vampire = Vampire{"test", characters, memories, resources, marks}
	AddMemory(&vampire, newMemory)

	if vampire.Memories[2].Experiences[0] != newExperiences[0] {
		t.Errorf("New memory was not added")
	}
}

func Test_AddExperience(t *testing.T) {

}
