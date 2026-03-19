// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package validators

import "testing"

func TestMockAccessor_GetString_Present(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{"key": "value"})
	v, ok := m.GetString("key")
	if !ok || v != "value" {
		t.Errorf("GetString(key) = (%q, %v), want (\"value\", true)", v, ok)
	}
}

func TestMockAccessor_GetString_Missing(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{})
	v, ok := m.GetString("missing")
	if ok || v != "" {
		t.Errorf("GetString(missing) = (%q, %v), want (\"\", false)", v, ok)
	}
}

func TestMockAccessor_GetBool_Present(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{"enabled": true})
	v, ok := m.GetBool("enabled")
	if !ok || !v {
		t.Errorf("GetBool(enabled) = (%v, %v), want (true, true)", v, ok)
	}
}

func TestMockAccessor_GetBool_Missing(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{})
	v, ok := m.GetBool("missing")
	if ok || v {
		t.Errorf("GetBool(missing) = (%v, %v), want (false, false)", v, ok)
	}
}

func TestMockAccessor_GetStringList_Present(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{"items": []string{"a", "b"}})
	v, ok := m.GetStringList("items")
	if !ok || len(v) != 2 || v[0] != "a" || v[1] != "b" {
		t.Errorf("GetStringList(items) = (%v, %v), want ([a b], true)", v, ok)
	}
}

func TestMockAccessor_GetStringList_Missing(t *testing.T) {
	m := NewMockAccessor(map[string]interface{}{})
	v, ok := m.GetStringList("missing")
	if ok || v != nil {
		t.Errorf("GetStringList(missing) = (%v, %v), want (nil, false)", v, ok)
	}
}
