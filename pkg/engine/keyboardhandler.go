package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyBindingFunc func()

type KeyBinding struct {
	MainKey   ebiten.Key
	Modifiers []ebiten.Key
}

func (k *KeyBinding) Compare(keybinding *KeyBinding) bool {
	if keybinding == nil {
		return false
	}
	if k.MainKey != keybinding.MainKey {
		return false
	}
	for i, key := range k.Modifiers {
		if key != keybinding.Modifiers[i] {
			return false
		}
	}
	return true
}

type KeyboardHandler struct {
	*Base
	keybindings map[*KeyBinding][]KeyBindingFunc
}

func NewKeyboardHandler(name string) *KeyboardHandler {
	return &KeyboardHandler{
		Base:        NewBase(name),
		keybindings: make(map[*KeyBinding][]KeyBindingFunc),
	}
}

func (k *KeyboardHandler) AddKeyBindingForKey(mainkey ebiten.Key, modifiers []ebiten.Key, f KeyBindingFunc) {
	keybinding := &KeyBinding{
		MainKey:   mainkey,
		Modifiers: modifiers,
	}
	k.keybindings[keybinding] = append(k.keybindings[keybinding], f)
}

func (k *KeyboardHandler) GetKeyBindingsForKey(mainkey ebiten.Key, modifiers []ebiten.Key) []KeyBindingFunc {
	keybinding := &KeyBinding{
		MainKey:   mainkey,
		Modifiers: modifiers,
	}
	for trav, bindings := range k.keybindings {
		if trav.Compare(keybinding) {
			return bindings
		}
	}
	return nil
}

func isBindingPressed(binding *KeyBinding) bool {
	if !inpututil.IsKeyJustPressed(binding.MainKey) {
		return false
	}
	for _, modifier := range binding.Modifiers {
		if !inpututil.IsKeyJustPressed(modifier) {
			return false
		}
	}
	return true
}

func (k *KeyboardHandler) Update(args ...any) {
	for keybinding, bindings := range k.keybindings {
		if isBindingPressed(keybinding) {
			for _, f := range bindings {
				f()
			}
		}
	}
}
