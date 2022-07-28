package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapStack(t *testing.T) {
	mem := newMapStack()

	v := mem.get("abc")
	require.Equal(t, "Nil", v)

	count := mem.numEqualTo("abc")
	require.Equal(t, 0, count)

	mem.set("abc", "def")
	v = mem.get("abc")
	require.Equal(t, "def", v)
	mem.unSet("abc")
	v = mem.get("abc")
	require.Equal(t, "Nil", v)

	mem.set("abc", "def")
	mem.set("key2", "value2")
	mem.begin()
	mem.set("abc", "123")
	mem.begin()
	mem.unSet("abc")
	v = mem.get("key2")
	require.Equal(t, "value2", v)
	v = mem.get("abc")
	require.Equal(t, "Nil", v)
	ok := mem.commit()
	require.True(t, ok)
	ok = mem.rollBack()
	require.True(t, ok)
	v = mem.get("abc")
	require.Equal(t, "def", v)

	ok = mem.rollBack()
	require.False(t, ok)
	ok = mem.commit()
	require.False(t, ok)
}
