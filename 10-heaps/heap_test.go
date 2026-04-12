package heaps

import "testing"

func TestMinHeap_InsertAndExtract(t *testing.T) {
	h := NewMinHeap[string]()
	h.Insert(Entry[string]{Priority: 5, Value: "five"})
	h.Insert(Entry[string]{Priority: 3, Value: "three"})
	h.Insert(Entry[string]{Priority: 7, Value: "seven"})
	h.Insert(Entry[string]{Priority: 1, Value: "one"})

	// Should extract in ascending priority order
	expected := []string{"one", "three", "five", "seven"}
	for _, want := range expected {
		got, ok := h.Extract()
		if !ok {
			t.Fatalf("Extract() returned ok=false, want %s", want)
		}
		if got != want {
			t.Errorf("Extract() = %s, want %s", got, want)
		}
	}
}

func TestMaxHeap_InsertAndExtract(t *testing.T) {
	h := NewMaxHeap[int]()
	h.Insert(Entry[int]{Priority: 5, Value: 5})
	h.Insert(Entry[int]{Priority: 3, Value: 3})
	h.Insert(Entry[int]{Priority: 7, Value: 7})
	h.Insert(Entry[int]{Priority: 1, Value: 1})

	// Should extract in descending priority order
	expected := []int{7, 5, 3, 1}
	for _, want := range expected {
		got, ok := h.Extract()
		if !ok {
			t.Fatalf("Extract() returned ok=false, want %d", want)
		}
		if got != want {
			t.Errorf("Extract() = %d, want %d", got, want)
		}
	}
}

func TestHeap_Peek(t *testing.T) {
	h := NewMinHeap[int]()
	_, ok := h.Peek()
	if ok {
		t.Error("Peek() on empty heap should return ok=false")
	}

	h.Insert(Entry[int]{Priority: 10, Value: 10})
	h.Insert(Entry[int]{Priority: 5, Value: 5})
	val, ok := h.Peek()
	if !ok {
		t.Fatal("Peek() returned ok=false on non-empty heap")
	}
	if val != 5 {
		t.Errorf("Peek() = %d, want 5", val)
	}

	// Peek shouldn't remove the element
	if h.Size() != 2 {
		t.Errorf("Size() after Peek = %d, want 2", h.Size())
	}
}

func TestHeap_ExtractEmpty(t *testing.T) {
	h := NewMinHeap[int]()
	_, ok := h.Extract()
	if ok {
		t.Error("Extract() on empty heap should return ok=false")
	}
}

func TestHeap_Size(t *testing.T) {
	h := NewMaxHeap[int]()
	if h.Size() != 0 {
		t.Errorf("Size() on new heap = %d, want 0", h.Size())
	}
	h.Insert(Entry[int]{Priority: 1, Value: 1})
	h.Insert(Entry[int]{Priority: 2, Value: 2})
	if h.Size() != 2 {
		t.Errorf("Size() after 2 inserts = %d, want 2", h.Size())
	}
	h.Extract()
	if h.Size() != 1 {
		t.Errorf("Size() after extract = %d, want 1", h.Size())
	}
}

func TestMinHeap_Duplicates(t *testing.T) {
	h := NewMinHeap[int]()
	h.Insert(Entry[int]{Priority: 3, Value: 3})
	h.Insert(Entry[int]{Priority: 3, Value: 3})
	h.Insert(Entry[int]{Priority: 1, Value: 1})
	h.Insert(Entry[int]{Priority: 3, Value: 3})

	expected := []int{1, 3, 3, 3}
	for _, want := range expected {
		got, ok := h.Extract()
		if !ok {
			t.Fatalf("Extract() returned ok=false, want %d", want)
		}
		if got != want {
			t.Errorf("Extract() = %d, want %d", got, want)
		}
	}
}

func TestMaxHeap_SingleElement(t *testing.T) {
	h := NewMaxHeap[int]()
	h.Insert(Entry[int]{Priority: 42, Value: 42})

	val, ok := h.Peek()
	if !ok || val != 42 {
		t.Errorf("Peek() = (%d, %v), want (42, true)", val, ok)
	}

	got, ok := h.Extract()
	if !ok || got != 42 {
		t.Errorf("Extract() = (%d, %v), want (42, true)", got, ok)
	}

	if h.Size() != 0 {
		t.Errorf("Size() after extracting only element = %d, want 0", h.Size())
	}
}

func TestHeap_PeekPriority(t *testing.T) {
	h := NewMinHeap[string]()
	_, ok := h.PeekPriority()
	if ok {
		t.Error("PeekPriority() on empty heap should return ok=false")
	}

	h.Insert(Entry[string]{Priority: 10, Value: "ten"})
	h.Insert(Entry[string]{Priority: 3, Value: "three"})
	h.Insert(Entry[string]{Priority: 7, Value: "seven"})

	pri, ok := h.PeekPriority()
	if !ok {
		t.Fatal("PeekPriority() returned ok=false on non-empty heap")
	}
	if pri != 3 {
		t.Errorf("PeekPriority() = %d, want 3", pri)
	}

	// Should not remove the element
	if h.Size() != 3 {
		t.Errorf("Size() after PeekPriority = %d, want 3", h.Size())
	}
}

func TestHeap_PriorityVsValue(t *testing.T) {
	// Priority and Value are independent — extract by priority, get the value
	h := NewMinHeap[string]()
	h.Insert(Entry[string]{Priority: 3, Value: "low-priority"})
	h.Insert(Entry[string]{Priority: 1, Value: "high-priority"})
	h.Insert(Entry[string]{Priority: 2, Value: "mid-priority"})

	got, ok := h.Extract()
	if !ok || got != "high-priority" {
		t.Errorf("Extract() = (%s, %v), want (high-priority, true)", got, ok)
	}
	got, ok = h.Extract()
	if !ok || got != "mid-priority" {
		t.Errorf("Extract() = (%s, %v), want (mid-priority, true)", got, ok)
	}
	got, ok = h.Extract()
	if !ok || got != "low-priority" {
		t.Errorf("Extract() = (%s, %v), want (low-priority, true)", got, ok)
	}
}
